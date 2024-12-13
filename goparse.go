package goparse

import (
	"html/template"
	"io"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

var (
	TMP_DIR_NAME = ".goparse-tmp"
	HTML_EXT     = []string{"html", "gohtml"}
)

func NewTemplates(dir string) *template.Template {

	trimSlash(&dir)
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

	return template.Must(template.ParseGlob(TMP_DIR_NAME))

}

func trimSlash(dir *string) {

	d := *dir

	if string(d[len(d)-1]) == "/" {
		*dir = d[:len(d)-1]
	}
}

func handleEntries(rootDir string, entries []fs.DirEntry) error {

	for _, v := range entries {
		fullPath := rootDir + "/" + v.Name()

		if v.IsDir() {
			ents, err := os.ReadDir(fullPath)
			if err != nil {
				return err
			}
			handleEntries(fullPath, ents)
		} else {
			handleEntry(fullPath, v.Name())
		}

	}

	return nil

}

func handleEntry(fullpath, filename string) {
	if isValidHtml(getExtensionFile(filename)) {
		parserHtml(fullpath)
	}
}

func isValidHtml(extension string) bool {
	valid := false

	for _, v := range HTML_EXT {
		if v == extension {
			valid = true
		}
	}

	return valid
}

func getExtensionFile(filename string) string {
	split := strings.Split(filename, ".")
	return split[len(split)-1]
}

func parserHtml(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	replacedhtml := replaceHtmlTemplate(b)

}

func replaceHtmlTemplate(b []byte) string {
	re := regexp.MustCompile(`{{\s*template\s*"[^/"]*/([^"]+)"\s*}}`)

	return re.ReplaceAllString(string(b), `{{ template "$1" }}`)
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
