package distribute

import (
	"cw/crawler"
	"fmt"
	"net/http"
	"net/rpc"
)

type Worker struct {
	addr          string
	addUrlChannel chan bool
}

func initWorker(addr string) *Worker {
	return &Worker{
		addr:          addr,
		addUrlChannel: make(chan bool),
	}
}

// 向master中注册worker（调用Master.Register）
func register(workerAddress, masterAddress string) {
	args := &RegisterArgs{Worker: workerAddress}
	var reply RegisterReply
	Call(masterAddress, "Master.Register", args, &reply)
}

// DoJob 执行爬取数据
func (receiver Worker) DoJob(args *DoJobArgs, reply *DoJobReply) error {
	switch args.JobType {
	case "pc":
		crawler.Do(args.Urls)
		receiver.addUrlChannel <- true
	}
	return nil
}

func RunWorker(masterAddress, workerAddress string) {
	fmt.Println("------start rpc worker------")
	w := initWorker(workerAddress)
	go startRpcWorker(w)
	register(workerAddress, masterAddress)
	fmt.Println("------worker bind to master ok----")
	for {
		select {
		case <-w.addUrlChannel:
			fmt.Println("爬虫已开始处理")
		}
	}
}

func startRpcWorker(w *Worker) {
	rpc.Register(w)
	rpc.HandleHTTP()
	err := http.ListenAndServe(w.addr, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("开始退出")
	}
}
