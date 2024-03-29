package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Repo struct {
	Reposhort, RepoHttp, Branch, Path string
}

func WriteString(file string, string string, perm os.FileMode) error {
	data := []byte(string)
	err := ioutil.WriteFile(file, data, perm)
	return err
}

func Mkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		return nil
	}
	return fmt.Errorf("Problem in pkg.Mkdir. Could not create: %s\n", path)
}

func Rmdir(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		os.RemoveAll(path)
	}
}

func Sample() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}

func TmpCreate(templateText string, r interface{}) bytes.Buffer {

	var b bytes.Buffer

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title": strings.Title,
	}

	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(&b, r)
	//b.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
	return b
}

func Sample2() {

	var b bytes.Buffer
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title": strings.Title,
	}

	// A simple template definition to test our function.
	// We print the input text several ways:
	// - the original
	// - title-cased
	// - title-cased and then printed with %q
	// - printed with %q and then title-cased.
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
`

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(&b, "the go programming language")
	b.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}

func GetPWD() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return dir, err
	}
	return dir, err
}

func SliceIndex(limit int, hop int, predicate func(i int) bool) int {
	if limit == 0 {
		return -1
	}

	for i := limit - 1; i >= 0; i-- {
		if predicate(i) {
			if i+hop < limit {
				return i + hop
			}
		}
	}
	return -1
}

func GithubRepo(dir string) (string, bool) {
	xs := strings.Split(dir, "/")
	idx := SliceIndex(len(xs), 2, func(i int) bool { return xs[i] == "github.com" })
	if idx != -1 {
		return xs[idx], true
	}
	return "", false
}

func GithubUser(dir string) (string, bool) {
	xs := strings.Split(dir, "/")
	idx := SliceIndex(len(xs), 1, func(i int) bool { return xs[i] == "github.com" })
	if idx != -1 {
		return xs[idx], true
	}
	return "", false
}
