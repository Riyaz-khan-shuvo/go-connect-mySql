package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:123456Ri@tcp(127.0.0.1:3306)/hosting_db")
	if err != nil {
		panic(err.Error())
	}

	// insert , err:= db.Query(INSERT INTO `loruki` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, 'Riyaz Hossain', 'Riyaz dot com', 'mdriyaz5965@gmail.com', '1');)

	// insert, err := db.Query("INSERT INTO `loruki` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, 'Riyaz Hossain', 'riyaz dot com', 'mdriyaz5965@gmail.com', '1');")

	if err != nil {
		panic(err.Error())
	}

	// INSERT INTO `loruki` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, 'Riyaz Hossain', 'Riyaz dot com', 'mdriyaz5965@gmail.com', '1');

	// defer the close till after the main function has finished
	// executing
	fmt.Println("Database Connected Successful")
	// defer db.Close()
	// defer insert.Close()
}

func main() {
	fmt.Println("I am working")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/features", featuresPage)
	http.HandleFunc("/docs", docsPage)
	http.HandleFunc("/request", sendData)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":9000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	pToTem, err := template.ParseFiles("./Template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)
}

func featuresPage(w http.ResponseWriter, r *http.Request) {
	pToTem, err := template.ParseFiles("./Template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem, err = pToTem.ParseFiles("./wPage/feature.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)
}

func docsPage(w http.ResponseWriter, r *http.Request) {
	pToTem, err := template.ParseFiles("./Template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem, err = pToTem.ParseFiles("./wPage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)
}

func sendData(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Data send successful!!!")

	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	data := "INSERT INTO `loruki` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"

	sql := fmt.Sprintf(data, name, company, email)

	insert, err := db.Query(sql)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer insert.Close()
}
