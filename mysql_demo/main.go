package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//数据库信息
	// dsn := "root:root_wking@tcp(192.168.88.88:3306)/qingcheng_goods"
	dsn := "root:root@tcp(localhost:3306)/qingcheng_goods"

	//链接数据库
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("open %s failed, err: %v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("链接数据库失败！！")
	} else {
		fmt.Println("链接数据库成功！！")
	}

}
