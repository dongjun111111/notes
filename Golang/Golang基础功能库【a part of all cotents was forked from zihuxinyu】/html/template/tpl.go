package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t, err := template.ParseFiles("./tpl.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title string
	}{
		Title: "golang html template demo",
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

}
