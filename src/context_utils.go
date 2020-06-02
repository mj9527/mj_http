/**
 * @Author: mjzheng
 * @Description:
 * @File:  context_utils.go
 * @Version: 1.0.0
 * @Date: 2020/6/1 下午5:24
 */

package main

import (
	"context"
	"fmt"
	"time"
)

func UseContext() {
	ctx := context.Background()

	//go monitor(ctx, 1)

	ctx1, cancel := context.WithCancel(ctx)

	go monitor(ctx1, 2)

	ctx2, _ := context.WithTimeout(ctx1, time.Second)

	deadline, ok := ctx2.Deadline()
	if ok {
		fmt.Println("with dealine ", deadline)
	}

	go monitor(ctx2, 1)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()

	time.Sleep(5 * time.Second)
}

func monitorUseChan(stop chan bool, index int) {
	for {
		select {
		case <-stop:
			fmt.Println("监控退出，停止了...", index)
			return
		default:
			fmt.Println("goroutine监控中...", index)
			time.Sleep(2 * time.Second)
		}
	}
}

func UseChan() {
	stop := make(chan bool)

	go monitorUseChan(stop, 1)

	go monitorUseChan(stop, 2)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	time.Sleep(5 * time.Second)

}
