package main

import (
	"fmt"
	"github.com/goroutinePool/src"
	"runtime"
	"time"
)

func main(){
	t := src.NewTask(func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})

	p := src.NewPool(4)

	tN := 0//测试数量
	go func() {
		for{
			p.AddTask(t)
			tN++
			if 0 == 1000%10 {
				fmt.Println(runtime.NumGoroutine())
				time.Sleep(time.Second)
			}
		}
	}()
	fmt.Println(runtime.NumGoroutine())
	p.Run()
}
