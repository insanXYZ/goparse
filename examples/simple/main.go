package main

import (
	"fmt"
	"net/http"

	"github.com/insanXYZ/goparse"
)

const Port = ":8080"

func main() {

	template := goparse.NewTemplates("views/*.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.ExecuteTemplate(w, "index.html", nil)
	})

	fmt.Printf("server running , port %s", Port)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		panic(err.Error())
	}

}
