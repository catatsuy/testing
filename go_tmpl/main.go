package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	fmap := template.FuncMap{
		"con": func(str string) string {
			if str == "abc" {
				return "aaaa"
			}
			return "bbb"
		},
	}

	v := map[string]interface{}{
		"Header": "!!! へっだー !!!",
		"Footer": "zzz ふったー zzz",
	}

	fmt.Println("\n===foo===")
	t := template.Must(template.New("foo.html").Funcs(fmap).ParseGlob("template/*"))
	t.Execute(os.Stdout, v)

	fmt.Println("\n===bar===")
	t.ExecuteTemplate(os.Stdout, "bar.html", v)
}
