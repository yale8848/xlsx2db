// Create by Yale 2019/9/12 16:54
package internal

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DB interface {
	Connect(driverName, dataSourceName string) error
	GetDBNames() ([]string, error)
	GetTableNames(db string) ([]string, error)
	GetTabScheme(db, tb string) ([]string, error)
	Insert(db, tb string, params []map[string]interface{}) error
}

type db struct {
	con *sqlx.DB
	eng *xorm.Engine
}

func (d *db) getNames(sql string) ([]string, error) {
	dbNames := make([]string, 0)
	err := d.eng.SQL(sql).Find(&dbNames)
	return dbNames, err
}
func (d *db) Insert(db, tb string, params []map[string]interface{}) error {

	sql := `Insert into %s.%s (%s) values(%s) `
	clos := make([]string, 0)
	vs := make([]string, 0)
	for k, _ := range params[0] {
		clos = append(clos, k)
		vs = append(vs, "?")
	}
	sqls := make([]string, 0)
	for i := 0; i < len(params); i++ {
		sqls = append(sqls, fmt.Sprintf(sql, db, tb, strings.Join(clos, ","), strings.Join(vs, ",")))
	}
	_, e := d.eng.Exec(sqls, &params)

	fmt.Println(e)

	return nil

}
func (d *db) GetTabScheme(db, tb string) ([]string, error) {
	sql := `select COLUMN_NAME from information_schema.columns where table_schema = '%s'and TABLE_NAME = '%s'`
	return d.getNames(fmt.Sprintf(sql, db, tb))
}
func (d *db) GetTableNames(db string) ([]string, error) {
	sql := `select TABLE_NAME from information_schema.columns where table_schema = '%s' `
	return d.getNames(fmt.Sprintf(sql, db))

}
func (d *db) GetDBNames() ([]string, error) {
	sql := `select table_schema from information_schema.columns group by table_schema`
	return d.getNames(sql)
}
func (d *db) Connect(driverName, dataSourceName string) error {
	eng, err := xorm.NewEngine(driverName, dataSourceName)

	if err != nil {
		return err
	}
	d.eng = eng

	return nil
}
