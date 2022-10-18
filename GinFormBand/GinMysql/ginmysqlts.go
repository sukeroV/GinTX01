package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sukerov/myFunctions"
)

var db *sql.DB

func main() {
	connStr := "root:root@tcp(127.0.0.1:3306)/ginmysqlts"
	db = myFunctions.MySQL_Connect(connStr)

	//var u1 []name
	u1 := make([]name, 0)
	n, _ := db.Query(`select name,age from person where id<250;`)
	for n.Next() {
		u2 := name{}
		n.Scan(&u2.Name, &u2.Age)
		u1 = append(u1, u2)
	}
	fmt.Printf("%v", u1)
}

type name struct {
	Name string
	Age  int
}
