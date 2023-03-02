package services

type Contractor interface {
	RetriveData01(sitename string) int                                     //retrive data using sitename
	RetriveData02(sitename string, date string) int                        //retrive  data using sitename and date
	DeleteData01(sitename string) int                                      //delete  data using sitename
	DeleteData02(sitename string, date string) int                         //delete data using sitename and date
	laborRateChange(rate int) int                                          //change in bebor rate
	workerRateChange(rate int) int                                         //change in worker rate
	cementRateChange(rate int) int                                         //change in cement rate
	sandRateChange(rate int) int                                           //change in sand rate
	brickRateChange(rate int) int                                          //change in brick rate
	insertUser(sitename string, role string) string                        //inert new user
	supervisorSiteChange(username int, newsite string) int                 //change in supervisor supervision site
	customerSiteChange(username int, password string, sitename string) int //change in customer site in this case customer password is must

}
