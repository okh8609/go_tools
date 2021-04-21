package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBInfo struct {
	DBType   string
	Hostname string
	UserName string
	Password string
	Charset  string
}

type DBModel struct {
	DBEngin *sql.DB
	DBInfo  *DBInfo
}

type TableMember struct {
	Name       string
	DataType   string
	IsNullable string
	IsKey      string
	DataType2  string
	Comment    string
}

func (m *TableMember) String() string {
	return fmt.Sprintf("{Name:%s, DataType:%s, IsNullable:%s, IsKey:%s, DataType2:%s, Comment:%s}",
		m.Name, m.DataType, m.IsNullable, m.IsKey, m.DataType2, m.Comment)
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	ss := fmt.Sprintf(s, m.DBInfo.UserName, m.DBInfo.Password, m.DBInfo.Hostname, m.DBInfo.Charset)
	m.DBEngin, err = sql.Open(m.DBInfo.DBType, ss)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetTableMembers(dbName string, tableName string) ([]*TableMember, error) {
	query := "select COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT" +
		" from INFORMATION_SCHEMA.COLUMNS where table_schema = ? AND table_name = ?;"
	rows, err := m.DBEngin.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("無資料")
	}
	defer rows.Close()

	var tableMembers []*TableMember
	for rows.Next() {
		var tmb TableMember
		err := rows.Scan(&tmb.Name, &tmb.DataType, &tmb.IsKey, &tmb.IsNullable, &tmb.DataType2, &tmb.Comment)
		if err != nil {
			return nil, err
		}
		tableMembers = append(tableMembers, &tmb)
	}
	return tableMembers, nil
}
