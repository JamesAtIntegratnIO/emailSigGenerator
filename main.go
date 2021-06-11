package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

const (
	Host = "localhost"
	Port = "8181"
)

type User struct {
	Name        string `yaml:"Name"`
	Title       string `yaml:"Title"`
	Phone       string `yaml:"Phone"`
	Logo        string `yaml:"Logo"`
	CompanyName string `yaml:"CompanyName"`
	CompanyURL  string `yaml:"CompanyURL"`
}

const SignatureTemplate = `
<!DOCTYPE html>
<table>
    <tbody>
        <tr>
            <td><img src={{.Logo}}>
            </td>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
            <td>
                {{.Name}}<br>
                <b>{{.Title}}</b><br>
                <a href={{.CompanyURL}} target="_blank"><span zeum4c4="PR_2_0" data-ddnwab="PR_2_0"
                        aria-invalid="grammar" class="Lm ng">{{.CompanyName}}</span></a><br>
                Phone: {{.Phone}}<br>
            </td>
        </tr>
    </tbody>
</table>
`

func (u *User) getUserFromYaml() *User {
	yamlFile, err := ioutil.ReadFile("./input.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, u)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return u
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	var user User
	user.getUserFromYaml()

	parsedTemplate, _ := template.New("SignatureTemplate").Parse(SignatureTemplate)
	err := parsedTemplate.Execute(w, user)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func main() {
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}
