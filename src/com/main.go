package main

import (
	"fmt"
	"os"
	"strconv"

	"bitsys.sys/construction_management_system/aspect"
	"bitsys.sys/construction_management_system/serviceImplement"
)

var press int
var date string
var days int
var labor []int
var worker []int
var watchman int
var s1 []string

type UserInterface struct {
	press    int
	date     string
	days     int
	labor    []int
	worker   []int
	watchman int
	s1       []string
}

func changes(userid int, password string, u UserInterface) {
	fmt.Println("1.change in site\n2.change in date\n3.change in labor worker on site\n4.change in watchman \n5.back\nAny no for exit")
	var press2 int
	fmt.Scanln(&press2)

	if press2 == 1 {
		s1 = serviceImplement.RetriveSite(userid, password) //for get supervision  sits from database
		press := site(userid, password, s1)
		u.press = press
		fmt.Println("site changed to" + s1[press])
		changes(userid, password, u)
	} else if press2 == 2 {
		fmt.Println("Enter date (DD-MM-YYYY)")
		var date string
		fmt.Scanln(&date)
		u.date = date
		changes(userid, password, u)
	} else if press2 == 3 {
		var days int
		fmt.Println(labor)
		fmt.Println(worker)
		fmt.Println("Enter days:")
		fmt.Scanln(&days)
		for i := 0; i < days; i++ {
			fmt.Println("day " + strconv.Itoa(i) + " labor:")
			fmt.Scanln(&labor[i])
			fmt.Println("day " + strconv.Itoa(i) + " worker")
			fmt.Scanln(&worker[i])
			fmt.Println("")
		}
		u.days = days
		changes(userid, password, u)
	} else if press2 == 4 {
		fmt.Println("Enter no of watchman")
		fmt.Scanln(&watchman)
		u.watchman = watchman
		changes(userid, password, u)
	} else if press2 == 5 {
		mainSupervisor(userid, password, u)
	} else {
		os.Exit(1)
	}
}

func site(userid int, password string, s1 []string) int {
	fmt.Println("Enter site:")
	fmt.Println(s1)
	press := 0
	for i := 1; i < len(s1)+1; i++ {
		fmt.Println(strconv.Itoa(i) + " " + s1[i-1])
	}
	fmt.Scanln(&press)
	return press
}

func supervisor(userid int, password string, u UserInterface) {
	fmt.Println("Enter date (DD-MM-YYYY)")
	fmt.Scanln(&date)
	s1 := serviceImplement.RetriveSite(userid, password) //for get supervision  sits from database
	press = site(userid, password, s1) - 1
	fmt.Println("Enter no of days")
	fmt.Scanln(&days)
	labor := []int{}
	worker := []int{}
	var tempInputlabor int
	var tempInputWorker int
	for i := 0; i < days; i++ {
		fmt.Println("day " + strconv.Itoa(i) + " labor:")
		fmt.Scanln(&tempInputlabor)
		labor = append(labor, tempInputlabor)
		fmt.Println("day " + strconv.Itoa(i) + " worker")
		fmt.Scanln(&tempInputWorker)
		worker = append(worker, tempInputWorker)
		fmt.Println("")
	}
	u.labor = labor
	u.worker = worker
	fmt.Println("Enter no of watchman")
	fmt.Scanln(&watchman)
	u.watchman = watchman

	fmt.Println("Enter cement count in bags")
	var cement int
	fmt.Scanln(&cement)
	fmt.Println("Enter sand count in trip")
	var sand int
	fmt.Scanln(&sand)
	fmt.Println("Enter brick count")
	var brick int
	fmt.Scanln(&brick)

	fmt.Println(serviceImplement.DataEntry01withcommit(s1[press], date, labor, worker, watchman, 0)) //for show temporary data before save in database
	fmt.Println(serviceImplement.MaterialEntrywithcommit(s1[press], date, cement, sand, brick, 0))
	fmt.Println("Save data \n1.yes \n2.no \n3.back")
	var press1 int
	fmt.Scanln(&press1)
	if press1 == 1 {
		serviceImplement.DataEntry01withcommit(s1[press], date, labor, worker, watchman, 1)
		serviceImplement.MaterialEntrywithcommit(s1[press], date, cement, sand, brick, 0) //Save in database
		mainSupervisor(userid, password, u)
	} else if press1 == 2 {
		changes(userid, password, u)
		mainSupervisor(userid, password, u)
	} else if press1 == 3 {
		supervisor(userid, password, u)
		mainSupervisor(userid, password, u)
	}

}

