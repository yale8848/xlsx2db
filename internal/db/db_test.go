// Create by Yale 2019/9/12 17:59
package db

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
	"time"
)

func getConnect() DB {
	b, _ := ioutil.ReadFile("db.info")
	d := db{}
	err := Connect("mysql", string(b))
	if err != nil {
		panic(err)
	}
	return &d
}
func TestDb_GetDBNames(t *testing.T) {
	db := getConnect()
	n, _ := GetDBNames()
	fmt.Sprintln(n)
}

func TestDb_GetTabScheme(t *testing.T) {
	db := getConnect()
	n, _ := GetTabScheme("db_test", "tb_test")
	fmt.Sprintln(n)

}
func TestDb_Insert(t *testing.T) {
	p := make([]map[string]interface{}, 0)

	n := time.Now().Format("2006-01-02 15:04:05")
	for i := 2 * 10; i < 2*10+10; i++ {
		mp := make(map[string]interface{})
		mp["test_number"] = strconv.Itoa(i)
		mp["test_name"] = strconv.Itoa(i)
		mp["test_time"] = n
		mp["test_f"] = "3.14"
		mp["test_d"] = "3.14"
		mp["test_date"] = n
		mp["test_timestamp"] = n
		p = append(p, mp)
	}

	db := getConnect()
	Insert("db_test", "tb_test", p)

}
