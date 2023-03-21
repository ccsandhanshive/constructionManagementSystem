package serviceImplement

import (
	"fmt"

	"bitsys.sys/construction_management_system/aspect"
)
func Login(userid int, password string) string {
	db  = aspect.ProvideConnection()
	result, err := db.Query(fmt.Sprintf("select userid,password,role from login where userid=%d and password= %v", userid, password))

	if err != nil {
		fmt.Print(err)
		return ""
	}
	var role string

	for result.Next() {

		result.Scan(&userid, &password, &role)
	}
	/*try {
		p=con.prepareStatement("select userid,password,role from login where userid=? and password= ?");
		p.setInt(1,userid);
		p.setString(2,password);
		rs=p.executeQuery();
		rs.next();
		s=rs.getString(3);
		} catch (SQLException e) {
		// TODO Auto-generated catch block
	} */
	return role
}

func ChangePass(userid int, password string, newpassword string) int {
	_, err := db.Query("update login set password=%v where userid=%d and password=%v", newpassword, userid, password)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return userid
}