func passChange(userid int, password string) {
	fmt.Println("Press 1 for change password\nAny other no for skip")
	var press int
	fmt.Scanln(&press)
	if press == 1 {
		fmt.Println("Enter New Password")
		var newpass string
		fmt.Scanln(&newpass)
		fmt.Println("Are you sure\n1.yes 2.no 3.back")
		fmt.Scanln(&press)
		if press == 1 {
			serviceImplement.ChangePass(userid, password, newpass)
			fmt.Println(newpass)
			passChange(userid, password)
		}
	}
}

func mainSupervisor(userid int, password string, u UserInterface) {
	fmt.Println("1.Enter Data\n2.retrive data via site\n3.retrive data via site and date\n4.logout\nAny no for exit")
	var c int
	fmt.Scanln(&c)
	if c == 1 {
		supervisor(userid, password, u)
		mainSupervisor(userid, password, u)
	} else if c == 2 {
		fmt.Println("Enter sitename")
		var site string
		fmt.Scanln(&site)
		serviceImplement.RetriveData01(site)
		mainSupervisor(userid, password, u)
	} else if c == 3 {
		fmt.Println("Enter sitename:")
		var site string
		fmt.Scanln(&site)
		fmt.Println("Enter date")
		var date string
		fmt.Scanln(&date)
		serviceImplement.RetriveData02(site, date)
		mainSupervisor(userid, password, u)
	} else if c == 4 {
		fmt.Scanf("User logout successfully")
		u.login(u)
	} else {
		os.Exit(1)
	}

}
func (UserInterface) contractor(u UserInterface) {
	fmt.Println("1.retrive data\n2.deletedata\n3.changes in rate\n4.User related changes\n5.logout")
	var press int
	fmt.Scanln(&press)
	if press == 1 {
		fmt.Println("1.with site name\n2.with site name and date\n3.back\nAny other no for exit")
		var press1 int
		fmt.Scanln(&press1)
		if press1 == 1 {
			fmt.Println("Enter Site name")
			var sitename string
			fmt.Scanln(&sitename)
			serviceImplement.RetriveData01(sitename)
			u.contractor(u)
		} else if press1 == 2 {
			fmt.Println("Enter site name")
			var sitename string
			fmt.Scanln(&sitename)
			fmt.Println("Enter date")
			var date string
			fmt.Scanln(&date)
			serviceImplement.RetriveData02(sitename, date)
			u.contractor(u)
		} else if press1 == 3 {
			press = 0
			u.contractor(u)
		} else {
			os.Exit(1)
		}
	}
	if press == 2 {
		fmt.Println("1.with site\n2.site and date\n3.back\nAny other no for exit")
		var press2 int
		fmt.Scanln(&press2)
		if press2 == 1 {
			fmt.Println("Enter site name")
			var sitename string
			fmt.Scanln(&sitename)
			serviceImplement.DeleteData01(sitename)
			u.contractor(u)
		} else if press2 == 2 {
			fmt.Println("Enter site name")
			var sitename string
			fmt.Scanln(&sitename)
			fmt.Println("Enter date")
			var date string
			fmt.Scanln(&date)
			serviceImplement.DeleteData02(sitename, date)
			u.contractor(u)
		} else if press2 == 3 {
			u.contractor(u)
		} else {
			os.Exit(1)
		}

	}
	if press == 3 {
		fmt.Println("1.labor\n2.worker\n3.watchman\n4.cement\n5.sand\n6.brick\n7.back\nAny other no for exit")
		var ans int
		fmt.Scanln(&ans)
		fmt.Println("Enter rate")
		var rate int
		fmt.Scanln(&rate)
		fmt.Println("save rate\n1.yes\n2.no")
		var ans1 int
		fmt.Scanln(&ans1)
		if ans == 1 {
			serviceImplement.LaborRateChangewithcommit(rate, ans1)
		} else if ans == 2 {
			serviceImplement.WorkerRateChangewithcommit(rate, ans1)
		} else if ans == 3 {
			serviceImplement.WatchmanRateChangewithcommit(rate, ans1)
		} else if ans == 4 {
			serviceImplement.CementRateChangewithcommit(rate, ans1)
		} else if ans == 5 {
			serviceImplement.SandRateChangewithcommit(rate, ans1)
		} else if ans == 6 {
			serviceImplement.BrickRateChangewithcommit(rate, ans1)
		} else if ans == 7 {
			u.contractor(u)
		} else {
			os.Exit(1)
		}
		u.contractor(u)

	}
	if press == 4 {
		fmt.Println("1.change in site\n2.detele user\n3.insert user\n4.display data\n5.back\nAny other no for exit")
		var sp int
		fmt.Scanln(&sp)
		if sp == 1 {
			fmt.Println("Enter userid")
			var username int
			fmt.Scanln(&username)
			fmt.Println("Enter new site name")
			var newsite string
			fmt.Scanln(&newsite)
			fmt.Println("Save data \n1.yes\n2.no")
			var n int
			fmt.Scanln(&n)
			serviceImplement.SupervisorSiteChangewithcommit(username, newsite, n)
			u.contractor(u)
		} else if sp == 2 {
			fmt.Println("Enter userid")
			var userid int
			fmt.Println(userid)
			fmt.Println("Are you sure?")
			var commit int
			fmt.Println(commit)
			serviceImplement.DeleteUser(userid)
			u.contractor(u)
		} else if sp == 3 {
			fmt.Println("Enter sitename")
			var sitename string
			fmt.Scanln(&sitename)
			fmt.Println("Enter role")
			var role string
			fmt.Scanln(&role)
			fmt.Println("Are you sure\n1.yes\n2.no")
			var commit int
			fmt.Scanln(&commit)
			fmt.Println(serviceImplement.InsertUserwithcommit(sitename, role, commit))
			u.contractor(u)
		} else if sp == 4 {
			fmt.Println("1.all\n2.userid\n3.role based\n4.back\nAny no for exit")
			var r int
			fmt.Scanln(&r)
			if r == 1 {
				serviceImplement.RetriveUser01withcommit(1)
				u.contractor(u)
			} else if r == 2 {
				fmt.Println("Enter userid")
				var userid int
				fmt.Scanln(&userid)
				fmt.Println(serviceImplement.RetriveUser02withcommit(userid, 1))
				u.contractor(u)

			} else if r == 3 {
				fmt.Println("Enter role")
				var role int
				fmt.Scanln(&role)
				fmt.Println(serviceImplement.RetriveUser02withcommit(role, 1))
				u.contractor(u)
			} else if r == 4 {
				u.contractor(u)
			} else {
				os.Exit(0)
			}
		} else if sp == 5 {
			u.contractor(u)
		} else {
			os.Exit(1)
		}

	} else if press == 5 {
		u.login(u)
	} else {
		os.Exit(1)
	}
}

