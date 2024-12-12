package goparse

import (
	"html/template"
	"io/fs"
	"os"
)

const (
	TMP_DIR_NAME = ".goparse-tmp"
)

func NewTemplates(dir string) *template.Template {

	removeLastSlash(&dir)
	ents, err := os.ReadDir(dir)
	if err != nil {
		ReturnPanic(ErrDirDoesntExist(dir))
	}

	err = os.Mkdir(TMP_DIR_NAME, 0666)
	if err != nil {
		ReturnPanic(ERR_CREATE_TMP_DIR)
	}

	defer os.Remove(TMP_DIR_NAME)

	handleEntries(dir, ents)

	return nil

}

func removeLastSlash(dir *string) {

	d := *dir

	if string(d[len(d)-1]) == "/" {
		*dir = d[:len(d)-1]
	}
}

func handleEntries(rootDir string, entries []fs.DirEntry) error {

	for _, v := range entries {
		if v.IsDir() {

			dir := rootDir + "/" + v.Name()

			ents, err := os.ReadDir(dir)
			if err != nil {
				return err
			}

			handleEntries(dir, ents)
		}

	}

}

//
// func main() {
// 	text := `{{ template "views/index.html" }} dan {{template"views/home.html"}} dan {{template "partials/header.html"}}`
//
// 	re := regexp.MustCompile(`{{\s*template\s*"[^/"]*/([^"]+)"\s*}}`)
//
// 	result := re.ReplaceAllString(text, `{{ template "$1" }}`)
//
// 	fmt.Println(result)
// }
