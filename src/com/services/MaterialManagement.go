package services

type MaterialManagement interface {
	brickCal(count int) int  //calculation of brick amount
	sandCal(count int) int   //calculation of sand amount
	cementCal(count int) int //calculation of cement amount
}
