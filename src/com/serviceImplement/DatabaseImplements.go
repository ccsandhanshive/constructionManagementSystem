package serviceImplement

import (
	"database/sql"
	"fmt"

	"bitsys.sys/construction_management_system/aspect"
)

var db *sql.DB = aspect.ProvideConnection()

// @Override
func DataEntry01(sitename string, date string, no_of_labor int, labor_amt int, no_of_worker int, worker_amt int, no_of_watchman int, watchman_amt int, total_amt int) int {
	dataEntry := fmt.Sprintf("insert into %v values(%v,%d,%d,%d,%d,%d,%d,%d)", sitename, date, no_of_labor, labor_amt, no_of_worker, worker_amt, no_of_watchman, watchman_amt, total_amt)
	_, err := db.Query(dataEntry)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Print("Successfully  Inserted\n")
	defer db.Close()
	return 1
}

func RetriveData01(sitename string) int {
	dataEntry := fmt.Sprintf("select w.date_,w.no_of_labor,w.laber_amt,w.no_of_worker,w.worker_amt,w.no_of_watchman,w.watchman_amt,w1.cement_count,w1.cement_amt,w1.sand_count,w1.sand_amt,w1.brick_count,w1.brick_amt,w.total_amt+w1.total_amount as total_amount from %v w, %vMaterial w1 where w.date_=w1.date_", sitename, sitename)
	result, err := db.Query(dataEntry)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	for result.Next() {

		var date_ string
		var no_of_labor int
		var laber_amt int
		var no_of_worker int
		var worker_amt int
		var no_of_watchman int
		var watchman_amt int
		var cement_count int
		var cement_amt int
		var sand_count int
		var sand_amt int
		var brick_count int
		var brick_amt int
		var total_amount int

		// The result object provided Scan  method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = result.Scan(&date_, &no_of_labor, &laber_amt, &no_of_worker, &worker_amt, &no_of_watchman, &watchman_amt, &cement_count, &cement_amt, &sand_count, &sand_amt, &brick_count, &brick_amt, &total_amount)

		// handle error
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Printf(date_, no_of_labor, laber_amt, no_of_worker, worker_amt, no_of_watchman, watchman_amt, cement_count, cement_amt, sand_count, sand_amt, brick_count, brick_amt, total_amount)
	}

	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}

func RetriveData02(sitename string, date string) int {
	dataEntry := fmt.Sprintf("select\r\n"+
		" w.date_,\r\n"+
		"w.no_of_labor,\r\n"+
		"w.laber_amt,\r\n"+
		"w.no_of_worker,\r\n"+
		"w.worker_amt,\r\n"+
		"w.no_of_watchman,\r\n"+
		"w.watchman_amt,\r\n"+
		"w1.cement_count,\r\n"+
		"w1.cement_amt,\r\n"+
		"w1.sand_count,\r\n"+
		"w1.sand_amt,\r\n"+
		"w1.brick_count,\r\n"+
		"w1.brick_amt,\r\n"+
		"w.total_amt+w1.total_amount as total_amount \r\n"+
		"from %v w, %vMaterial w1 where w.date_=w1.date_ and w.date_='%v'", sitename, sitename, date)
	result, err := db.Query(dataEntry)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	for result.Next() {

		var date_ string
		var no_of_labor int
		var laber_amt int
		var no_of_worker int
		var worker_amt int
		var no_of_watchman int
		var watchman_amt int
		var cement_count int
		var cement_amt int
		var sand_count int
		var sand_amt int
		var brick_count int
		var brick_amt int
		var total_amount int

		// The result object provided Scan  method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = result.Scan(&date_, &no_of_labor, &laber_amt, &no_of_worker, &worker_amt, &no_of_watchman, &watchman_amt, &cement_count, &cement_amt, &sand_count, &sand_amt, &brick_count, &brick_amt, &total_amount)

		// handle error
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Printf(date_, no_of_labor, laber_amt, no_of_worker, worker_amt, no_of_watchman, watchman_amt, cement_count, cement_amt, sand_count, sand_amt, brick_count, brick_amt, total_amount)
	}

	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}

func DeleteData01(sitename string) int {
	dataEntry := fmt.Sprintf("truncate table %v", sitename)
	_, err := db.Query(dataEntry)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}

func DeleteData02(sitename string, date string) int {
	dataDelete := fmt.Sprintf("delete from %v where date_ = '%v'", sitename, date)
	_, err := db.Query(dataDelete)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	dataDelete = fmt.Sprintf("delete from %vMaterial  where date_='%v'", sitename, date)
	_, err = db.Query(dataDelete)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}

func LaborRateChange(rate int) int {
	labor_rate_change := fmt.Sprintf("update rate set labor=%d ", rate)
	_, err := db.Query(labor_rate_change)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}

func WorkerRateChange(rate int) int {
	worker_rate_change := fmt.Sprintf("update rate set worker=%d", rate)
	_, err := db.Query(worker_rate_change)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}

func CementRateChange(rate int) int {
	cement_rate_change := fmt.Sprintf("update rate set cement=%d ", rate)
	_, err := db.Query(cement_rate_change)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}

func SandRateChange(rate int) int {
	sand_rate_change := fmt.Sprintf("update rate set sand=%d ", rate)
	_, err := db.Query(sand_rate_change)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
}

func BrickRateChange(rate int) int {
	sand_rate_change := fmt.Sprintf("update rate set brick=%d where brick>0", rate)
	_, err := db.Query(sand_rate_change)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}

