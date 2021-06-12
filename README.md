# Email Signature Generator

This is a quick tool I built that will generate a email signature based on input of a yaml file.
## Usage

### From the Binary
1. Grab the latest binary for your Operating System from [here](https://github.com/JamesAtTensure/emailSigGenerator/releases)
2. Extract it into its own folder
3. Edit `emailSigGenerator.yaml` and populate the fields
4. Run `emailSigGenerator`
5. If your browser doesn't open on its own. Open your browser and navigate to http://localhost:8181/
6. You will be given the option to generate your signature from the yaml file that was next to the executable. Or you can use the form and populate the data there
7. You can then copy the signature to your email client of choice. 
#### Gotchas
This app is currently unsigned. This means when you run it the first time you will get a warning and you won't be able to continue. On mac you will need to click the :question: mark and follow the instructions to allow it and then run it again.
### From Source
1. Clone the repository
2. Have Go [setup](https://golang.org/doc/install) 
3. Download the dependencies `go mod download`
4. Edit `emailSigGenerator.yaml`
5. Run the app `go run main.go`
6. You will be given the option to generate your signature from the yaml file that was next to the executable. Or you can use the form and populate the data there

This service is also available with curl
```shell
curl --request POST \
  --url http://localhost:8181/form \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data 'Name=First Last' \
  --data 'Title=Awesome Sauce' \
  --data Phone=555-555-5555 \
  --data 'Logo=https://tensure.io/icons/icon-144x144.png?v=669fd962b090ca24382d97e5c236b611' \
  --data CompanyName=Tensure \
  --data CompanyURL=https://tensure.io
  ```