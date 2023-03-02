package bean

import (
	"fmt"
)

type Contractor struct {
	userid   int
	password string
}

func CreateContractor(userid int, password string) Contractor {
	return Contractor{userid: userid, password: password}
}
func (c Contractor) getUserid() int {
	return c.userid
}
func (c Contractor) setUserid(userid int) {
	c.userid = userid
}
func (c Contractor) getPassword() string {
	return c.password
}
func (c Contractor) setPassword(password string) {
	c.password = password
}
func (c Contractor) String() string {
	return fmt.Sprintf("Contractor [userid=%d, password=%s]", c.userid, c.password)
}
