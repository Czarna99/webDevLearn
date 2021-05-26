package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseGlob("testFiles/*.gohtml")
	if err != nil {
		log.Fatal(err)

	}

	err = tpl.ExecuteTemplate(os.Stdout, "szop.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "junior.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "best.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("File.txt")
	if err != nil {

		log.Fatal(err)

		defer nf.Close()
	}
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}

}