func InsertUser(sitename string, role string) string {
	userid := RetriveData03() + 1
	password := fmt.Sprintf("%v/%v/%d", sitename, role, userid)
	insert_user := fmt.Sprintf("insert into login values(%d,%v,%v,%v)", userid, password, role, sitename)
	_, err := db.Query(insert_user)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	fmt.Println("new user inserted")
	defer db.Close()
	return fmt.Sprintln("userid= %v pass= %v", userid, password)
}

func SupervisorSiteChange(username int, newsite string) int {
	supervisor_site_change := fmt.Sprintf("update login set sitename=%v where userid=%d", newsite, username)
	_, err := db.Query(supervisor_site_change)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
}

func CustomerSiteChange(userid int, password string, sitename string) int {
	customer_site_change := fmt.Sprintf("update login set sitename=? where userid=? and password=?", sitename, userid, password)
	_, err := db.Query(customer_site_change)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
}
func RetriveData03() int {
	result, err := db.Query("select max(userid) from login")

	if err != nil {
		fmt.Println(err)
		return 0
	}

	result.Next(){
		
	}
	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
	return 1
	/* try {
		p=con.prepareStatement("select max(userid) from login");
	rs=p.executeQuery();
	rs.next();
	c=rs.getInt(1);


	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */
	return 0
}

func GetRates(ratenm string) int {
	/*	try {
			p=con.prepareStatement("select * from rate");
			rs=p.executeQuery();
			rs.next();
			c=rs.getInt(ratenm);
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} */
	return 3

}

func RetriveSite(userid int, pass string) []string {
	s := []string{"wagholi", "rajpark", "chikhali", "tathawade"}
	/* try {
		p=con.prepareStatement("select sitename from login where userid=? and password=?");
		p.setInt(1,userid);
		p.setString(2,pass);
		rs=p.executeQuery();
		rs.next();
	 s=rs.getString(1);
	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */
	return s
}
func RetriveSite01(userid int, pass string) string {
	/* try {
		p=con.prepareStatement("select sitename from login where userid=? and password=?");
		p.setInt(1,userid);
		p.setString(2,pass);
		rs=p.executeQuery();
		rs.next();
	 s=rs.getString(1);
	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */
	return "wagholi"
}

func WatchmanRateChange(rate int) int {
	/* try {
	p=con.prepareStatement("update rate set watchman=? ");
	p.setInt(1,rate);
	c=p.executeUpdate();
	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */
	return 0
}

func DeleteUser(userid int) int {
	/* try {
		s=con.createStatement();
	    c=s.executeUpdate("delete from login where userid ="+userid);
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}catch(Exception e1){
			e1.printStackTrace();
		} */
	return 0
}

func RetriveUser01() string {
	//String s=null;
	/* try {
		p=con.prepareStatement("select userid,role,sitename from login ");
		rs=p.executeQuery();
		while(rs.next()) {
			System.out.println("userid= "+rs.getInt(1)+"\nrole = "+rs.getString(2)+"\nsitename = "+rs.getString(3));
		}
	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */
	return "123"
}

func RetriveUser02(userid int) string {
	/*	String s=null;
		try {
			p=con.prepareStatement("select userid,role,sitename from login where userid=?");
			p.setInt(1,userid);
			rs=p.executeQuery();
			while(rs.next()) {
				System.out.println("userid= "+rs.getInt(1)+"\nrole = "+rs.getString(2)+"\nsitename = "+rs.getString(3));
			}
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} */
	return "123"
}

func RetriveUser03(role string) string {
	/*String s=null;
	try {
		p=con.prepareStatement("select userid,role,sitename from login where role=?");
		p.setInt(1,userid);
		rs=p.executeQuery();
		while(rs.next()) {
			System.out.println("userid= "+rs.getInt(1)+"\nrole = "+rs.getString(2)+"\nsitename = "+rs.getString(3));
		}
	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */
	return "123"
}

func CreateSiteTable(sitename string) bool {
	/* boolean b=false;
	try {
		p=con.prepareStatement("create table "+sitename+"(date_ varchar2(30),no_of_labor int,laber_amt int,no_of_worker int,worker_amt int,no_of_watchman int,watchman_amt int,total_amt int)");
		b=p.execute();
	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */
	return true
}

func MaterialEntry(sitename string, date string, cement_count int, cement_amt int, sand_count int, sand_amt int, brick_count int, brick_amt int, total_amt int) int {
	/* c=0;
	System.out.println("In materialEntry "+sitename+"Material");

		try {
			p=con.prepareStatement("insert into "+sitename+"Material values(?,?,?,?,?,?,?,?)");
			p.setString(1,date);
			p.setInt(2,cement_count);
			p.setInt(3,cement_amt);
			p.setInt(4,sand_count);
			p.setInt(5,sand_amt);
			p.setInt(6,brick_count);
			p.setInt(7,brick_amt);
			p.setInt(8,total_amt);
			c=p.executeUpdate();
		} catch (SQLException e) {
			// TODO Auto-generated catch block

		boolean b=createMaterialTable(sitename+"Material");//if table not exit then create table table name is site name

		materialEntry(sitename,date,cement_count,cement_amt,sand_count,sand_amt,brick_count,brick_amt,total_amt); //Try to enter data again

			//e.printStackTrace();
		} */
	return 0
}
func CreateMaterialTable(sitename string) bool {
	/* System.out.println("In create "+sitename);
	boolean b=false;
	try {
		p=con.prepareStatement("create table "+sitename+"(date_ varchar2(30),cement_count int,cement_amt int,sand_count int,sand_amt int,brick_count int,brick_amt int,total_amount int)");
		b=p.execute();
	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */
	return true
}
