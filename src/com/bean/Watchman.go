package bean

type Watchman struct {
	count int
	sal   int
}

func CreateWatchman(count int, sal int) Watchman {
	return Watchman{count: count, sal: sal}
}
func (w Watchman) getCount() int {
	return w.count
}
func (w Watchman) setCount(count int) {
	w.count = count
}
func (w Watchman) getSal() int {
	return w.sal
}
func (w Watchman) setSal(int sal) {
	w.sal = sal
}
