package serviceImplement

import (
	"fmt"

	"bitsys.sys/construction_management_system/bean"
)

var c int

// if Commit ==1 then and then only data will be save on database after save data in database user cannot be modify it only Contractor can be deleted this record
// and only supervisor has authority to add data and modify before save	its applicable to all operations
//
//	Connection con=ConnectionProvider.provideConnection();
//		DatabaseImplements d=new DatabaseImplements();
func DataEntry01withcommit(sitename string, date string, no_of_labor []int, no_of_worker []int, no_of_watchman int, commit int) string {
	var l bean.Labor
	var w bean.Worker
	var w1 bean.Watchman
	var labor_amt int    //total amount of labor
	var worker_amt int   //total amount of workers
	var watchman_amt int //total amount of watchman
	var total_amt int    //total amount
	var lcount int       //count total number of labors
	var wcount int       //count total number of worker

	for i := 0; i < len(no_of_labor); i++ {
		lcount += no_of_labor[i]
	}

	for i := 0; i < len(no_of_worker); i++ {
		wcount += no_of_worker[i]
	}
	labor_amt = lcount * GetRates("labor") //get rate from rate table and multiply by there counts
	worker_amt = wcount * GetRates("worker")
	watchman_amt = no_of_watchman * GetRates("watchman")
	total_amt = labor_amt + worker_amt + watchman_amt
	l = bean.CreateLabor(lcount, labor_amt)
	w = bean.CreateWorker(wcount, worker_amt)
	w1 = bean.CreateWatchman(no_of_watchman, watchman_amt)
	if commit == 1 { //if user wish to sava data in database
		DataEntry01(sitename, date, lcount, labor_amt, wcount, worker_amt, no_of_watchman, watchman_amt, total_amt)
	}
	return fmt.Sprintf("%s \n %s \n %s \n total amt= %d", l, w, w1, total_amt) //print all details for check

}

func RetriveDatawithcommit(sitename string, commit int) int {
	if commit == 1 {
		c = RetriveData01(sitename)
	}
	return c
}

func RetriveData02withcommit(sitename string, date string, commit int) int {
	if commit == 1 {
		c = RetriveData02(sitename, date)
	}

	return c
}

func DeleteData01withcommit(sitename string, commit int) int {
	if commit == 1 {
		c = DeleteData01(sitename)
	}

	return c
}

func DeleteData02withcommit(sitename string, date string, commit int) int {
	if commit == 1 {
		c = DeleteData02(sitename, date)
	}

	return c
}

func LaborRateChangewithcommit(rate int, commit int) int {
	if commit == 1 {
		c = LaborRateChange(rate)
	}

	return c
}

func WorkerRateChangewithcommit(rate int, commit int) int {
	if commit == 1 {
		c = WorkerRateChange(rate)
	}
	return c
}

func CementRateChangewithcommit(rate int, commit int) int {
	if commit == 1 {
		c = CementRateChange(rate)
	}
	return c
}

func SandRateChangewithcommit(rate int, commit int) int {
	if commit == 1 {
		c = SandRateChange(rate)
	}
	return c
}

func BrickRateChangewithcommit(rate int, commit int) int {
	if commit == 1 {
		return BrickRateChange(rate)
	}
	return 0
}

func InsertUserwithcommit(sitename string, role string, commit int) string {
	if commit == 1 {
		return InsertUser(sitename, role)
	}
	return ""
}

func SupervisorSiteChangewithcommit(username int, newsite string, commit int) int {
	if commit == 1 {
		c = SupervisorSiteChange(username, newsite)
	}
	return c
}

func CustomerSiteChangewithcommit(username int, password string, sitename string, commit int) int {
	if commit == 1 {
		c = CustomerSiteChange(username, password, sitename)
	}
	return c
}

func WatchmanRateChangewithcommit(rate int, commit int) int {
	if commit == 1 {
		c = WatchmanRateChange(rate)
	}
	return c
}

func DeleteUserwithcommit(userid int, commit int) int {
	if commit == 1 {
		c = DeleteUser(userid)
	}
	return c
}

func RetriveUser01withcommit(commit int) string {
	if commit == 1 {
		return RetriveUser01()
	}
	return ""
}
func RetriveUser02withcommit(userid int, commit int) string {
	if commit == 1 {
		return RetriveUser02(userid)
	}
	return ""
}

func RetriveUser03withcommit(role string, commit int) string {
	if commit == 1 {
		return RetriveUser03(role)
	}
	return ""
}

var mc MaterialCal = CreateMaterialCal()

func MaterialEntrywithcommit(sitename string, date string, cement_count int, sand_count int, brick_count int, commit int) string {
	var cement_amt int
	var sand_amt int
	var brick_amt int
	var total_amt int
	cement_amt = mc.CementCal(cement_count)
	sand_amt = mc.SandCal(sand_count)
	brick_amt = mc.BrickCal(brick_count)
	total_amt = cement_amt + sand_amt + brick_amt
	if commit == 1 {
		c := MaterialEntry(sitename, date, cement_count, cement_amt, sand_count, sand_amt, brick_count, brick_amt, total_amt)
		fmt.Println(c)
	}

	return fmt.Sprintf("cement count= %d cement amount=%d \n sand count=%d sand amount= %d \n brick count=%d brick amount=%d \n Total amount=%d", cement_count, cement_amt, sand_count, sand_amt, brick_count, brick_amt, total_amt)
}