func (UserInterface) login(u UserInterface) {
	fmt.Println("Enter 000 userid for exit")
	fmt.Println("Enter userid:")
	var userid int
	fmt.Scanln(&userid)
	if userid == 000 {
		os.Exit(1)
	}

	fmt.Println("Enter password:")
	var password string
	fmt.Scanln(&password)
	role := serviceImplement.Login(userid, password) //login method in LoginImplement return role of user
	fmt.Println(role)
	if role != "" {
		passChange(userid, password)
		if role == "supervisor" {
			mainSupervisor(userid, password, u)
		} else if role == "customer" {
			fmt.Println("Your site data:-")
			serviceImplement.RetriveData01(serviceImplement.RetriveSite01(userid, password))
			fmt.Println("End")
			u.login(u)
		} else if role == "contractor" {
			u.contractor(u)
			u.login(u)
		} else {
			os.Exit(1)
		}
	} else //if invalid user
	{
		fmt.Println("Invalid user")
		u.login(u)
	}
}

func main() {
	aspect.ProvideConnection()
	u := UserInterface{
		press:    press,
		date:     date,
		days:     days,
		labor:    labor,
		worker:   worker,
		watchman: watchman,
		s1:       s1,
	}
	fmt.Println(u)
	fmt.Println("1.Login\nAny other no for Exit")
	var ch int
	fmt.Scanln(&ch)
	if ch == 1 {
		u.login(u)
	}

	os.Exit(1)

}
