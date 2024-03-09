package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func render(writer http.ResponseWriter, tpl *template.Template) {
	err := tpl.ExecuteTemplate(writer, tpl.Name(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func AutoRegisterAndRender() {
	t, err := template.ParseGlob("template/*.xtl")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(t.Templates()))
	for _, tp := range t.Templates() {
		tpl := tp
		if strings.Contains(tpl.Name(), ".xtl") {
			continue
		}
		fmt.Println("/" + tpl.Name())
		http.HandleFunc("/"+tpl.Name(), func(writer http.ResponseWriter, request *http.Request) {
			render(writer, tpl)
		})
	}
}

func main() {
	AutoRegisterAndRender()
	log.Fatal(http.ListenAndServe(":9503", nil))
}
