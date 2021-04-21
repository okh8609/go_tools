package cmd

import (
	"log"

	"github.com/okh8609/go_tools/internal/sql2struct"
	"github.com/spf13/cobra"
)

var DBType string
var Hostname string
var UserName string
var Password string
var Charset string
var DBName string
var TableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql轉換和處理",
	Long:  "sql轉換和處理",
	Run: func(cmd *cobra.Command, args []string) {
		info := sql2struct.DBInfo{
			DBType:   DBType,
			Hostname: Hostname,
			UserName: UserName,
			Password: Password,
			Charset:  Charset,
		}

		db := sql2struct.NewDBModel(&info)
		err := db.Connect()
		if err != nil {
			log.Fatalf("db.Connect err: %v", err)
		}
		result, err := db.GetTableMembers(DBName, TableName)
		if err != nil {
			log.Fatalf("db.GetTableMembers err: %v", err)
		}
		err = sql2struct.Generate(TableName, result)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.Flags().StringVarP(&DBType, "dbtpye", "t", "mysql", `資料庫類型`)
	sqlCmd.Flags().StringVarP(&Hostname, "hostname", "o", "127.0.0.1:3306", `伺服器名稱或IP`)
	sqlCmd.Flags().StringVarP(&UserName, "username", "u", "", `使用者名稱`)
	sqlCmd.Flags().StringVarP(&Password, "password", "p", "", `密碼`)
	sqlCmd.Flags().StringVarP(&Charset, "charset", "c", "utf8mb4", `編碼`)
	sqlCmd.Flags().StringVarP(&DBName, "dbname", "s", "", `資料庫名稱`)
	sqlCmd.Flags().StringVarP(&TableName, "tablename", "n", "", `表格名稱`)
}
