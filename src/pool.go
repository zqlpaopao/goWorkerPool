/**
 * @Author: zhangsan
 * @Description:
 * @File:  pool
 * @Version: 1.0.0
 * @Date: 2021/2/23 上午9:50
 */

package src

import "fmt"

//-- --------------------------------------------
//--> @Description Initialize the coroutine pool
//--> @Param
//--> @return
//-- ----------------------------
func NewPool(n int)*pool {
	return &pool{
		receiveCh: make(chan *task),
		runCh: make(chan *task),
		workerNum: n,
	}
}

func(p *pool)AddTask(f *task){
	p.receiveCh<-f
}

//-- ----------------------------
//--> @Description worker执行器
//--> @Param
//--> @return
//-- ----------------------------
func(p *pool)worker(i int){
	for task := range p.runCh{
		task.execute()
		fmt.Println("执行任务的协程id",i)
	}
}

//-- --------------------------------------------
//--> @Description Coroutine work initialization
//--> @Param
//--> @return
//-- ----------------------------
func(p *pool)Run(){
	for i := 0;i < p.workerNum;i++{
		go p.worker(i)
	}
	for task:= range p.receiveCh{
		p.runCh<-task
	}
}