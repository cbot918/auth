package db

import (
	"auth/internal/util"
	"database/sql"
	"fmt"

	"auth/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	db    *sql.DB
	table string
}

func NewDB(driver string, user string, password string, host string, port int32, use string) *DB {
	obj := new(DB)

	// mysql string
	// cString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, use)

	// postgres string
	cString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, use)

	var err error
	obj.db, err = sql.Open(driver, cString)
	util.Checke(err, "db open failed")

	return obj
}

func (d *DB) SetTable(tableName string) {
	d.table = tableName
}

// func (d *DB) ExeQuery() {
// }

func (d *DB) CreateRow(name string, email string, pass string) {
	sql := fmt.Sprintf("insert into %s (name, email, password) values ('%s','%s','%s')", d.table, name, email, pass)
	_, err := d.db.Exec(sql)
	util.Checke(err, "CreateRow failed")
	util.Logg("CreateRow success")
}

func (d *DB) ReadRow(uid int32) {
	var user model.User

	sql := fmt.Sprintf(`select * from %s where id=%d`, d.table, uid)
	row := d.db.QueryRow(sql)
	err := row.Scan(&user.Uid, &user.Name, &user.Email, &user.Password)
	util.Checke(err, "row.Scan failed")
	util.Logg(user)
	util.Logg("ReadRow success")
}

// func (d *DB) ReadTable() {
// 	sql := fmt.Sprintf(`select * from %s`, d.table)
// 	rows, err := d.db.Query(sql)
// 	util.Checke(err, "Query failed")
// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err = rows.Scan(&user.uid, &user.name, &user.email, &user.password)
// 		util.Checke(err, "rows.Scan failed")
// 		users = append(users, user)
// 	}
// 	util.Logg(users)
// 	util.Logg("ReadTable success")
// }

func (d *DB) UpdateRow(target string, value string, where int32) {
	sql := fmt.Sprintf("UPDATE %s SET %s = '%s' where uid = %d", d.table, target, value, where)
	_, err := d.db.Exec(sql)
	util.Checke(err, "update row failed")
	util.Logg("UpdateRow success")
}

func (d *DB) DeleteRow(where int32) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE uid=%d", d.table, where)
	_, err := d.db.Exec(sql)
	util.Checke(err, "delete row failed")
	util.Logg("DeleteRow success")
}
