package serviceImplement

import "fmt"

type MaterialCal struct {
}

func CreateMaterialCal() MaterialCal {
	return MaterialCal{}
}

// DatabaseImplements db=new DatabaseImplements();
func (MaterialCal) cementCal(count int) int {
	if count >= 0 {
		count := count * GetRates("cement")
		return count
	} else {
		fmt.Println("Cement count must be positive value")
	}
	return 0
}
func (MaterialCal) sandCal(count int) int {
	if count >= 0 {
		count := count * GetRates("sand")
		return count
	} else {
		fmt.Println("Sand count must be positive value")
	}
	return 0
}
func (MaterialCal) brickCal(count int) int {
	if count >= 0 {
		count := count * GetRates("brick")
		return count
	} else {
		fmt.Println("Brick count must be positive value")
	}
	return 0
}
