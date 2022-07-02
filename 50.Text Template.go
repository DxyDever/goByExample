package main

import (
	"html/template"
	"os"
)

func main() {

	t1 := template.New("t1")

	t1, err := t1.Parse("Value is {{.}}\n")

	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)

	t1.Execute(os.Stdout, []string{
		"go",
		"rust",
		"c++",
		"c#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Done"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "dxy",
	})

	/*
		裁剪空格
		// 裁剪 content 前后的空格
		{{- content -}}

		// 裁剪 content 前面的空格
		{{- content }}

		// 裁剪 content 后面的空格
		{{ content -}}
	*/
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")

	t4.Execute(os.Stdout,
		[]string{
			"go",
			"rust",
			"c++",
			"c#",
		})
}
