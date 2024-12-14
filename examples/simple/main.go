package main

import (
	"fmt"
	"io"
	"net/http/httptest"

	"github.com/insanXYZ/goparse"
)

func main() {
	t := goparse.NewTemplates("views/*.html")

	recorder := httptest.NewRecorder()
	err := t.ExecuteTemplate(recorder, "header.html", nil)
	if err != nil {
		panic(err.Error())
	}

	b, err := io.ReadAll(recorder.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(b))
}
