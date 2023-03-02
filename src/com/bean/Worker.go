package bean

import (
	"fmt"
)

type Worker struct {
	workercount int
	sal         int
}

func CreateWorker(workercount int, sal int) Worker {
	return Worker{workercount: workercount, sal: sal}
}
func (w Worker) getWorkercount() int {
	return w.workercount
}
func (w Worker) setWorkercount(workercount int) {
	w.workercount = workercount
}
func (w Worker) getSal() int {
	return w.sal
}
func (w Worker) setSal(sal int) {
	w.sal = sal
}
func (w Worker) String() string {
	return fmt.Sprintf(" Worker [workercount =%d, sal=%d]", w.workercount, w.sal)
}
