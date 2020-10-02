package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type BarCode struct {
	List    []StatusBase
	Barcode string
}

type StatusBase struct {
	Event_id      int    `db:"event_id"`
	Glob_barcode  string `db:"glob_barcode"`
	Barcode       string `db:"barcode"`
	Entry_date    string `db:"entry_date"`
	Status_date   string `db:"status_date"`
	Status_zip    string `db:"status_zip"`
	Status_place  string `db:"status_place"`
	Status_event  string `db:"status_event"`
	Status_weight string `db:"status_weight"`
	Status_add    string `db:"status_add"`
	Additional    string `db:"additional"`
	Country_org   string `db:"country_org"`
	Country_dest  string `db:"country_dest"`
	Barcode_dest  string `db:"barcode_dest"`
	Orgstatus     string `db:"orgstatus"`
}

func main() {

	var xar = []BarCode{}
	var list = GetTable()
	for _, h := range list {
		xar = SetArray(xar, h)
	}
	bytes, err := json.Marshal(xar)
	if err != nil {
		print("Can't serislize", xar)
	}
	print("var  = '%v'\n", string(bytes))
	WriteToFile("out.json", string(bytes))
	if false {
		/*conn, err := sqlx.Connect("mysql", "art:art2@tcp(localhost:3306)/test")
		if err != nil {
			panic(err)
		}
		conn.MustExec(schema)
		res, err := conn.Exec("INSERT INTO users2 (name) VALUES(\"Peter\")")
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created user with id:%d", id)
		var user User
		err = conn.Get(&user, "select * from users2 where id=?", id)
		if err != nil {
			panic(err)
		}
		_, err = conn.Exec("UPDATE users set name=\"John\" where id=?", id)
		if err != nil {
			panic(err)
		}
		_, err = conn.Exec("DELETE FROM users where id=?", id)
		if err != nil {
			panic(err)
		}*/
	}
	scanner()
}

func SetArray(ar []BarCode, stat StatusBase) []BarCode {
	if Contains(ar, stat) {
		ar = SetContains(ar, stat)
	} else {
		ar = AddBarar(ar, stat)
	}
	return ar
}

func GetTable() []StatusBase {

	u := []StatusBase{}
	conn, err := sqlx.Connect("mysql", "art:art2@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	err = conn.Select(&u, "SELECT * FROM status_base")
	if err != nil {
		panic(err)
	}

	return u
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
