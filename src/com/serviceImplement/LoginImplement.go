package serviceImplement

import (
	"fmt"
)

func Login(userid int, password string) string {
	query := fmt.Sprintf("select userid,password,role from login where userid=%d and password= %v", userid, password)
	result, err := db.Query(query)

	if err != nil {
		fmt.Print(err)
		return ""
	}
	var role string

	for result.Next() {

		result.Scan(&userid, &password, &role)
	}
	return role
}

func ChangePass(userid int, password string, newpassword string) int {
	query := fmt.Sprintf("update login set password=%v where userid=%d and password=%v", newpassword, userid, password)
	_, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return userid
}
