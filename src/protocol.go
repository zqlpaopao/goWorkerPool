/**
 * @Author: zhangsan
 * @Description:
 * @File:  protocol
 * @Version: 1.0.0
 * @Date: 2021/2/23 上午9:50
 */

package src

type (
	//Assist in fibers
	poolInter interface {
		NewPool()
		AddTask()
		Run()
	}
	taskInter interface {
		execute()
	}
	//Task metadata
	task struct {
		f func()
	}
	//Concord pool role
	pool struct{
		receiveCh chan *task
		runCh chan *task
		workerNum int
	}
)

func HandlerExecute(f taskInter){f.execute()}

func HandlerNewPool(f poolInter){f.NewPool()}
func HandlerAddTask(f poolInter){f.AddTask()}
func HandlerRun(f poolInter){f.Run()}