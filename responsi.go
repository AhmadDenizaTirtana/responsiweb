package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func utama(res http.ResponseWriter, req *http.Request) {
	halaman, _ := template.New("tmp1").Parse(html)
	data := map[string]string{
		"nama": "Deni",
	}
	halaman.Execute(res, data)
}

func connect() *sql.DB {
	var db, err = sql.Open("mysql", "root@(localhost:3306)/db1")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return db
}

const html = `
<html>
<body>
selamat pagi {{.nama}}
<form method="post" action="/input_data">
			id : <br>
			<input type="text" name="id"><br>
			nama : <br>
			<input type="text" name="nama"><br>
			<input type="submit" value="input data">
			</form>
</body>
</html>
			`

func input_data(res http.ResponseWriter, req *http.Request) {
	var db = connect()
	defer db.Close()

	id := req.FormValue("id")
	name := req.FormValue("nama")

	var err error

	_, err = db.Exec("insert into data values (?, ?)", id, name)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	http.HandleFunc("/", utama) // set router
	http.HandleFunc("/input_data", input_data)
	fmt.Println("starting web server at http://localhost:9090/")
	http.ListenAndServe(":9090", nil) // set listen port

}
