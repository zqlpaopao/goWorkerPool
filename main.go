package main

import (
	"fmt"
	"runtime"
	"time"
)

//定义任务类型
type Task struct {
	f func() error
}

//创建一个任务
func NewTask(argF func() error) *Task {
	t := &Task{f: argF}
	return t
}

//Task执行业务的方法
func (t *Task) Execute() {
	t.f()
}

//------------------------------有关协程池Pool角色的功能------------------
//定义协程池
type Pool struct {
	//对外的入口,接收task
	EntryChannel chan *Task
	//内部Task队列
	JobsChannel chan *Task
	//协程的数量
	workerNum int
}

//创建协程池
func NewPool(n int) *Pool {
	p := &Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		workerNum:    n,
	}
	return p
}

//协程池的Worker，并且run
func (p *Pool) worker(workerId int) {
	//一直从内部取任务，渠道任务在执行
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("执行的任务的id", workerId)
	}
}

//让协程池开始工作
func (p *Pool) run() {
	//根据workerNum去工作
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}

	//将入口的任务加入到内部任务队列
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}

func main() {
	//创建任务
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	//创建协程池Pool
	p := NewPool(4)
	//将任务交给协程池
	taskNum := 0
	go func() {
		for {
			p.EntryChannel <- t
			taskNum++
			fmt.Println(taskNum)
			if 0 == 1000%10 {
				fmt.Println(1111)
				fmt.Println(runtime.NumGoroutine())
				time.Sleep(time.Second)
			}
		}
	}()
	p.run()
	fmt.Println(runtime.NumGoroutine())
}
