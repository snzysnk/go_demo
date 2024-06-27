package distribute

import (
	"fmt"
	"net/rpc"
)

type DoJobArgs struct {
	JobType string
	Urls    []string
}

type DoJobReply struct {
	Ok bool
}

type RegisterArgs struct {
	Worker string
}

type RegisterReply struct {
	Ok bool
}

func Call(address string, rpcName string, args interface{}, reply interface{}) bool {
	http, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		return false
	}
	defer http.Close()

	err = http.Call(rpcName, args, reply)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
