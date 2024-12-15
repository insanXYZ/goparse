package goparse

import (
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	TmpDirName = ".goparse-tmp"
	wg         = new(sync.WaitGroup)
)

func NewTemplates(pattern string) *template.Template {

	err := os.MkdirAll(TmpDirName, os.ModePerm)
	if err != nil {
		returnPanic(ErrCreateTmpDir)
	}

	defer os.RemoveAll(TmpDirName)

	path, singlePattern, err := splitSlash(pattern)
	if err != nil {
		returnPanic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = handlePath(path, singlePattern, wg)
		if err != nil {
			returnPanic(err)
		}
	}()

	wg.Wait()

	return template.Must(template.ParseGlob(filepath.Join(TmpDirName, singlePattern)))
}

func splitSlash(pattern string) (string, string, error) {
	splits := strings.Split(pattern, "/")

	if len(splits) < 2 {
		return "", "", ErrInvalidPattern
	}

	return filepath.Join(splits[:len(splits)-1]...), splits[len(splits)-1], nil
}

func handlePath(path, pattern string, wg *sync.WaitGroup) error {

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, v := range entries {

		fullpath := filepath.Join(path, v.Name())

		if v.IsDir() {
			wg.Add(1)
			go func(fp string) {
				defer wg.Done()
				err := handlePath(fp, pattern, wg)
				if err != nil {
					returnPanic(err)
				}
			}(fullpath)
		} else if m, _ := filepath.Match(pattern, v.Name()); m {
			wg.Add(1)
			go func(fp, fname string) {
				defer wg.Done()
				err := handleEntry(fp, fname)
				if err != nil {
					returnPanic(err)
				}
			}(fullpath, v.Name())
		}
	}

	return nil
}

func handleEntry(path, filename string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	replacedTemplate := replaceTemplate(b)
	return createTemplate(filename, replacedTemplate)
}

func createTemplate(filename, value string) error {

	file, err := os.Create(filepath.Join(TmpDirName, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write([]byte(value))
	return err
}

func replaceTemplate(b []byte) string {
	re := regexp.MustCompile(`{{\s*template\s*"[^/"]*/([^"]+)"\s*}}`)

	return re.ReplaceAllString(string(b), `{{ template "$1" }}`)
}
