package secondary

import (
	"bytes"
	"html/template"
	"os"
)

func ProcessTemplate(html []byte, data interface{}) ([]byte, error) {
	tmpl, err := template.New("template").Parse(string(html))
	if err != nil {
		return nil, err
	}

	var processedHTML bytes.Buffer
	err = tmpl.Execute(&processedHTML, data)
	if err != nil {
		return nil, err
	}

	return processedHTML.Bytes(), nil
}

func GetTemplate(name string) []byte {
	file, err := os.Open("./web/templates/" + name + ".html")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		panic(err)
	}

	return bs
}
