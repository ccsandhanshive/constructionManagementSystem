package services

type Customer interface {
	RetriveData01(sitename string) int              //retrive data of perticuler site
	RetriveData02(sitename string, date string) int //retrive data of perticuler site using date

}
