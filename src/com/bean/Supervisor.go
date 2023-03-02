package bean

import (
	"fmt"
)

type Supervisor struct {
	userid   int
	password string
	sitename string
}

func CreateSupervisor(userid int, password string, sitename string) Supervisor {
	return Supervisor{userid: userid, password: password, sitename: sitename}
}
func (s Supervisor) getUserid() int {
	return s.userid
}
func (s Supervisor) setUserid(userid int) {
	s.userid = userid
}
func (s Supervisor) getPassword() string {
	return s.password
}
func (s Supervisor) setPassword(password string) {
	s.password = password
}
func (s Supervisor) getSitename() string {
	return s.sitename
}
func (s Supervisor) setSitename(sitename string) {
	s.sitename = sitename
}
func (s Supervisor) String() string {
	return fmt.Sprintf("Supervisor [userid=%d, password=%s ,sitename=%s]", s.userid, s.password, s.sitename)
}
