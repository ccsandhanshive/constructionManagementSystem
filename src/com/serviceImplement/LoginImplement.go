package serviceImplement

import "fmt"

func Login(userid int, password string) string {

	s := "supervisor"
	fmt.Println(userid, password)
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
	return s
}

func changePass(username int, password string, newpassword string) int {
	fmt.Println(username, password, newpassword)
	/* try {
		p=con.prepareStatement("update  login set password=? where userid=? and password=? ");
		p.setString(1,newpassword);
		p.setInt(2,username);
		p.setString(3,password);
		c=p.executeUpdate();

	} catch (SQLException e) {
		// TODO Auto-generated catch block
		e.printStackTrace();
	} */

	return username
}
