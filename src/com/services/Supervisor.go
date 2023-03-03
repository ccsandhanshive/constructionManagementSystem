package services

type Supervisor interface {
	dataEntry(sitename string, date string, no_of_labor int, labor_amt int, no_of_worker int, worker_amt int, no_of_watchman int, watchman_amt int, total_amt int) int //insert lebor, worker //and watchaman count
	retriveData01(sitename string) int                                                                                                                                 // retrive data of perticuler site
	retriveData02(sitename string, date string) int                                                                                                                    //retrive data of perticuler site using date

}
