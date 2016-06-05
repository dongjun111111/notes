package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	// 声明模板内容
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// 创建一个新的模板，并且载入内容
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	// 定义传入到模板的数据，并在终端打印
	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}
	err = t.Execute(os.Stdout, data)
	check(err)

	fmt.Println()

	// 定义Items为空的数据
	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}
	err = t.Execute(os.Stdout, noItems)
	check(err)

}
