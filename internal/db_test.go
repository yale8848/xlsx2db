// Create by Yale 2019/9/12 17:59
package internal

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
)

func getConnect() DB {
	b, _ := ioutil.ReadFile("db.info")
	d := db{}
	err := d.Connect("mysql", string(b))
	if err != nil {
		panic(err)
	}
	return &d
}
func TestDb_GetDBNames(t *testing.T) {
	db := getConnect()
	n, _ := db.GetDBNames()
	fmt.Sprintln(n)
}

func TestDb_GetTabScheme(t *testing.T) {
	db := getConnect()
	n, _ := db.GetTabScheme("db_test", "tb_test")
	fmt.Sprintln(n)

}
func TestDb_Insert(t *testing.T) {
	p := make([]map[string]interface{}, 0)

	for i := 10000; i < 10000+10; i++ {
		mp := make(map[string]interface{})
		mp["test_number"] = i
		mp["test_name"] = strconv.Itoa(i + 800)
		p = append(p, mp)
	}

	db := getConnect()
	db.Insert("db_test", "tb_test", p)

}
