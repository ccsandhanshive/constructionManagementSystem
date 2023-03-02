package bean

import (
	"fmt"
)

type Labor struct {
	laborcount int
	sal        int
}

func CreateLabor(laborcount int, sal int) Labor {
	return Labor{laborcount: laborcount, sal: sal}
}
func (l Labor) getLaborcount() int {
	return l.laborcount
}
func (l Labor) setLaborcount(laborcount int) {
	l.laborcount = laborcount
}
func (l Labor) getSal() int {
	return l.sal
}
func (l Labor) setSal(sal int) {
	l.sal = sal
}
func (l Labor) String() string {
	return fmt.Sprintf("Labor [Labor=%d, sal=%d]", l.laborcount, l.sal)
}
