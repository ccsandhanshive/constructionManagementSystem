package aspect

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
import (

	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)
*/
func GetProperties() {
	/*	try {
			fis=new FileInputStream(".//resources//DBConfig.properties");
			int x=0;
			p=new Properties();
			p.load(fis);


		s[0]=p.getProperty("DriverName");
		s[1]=p.getProperty("Url");
		s[2]=p.getProperty("UserName");
		s[3]=p.getProperty("Password");
			}catch(Exception e) {
				e.getStackTrace();
			} */
}

func ProvideConnection() {
	db, err := sql.Open("mysql", "admin:adminPass@tcp(database-1.ccbunm4fdi2j.ap-northeast-1.rds.amazonaws.com:3306)/construction")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//Call db.Ping() to check the connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	/*
	   var dbName string = "database-1"
	   var dbUser string = "admin"
	   var dbHost string = "database-1.ccbunm4fdi2j.ap-northeast-1.rds.amazonaws.com"
	   var dbPort int = 3306
	   var dbEndpoint string = fmt.Sprintf("%s:%d", dbHost, dbPort)
	   var region string = "ap-northeast-1"

	   cfg, err := config.LoadDefaultConfig(context.TODO())

	   	if err != nil {
	   		panic("configuration error: " + err.Error())
	   	}

	   authenticationToken, err := auth.BuildAuthToken(

	   	context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)

	   	if err != nil {
	   		panic("failed to create authentication token: " + err.Error())
	   	}

	   dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&allowCleartextPasswords=true",

	   	dbUser, authenticationToken, dbEndpoint, dbName,

	   )

	   db, err := sql.Open("mysql", dsn)

	   	if err != nil {
	   		panic(err)
	   	}

	   err = db.Ping()

	   	if err != nil {
	   		fmt.Println(err)
	   	}
	*/
}
