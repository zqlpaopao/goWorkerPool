/**
 * @Author: zHangSan
 * @Description:
 * @File:  task
 * @Version: 1.0.0
 * @Date: 2021/2/23 上午9:50
 */

package src


//-- ------------------------------------
//--> @Description Initialization task
//--> @Param
//--> @return
//-- ----------------------------
func NewTask(argF func())*task {
	return &task{f: argF}
}

//-- ----------------------------------------------------
//--> @Description Schedule the task execution principal
//--> @Param
//--> @return
//-- ----------------------------
func (t *task)execute(){
	t.f()
}