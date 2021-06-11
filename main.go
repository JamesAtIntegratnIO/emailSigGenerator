package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type User struct {
	Name        string `yaml:"Name"`
	Title       string `yaml:"Title"`
	Phone       string `yaml:"Phone"`
	Logo        string `yaml:"Logo"`
	CompanyName string `yaml:"CompanyName"`
	CompanyURL  string `yaml:"CompanyURL"`
}

const (
	Host = "localhost"
	Port = "8181"
)

var FileExists = false

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

const ErrorTemplate = `
<!DOCTYPE html>
<h3>
    <p><b>emailSigGenerator.yaml did not exist. 
	I have graciuosly thought of this and created a sample for you. 
	Please fill it out and reload this page</b></p>
<h3>
`

const SampleYaml = `
Name: "First Last"
Title: "Your Title"
Phone: "555-555-5555"
Logo: "Publicly available url to your logo Ex: https://tensure.io/icons/icon-144x144.png?v=669fd962b090ca24382d97e5c236b611"
CompanyName: "Your Company"
CompanyURL: "URL for your company Ex: https://tensure.io"

`

func createInputYaml() {
	f, err := os.Create("./emailSigGenerator.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(SampleYaml)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (u *User) getUserFromYaml() *User {
	_, err := os.Stat("./emailSigGenerator.yaml")
	if os.IsNotExist(err) {
		createInputYaml()
		FileExists = false
	} else {
		FileExists = true
	}
	yamlFile, err := ioutil.ReadFile("./emailSigGenerator.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, u)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return u
}

func getTemplate() string {
	if FileExists {
		return SignatureTemplate
	} else {
		return ErrorTemplate
	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	var user User
	user.getUserFromYaml()
	parsedTemplate, _ := template.New("SignatureTemplate").Parse(getTemplate())
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
