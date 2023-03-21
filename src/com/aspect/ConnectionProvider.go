package aspect

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func GetProperties() string {
	data, err := ioutil.ReadFile("../../resources/DBConfig.properties")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	//fmt.Println("Contents of file:")
	//fmt.Println(string(data))
	splitedData := strings.Split(string(data), "\n")
	//fmt.Println(splitedData)
	Url := strings.Split(splitedData[1], "=")[1]
	UserName := strings.Split(splitedData[2], "=")[1]
	Password := strings.Split(splitedData[3], "=")[1]
	Database := strings.Split(splitedData[4], "=")[1]
	return fmt.Sprintf("%v:%v@tcp(%v)/%v", UserName, Password, Url, Database)
}

func ProvideConnection() *sql.DB {

	db, err := sql.Open("mysql", GetProperties())

	if err != nil {
		log.Fatal(err)
	}
//	defer db.Close()
	//Call db.Ping() to check the connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
