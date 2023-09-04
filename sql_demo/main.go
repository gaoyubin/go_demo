package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	// "github.com/garyburd/redigo/redis"
	// "github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
)

var db *sql.DB

type User struct {
	realname string
	password string
	idx      int
}

func init() {
	var err error
	dsn := "user:password@tcp(127.0.0.1:3306)/mybatis"
	// dsn := "root:admin@tcp(127.0.0.1:3306)/mybatis?charset=utf8"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println()
}

// func init(){

// }
func prepareQueryDemo() {
	sqlStr := "select * from  user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare err:", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(2)
	if err != nil {
		fmt.Println("query err:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.idx, &u.realname, &u.password)
		if err != nil {
			fmt.Println("scan err", err)
			return
		}
		fmt.Println("show db", u)

	}
}
func main() {
	// prepareQueryDemo()
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	// redis.DialConnectTimeout

	if err != nil {
		// fmt.Println("conn err", err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println(conn)

	_, err = conn.Do("auth", "tl")
	if err != nil {
		panic(err)
	}

	_, err = conn.Do("set", "id", 1001)
	if err != nil {
		panic(err)
	}
	str, err := redis.String(conn.Do("get", "id"))
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
