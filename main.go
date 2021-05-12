package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("index.html"))

}

func main() {
	fileName := os.Args[1]

	someSlice := []byte("Uno due tre\n")

	err := tmpl.Execute(os.Stdout, someSlice)
	if err != nil {
		log.Fatal(err)
	}

	tre := "9"

	conTre, err := strconv.ParseInt(tre, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	natra, err := strconv.Atoi(tre)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conTre, natra)

	ioutil.WriteFile(fileName, someSlice, 0644)
}
