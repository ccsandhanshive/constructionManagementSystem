package bean

import (
	"fmt"
)

type Customer struct {
	userid   int
	password string
}

func CreateCustomer(userid int, password string) Customer {
	return Customer{userid: userid, password: password}
}
func (c Customer) getUserid() int {
	return c.userid
}
func (c Customer) setUserid(userid int) {
	c.userid = userid
}
func (c Customer) getPassword() string {
	return c.password
}
func (c Customer) setPassword(password string) {
	c.password = password
}
func (c Customer) String() string {
	return fmt.Sprintf("Customer [userid=%d password=%s]", c.userid, c.password)
}

//@Override
//public String toString() {
//	return "Customer [userid=" + userid + ", password=" + password + "]";
//}
