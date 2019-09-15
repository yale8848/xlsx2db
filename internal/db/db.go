// Create by Yale 2019/9/12 16:54
package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jmoiron/sqlx"
)

type tableInfo struct {
	ColumnName string
	ColumnType string
}
type DB interface {
	Connect(driverName, dataSourceName string) error
	GetDBNames() ([]string, error)
	GetTableNames(db string) ([]string, error)
	GetTabScheme(db, tb string) (map[string]string, error)
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

	clos := make([]string, 0)
	for k, _ := range params[0] {
		clos = append(clos, k)
	}
	_, e := d.eng.Table(db + "." + tb).Cols(clos...).Insert(params)

	fmt.Println(e)

	return nil

}
func (d *db) GetTabScheme(db, tb string) (map[string]string, error) {
	tbi := make([]tableInfo, 0)
	sql := `select COLUMN_NAME,column_type from information_schema.columns where table_schema = '%s'and TABLE_NAME = '%s'`
	err := d.eng.SQL(sql).Find(&tbi)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]string)
	for _, v := range tbi {
		mp[v.ColumnName] = v.ColumnType
	}
	return mp, nil
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
