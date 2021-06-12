package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"gopkg.in/yaml.v2"
)

type EmailSigData struct {
	Name        string `yaml:"Name"`
	Title       string `yaml:"Title"`
	Phone       string `yaml:"Phone"`
	LogoURL     string `yaml:"LogoURL"`
	CompanyName string `yaml:"CompanyName"`
	CompanyURL  string `yaml:"CompanyURL"`
}

const SignatureTemplate = `
<!DOCTYPE html>
<table>
    <tbody>
        <tr>
            <td><img src={{.LogoURL}}>
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

const FormTemplate = `
<h1>Email Signature Generator</h1>
<form method="POST">
	<label>Name:</label><br />
	<input type="text" name="Name"><br />
	<label>Title:</label><br />
	<input type="text" name="Title"><br />
	<label>Phone:</label><br />
	<input type="text" name="Phone"><br />
	<label>Logo URL:</label><br />
	<input type="text" name="LogoURL"><br />
	<label>Company Name:</label><br />
	<input type="text" name="CompanyName"><br />
	<label>Company URL:</label><br />
	<input type="text" name="CompanyURL"><br />
	<input type="submit">
</form>
`
const WelcomeMesage = `
  If your browser failed to open:
  Open your browser and navigate to http://localhost:8181
`

const Buttons = `
<html>
	<head>
		<title>Email Signature Generator</title>
	</head>
	<body>
		<h1>Email Signature Generator</h1>
		<button onclick="window.location.href='http://localhost:8181/generate';">
		Generate From Yaml
		</button>
		<button onclick="window.location.href='http://localhost:8181/form';">
		Generate From From
		</button>
	</body>
</html>
`

func getNewTemplate(templateName string) template.Template {
	tmpl := template.Must(template.New("Template").Parse(templateName))
	return *tmpl
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (e *EmailSigData) getSigDataFromYaml() *EmailSigData {
	yamlFile, err := ioutil.ReadFile("./emailSigGenerator.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, e)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return e
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := getNewTemplate(Buttons)
	tmpl.Execute(w, nil)
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	var data EmailSigData
	data.getSigDataFromYaml()
	tmpl := getNewTemplate(SignatureTemplate)
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func formFlow(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		tmpl := getNewTemplate(FormTemplate)
		tmpl.Execute(w, nil)
		return
	}
	formData := EmailSigData{
		Name:        r.FormValue("Name"),
		Title:       r.FormValue("Title"),
		Phone:       r.FormValue("Phone"),
		LogoURL:     r.FormValue("LogoURL"),
		CompanyName: r.FormValue("CompanyName"),
		CompanyURL:  r.FormValue("CompanyURL"),
	}
	tmpl := getNewTemplate(SignatureTemplate)
	tmpl.Execute(w, formData)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/generate", renderTemplate)
	http.HandleFunc("/form", formFlow)
	log.Println(WelcomeMesage)
	err := http.ListenAndServe("localhost:8181", nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}

func init() {
	openbrowser("http://localhost:8181")
}
