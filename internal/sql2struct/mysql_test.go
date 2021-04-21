package sql2struct_test

import "fmt"
import "github.com/okh8609/go_tools/internal/sql2struct"
import "testing"

/*
cd ~/okh8609/go_tools/internal/sql2struct
go test
*/
func TestMySQL(t *testing.T) {
	info := sql2struct.DBInfo{
		DBType:   "mysql",
		Hostname: "127.0.0.1",
		UserName: "kh",
		Password: "1234",
		Charset:  "utf8mb4",
	}

	db := sql2struct.NewDBModel(&info)
	db.Connect()
	result, _ := db.GetTableMembers("GoDB", "blog_tag")
	fmt.Print(result)
	sql2struct.Generate("blog_tag", result)

	// t.Log("TestMySQL PASS")
	// t.Fail()
	// t.Error("something wrong")
}
