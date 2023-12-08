package model

import (
	"database/sql"
	"fmt"

	"web/utils"

	_ "github.com/go-sql-driver/mysql"
)

var dbinstance *sql.DB

type User struct {
	User_name string
	User_code string
	User_id   string
	Password  string
}

func init() {
	var err error
	host, port, dbname, username, Password := utils.GetConfigInfo()
	dsn := username + ":" + Password + "@tcp(" + host + ":" + port + ")/" + dbname
	fmt.Println("dsn:", dsn)
	dbinstance, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = dbinstance.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}

}
func GetUserInfo(id string) []User {

	u := User{}
	row := make([]User, 0)
	sql := "select password,user_code,user_id,user_name from user where user_id= ?"
	stmt, err := dbinstance.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	result, err := stmt.Query(id)
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		err := result.Scan(&u.Password, &u.User_code, &u.User_id, &u.User_name)
		if err != nil {
			// 打印错误信息
			fmt.Println(err)
			// 抛出错误信息，阻止程序继续运行
			panic(err)
		}
		row = append(row, u)
	}
	return row
}
