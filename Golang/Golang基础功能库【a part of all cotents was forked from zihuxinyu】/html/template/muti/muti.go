package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t, err := template.ParseFiles("tpl/index.html", "tpl/public/header.html", "tpl/public/footer.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title string
	}{
		Title: "load common template",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
