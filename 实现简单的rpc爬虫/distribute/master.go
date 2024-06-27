package distribute

import (
	"fmt"
	"net/http"
	"net/rpc"
	"strconv"
)

type Master struct {
	addr         string
	regChan      chan string
	workDownChan chan string
	jobChan      chan string
	workers      map[*WorkInfo]bool
}

type WorkInfo struct {
	workAddr string
}

func initMaster(address string) (master *Master) {
	return &Master{
		addr:         address,
		regChan:      make(chan string),
		workDownChan: make(chan string),
		jobChan:      make(chan string, 2),
		workers:      make(map[*WorkInfo]bool),
	}
}

func (m *Master) Register(args *RegisterArgs, res *RegisterReply) error {
	m.regChan <- args.Worker
	res.Ok = true
	return nil
}

func getTarget(jobChan chan string) {
	for i := 0; i < 10; i++ {
		url := "https://github.com/search?q=go&type=Repositories&p=" + strconv.Itoa(i+1)
		jobChan <- url
	}
}

func RunRpcMaster(address string) {
	fmt.Println(33333)
	master := initMaster(address)
	go startRpcMaster(master)
	go getTarget(master.jobChan)
	fmt.Println("master节点已启动")
	for {
		select {
		case workAddr := <-master.regChan:
			fmt.Println("work:" + workAddr + "已成功绑定到master")
			work := &WorkInfo{workAddr: workAddr}
			master.workers[work] = true
			dispatchJob(work, master)
		}
	}
}

func dispatchJob(workInfo *WorkInfo, m *Master) {
	var urls []string

	for i := 0; i < 10; i++ {
		url := <-m.jobChan
		urls = append(urls, url)
	}
	args := &DoJobArgs{JobType: "pc", Urls: urls}
	var reply DoJobReply

	err := Call(workInfo.workAddr, "Worker.DoJob", args, &reply)
	if err == true {
		m.workers[workInfo] = false
		fmt.Println("已分配任务到worker:" + workInfo.workAddr)
	}
}

func startRpcMaster(m *Master) {
	rpc.Register(m)
	rpc.HandleHTTP()
	err := http.ListenAndServe(m.addr, nil)
	if err != nil {
		fmt.Println("start rpc master err")
	} else {
		fmt.Println("start rpc master terminate")
	}
}
