// Create by Yale 2019/9/12 16:54
package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type TableInfo struct {
	ColumnName string `json:"columnName"`
	ColumnType string `json:"columnType"`
}
type DB interface {
	Connect(driverName, dataSourceName string) error
	GetDBNames() ([]string, error)
	GetTableNames(db string) ([]string, error)
	GetTabScheme(db, tb string) ([]TableInfo, error)
	Insert(db, tb string, params []map[string]interface{}) error
}

type xdb struct {
	eng *xorm.Engine
}

func NewDB() DB {
	return &xdb{}
}

func (d *xdb) getNames(sql string) ([]string, error) {
	dbNames := make([]string, 0)
	err := d.eng.SQL(sql).Find(&dbNames)
	return dbNames, err
}
func (d *xdb) Insert(db, tb string, params []map[string]interface{}) error {

	clos := make([]string, 0)
	for k, _ := range params[0] {
		clos = append(clos, k)
	}
	_, e := d.eng.Table(db + "." + tb).Cols(clos...).Insert(params)
	return e

}
func (d *xdb) GetTabScheme(db, tb string) ([]TableInfo, error) {
	tbi := make([]TableInfo, 0)
	sql := `select column_name,column_type from information_schema.columns where table_schema = ? and TABLE_NAME = ? `
	err := d.eng.SQL(sql, db, tb).Find(&tbi)
	if err != nil {
		return nil, err
	}
	return tbi, nil
}
func (d *xdb) GetTableNames(db string) ([]string, error) {
	sql := `select distinct TABLE_NAME from information_schema.columns where table_schema = '%s' `
	return d.getNames(fmt.Sprintf(sql, db))

}
func (d *xdb) GetDBNames() ([]string, error) {
	sql := `select distinct table_schema from information_schema.columns group by table_schema`
	return d.getNames(sql)
}
func (d *xdb) Connect(driverName, dataSourceName string) error {
	eng, err := xorm.NewEngine(driverName, dataSourceName)

	if err != nil {
		return err
	}
	d.eng = eng

	return nil
}
