package user

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    string `json:"id"`
	Rank  int    `json:"rank"`
	State int    `json:"state"`
}

var driverName, dataSourceName string

func init() {
	driverName = "mysql"
	dataSourceName = "root:123456@tcp(127.0.0.1:3306)/battle?charset=utf8"
}

func Conn() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
	}
	return db
}

func RandStr(strlen int) string {
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = rand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = rand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}

func (user *User) CreateUser(uid string) {
	db := Conn()
	defer db.Close()

	stmt, err := db.Prepare(`INSERT user (id, rank, state) values (?,?,?)`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(uid, 0, 0)
	if err != nil {
		panic(err)
	}
}

func (user *User) UpdateRank(uid string, rk int) {
	db := Conn()
	defer db.Close()

	stmt, err := db.Prepare(`UPDATE user set rank=? where id=?`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(uid, rk)
	if err != nil {
		panic(err)
	}
}

func (user *User) UpdateState(uid string, s int) {
	db := Conn()
	defer db.Close()

	stmt, err := db.Prepare(`UPDATE user set state=? where id=?`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(uid, s)
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	var user User
// 	for i := 0; i < 100; i++ {
// 		str := RandStr(13)
// 		user.CreateUser(str)
// 	}
// }
