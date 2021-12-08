# Feature Flag Demo

## Run instructions

### Unleash

1. Start Unleash
`docker-compose -f docker-compose.unleash.yaml up -d`

2. Config Unleash in the address http://localhost:4242

User: admin
Password: unleash4all

* Create an api key and update the project file settings.yaml
* Create and flag and update the project file settings.yaml

3. Run program

**Windows**

Open a command prompt on directory build a run the following command:
`feature-flag-demo.exe unleash --config settings.yaml`

**Linux**

Open a command prompt on directory build a run the following command:
`feature-flag-demo unleash --config settings.yaml`

**Run with go**
`go run . unleash --config settings.yaml`

### GitLab

1. Start Unleash
`docker-compose -f docker-compose.gitlab.yaml up -d`

2. Run program

**Windows**

Open a command prompt on directory build a run the following command:
`feature-flag-demo.exe gitlab --config settings.yaml`

**Linux**

Open a command prompt on directory build a run the following command:
`feature-flag-demo gitlab --config settings.yaml`

**Run with go**
`go run . gitlab --config settings.yaml`
