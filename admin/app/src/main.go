package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func dbConnect() *sql.DB {
        connection := "user=pguser dbname=application2 password=pgpass host=db sslmode=disable"
        db, err := sql.Open("postgres", connection)
        if err != nil {
                panic(err.Error())
        }
        return db
}

type Customer struct {
        Id int
        Name string
        Email string
}

var temp = template.Must(template.ParseGlob("src/templates/*.html"))

func main() {
		http.HandleFunc("/", index)
        http.ListenAndServe(":8082", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	db := dbConnect()
	selectAll, err := db.Query("select * from customers")
	if err != nil {
			panic(err.Error())
	}

	c := Customer{}
	customers := []Customer{}

	for selectAll.Next(){
			var id int
			var name, email string
			err := selectAll.Scan(&id, &name, &email)
			if err != nil {
					panic(err.Error())
			}

			c.Name = name
			c.Email = email

			customers = append(customers, c)

	}

	temp.ExecuteTemplate(w, "Index", customers)
	defer db.Close()
}