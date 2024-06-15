package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getDbConnect() *sql.DB {
	db, err := sql.Open("mysql", "xieruixiang:Iam@123456@tcp(139.196.105.2:3306)/xim")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func setConnect(db *sql.DB) {
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}

type UserSchema struct {
	Name string
	Age  int
}

func dump(data interface{}) {
	marshal, _ := json.Marshal(data)
	fmt.Println(string(marshal))
}

func query(db *sql.DB) {
	var (
		users []UserSchema
	)
	rows, err := db.Query("select name,age from user where id > ?", 0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user UserSchema
		err := rows.Scan(&user.Name, &user.Age)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	dump(users) //[{"Name":"name1","Age":1},{"Name":"name2","Age":2},{"Name":"name3","Age":3},{"Name":"name4","Age":4}]
}

func Prepare(db *sql.DB) {
	var user UserSchema
	prepare, err := db.Prepare("SELECT name,age FROM user WHERE id > ?")
	if err != nil {
		log.Fatal(err)
	}
	row := prepare.QueryRow("3")
	_ = row.Scan(&user.Name, &user.Age)
	dump(user) //{"Name":"name2","Age":2}
	row = prepare.QueryRow("1")
	_ = row.Scan(&user.Name, &user.Age)
	dump(user) //{"Name":"name2","Age":2}
}

func queryRow(db *sql.DB) {
	var (
		user UserSchema
	)
	row := db.QueryRow("select name,age from user where id > ?", 1)
	err := row.Scan(&user.Name, &user.Age)
	if err != nil {
		log.Fatal(err)
	}

	dump(user) //{"Name":"name2","Age":2}
}

func exec(db *sql.DB) {
	prepare, err := db.Prepare("insert into user (name,age) values (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	result, err := prepare.Exec("name5", 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId()) //5,<nil>
	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(affected) //1
}

func affair(db *sql.DB) {
	begin, err := db.Begin()
	if err != nil {
		fmt.Println("affair open failure")
	}
	prepare, err := begin.Prepare("insert into user (id,name,age) values (?,?,?)")
	if err != nil {
		fmt.Println("prepare failure")
	}
	result, err := prepare.Exec(996, "name996", 996)
	if err != nil {
		fmt.Println("exec failure")
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("get insertId failure")
	} else {
		fmt.Println("insertId is", insertId) //insertId is 996
	}
	//begin.Commit()
	err = begin.Rollback()
	if err != nil {
		fmt.Println("rollback failure")
	} else {
		fmt.Println("rollback success") //rollback success
	}

	row := db.QueryRow("select id from user where id = ?", 996)
	var id int
	_ = row.Scan(&id)
	if id <= 0 {
		fmt.Println("row is not found") //row is not found
	}
}

func dispose(db *sql.DB) {
	prepare, err := db.Prepare("select id from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = prepare.Query(100)
	fmt.Println(err) //nil

	row := prepare.QueryRow(100)
	var id int
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no Rows") //no Rows
	}
}

func disposeNullFail(db *sql.DB) {
	var deletedAt string
	row := db.QueryRow("select deleted_at from user where id = ?", 1)
	err := row.Scan(&deletedAt)
	if err != nil {
		log.Fatal(err) //sql: Scan error on column index 0, name "deleted_at": converting NULL to string is unsupported
	}
	fmt.Println(deletedAt)
}

func disposeNullSuccess(db *sql.DB) {
	var deletedAt sql.NullString
	row := db.QueryRow("select deleted_at from user where id = ?", 1)
	err := row.Scan(&deletedAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deletedAt.String)
}

func disposeNullSuccess2(db *sql.DB) {
	var deletedAt string
	row := db.QueryRow("select COALESCE(deleted_at,'') as deleted_at from user where id = ?", 1)
	err := row.Scan(&deletedAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deletedAt)
}
func disposeColumns(db *sql.DB) {
	row, err := db.Query("select COALESCE(deleted_at,'') as deleted_at from user where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	columns, err := row.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(columns) //[deleted_at]
}

func disUnKnowType(db *sql.DB) {
	value := new(sql.RawBytes)
	row := db.QueryRow("select COALESCE(deleted_at,'') as deleted_at from user where id = ?", 1)
	err := row.Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(*value))
}

func monitor(db *sql.DB) {
	stats := db.Stats()
	fmt.Println("因为超过最大存活时间被关闭的数量", stats.MaxLifetimeClosed)
	fmt.Println("因为超过最大空闲连接被关闭的数量", stats.MaxIdleClosed)
	fmt.Println("空闲连接因为超过空闲的最大生存周期被关闭的数量", stats.MaxIdleTimeClosed)
	fmt.Println("等待数量", stats.WaitCount)
	fmt.Println("等待时间", stats.WaitDuration)
	fmt.Println("打开的连接数", stats.OpenConnections)
}

func main() {
	connect := getDbConnect()
	setConnect(connect)

	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	monitor(connect)
}
