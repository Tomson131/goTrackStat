package main

import (
	"bufio"
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

type StatusDesc struct {
	Post string `db:"postal_code"`
	Name string `db:"status_name"`
	Code string `db:"status_code"`
}

func main() {

	var xar = []BarCode{}
	var list = GetTable()
	for _, h := range list {
		xar = SetArray(xar, h)
	}
	var ld = GetStatDescs()
	for _, l := range ld {
		println(l.Name)
	}

	/*
		bytes, err := json.Marshal(xar)
		if err != nil {
			print("Can't serislize", xar)
		}
		print("var  = '%v'\n", string(bytes))
		WriteToFile("out.json", string(bytes))

	*/

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

func GetStatDescs() []StatusDesc {

	u := []StatusDesc{}
	conn, err := sqlx.Connect("mysql", "art:art2@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	err = conn.Select(&u, "SELECT postal_code,status_name,status_code FROM status_desc")
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
