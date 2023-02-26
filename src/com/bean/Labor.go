package bean

type Labor struct{
laborcount int
sal int
}
func CreateLabor(laborcount int, sal int) Labor {
	return Labor{laborcount: laborcount, sal: sal}
}
type (l Labor) getLaborcount() int{
	return l.laborcount;
}
type (l Labor) setLaborcount(laborcount int) {
	l.laborcount = laborcount
}
type (l Labor) getSal() int{
	return sal;
}
type (l Labor) setSal(sal int) {
	l.sal = sal;
}
