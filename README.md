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

3. Run the program

**Windows**

Open a command prompt on directory build a run the following command:
`feature-flag-demo.exe unleash --config settings.yaml`

**Linux**

Open a command prompt on directory build a run the following command:
`feature-flag-demo unleash --config settings.yaml`

**Run with go**
`go run . unleash --config settings.yaml`

### GitLab (Only linux)

1. Create an directory for gitlab docker volumes
`mkdir gitlab`

2. Create an environment variable for the gitlab docker volumes directory
`export GITLAB_HOME=./gitlab`

3. Start GitLab and wait a couple of minutes
`docker-compose -f docker-compose.gitlab.yaml up -d`

You can check the logs with:
`docker logs feature-flag-demo_gitlab_1 -f`

4. Modify the file ./gitlab/config/gitlab.rb
   * Uncomment the line external_url and alter its value for https://gitlab.example.com:8443
   * Uncomment the line gitlab_rails['gitlab_shell_ssh_port'] = and alter it's value for 822
   * Enter docker container with `docker exec -it feature-flag-demo_gitlab_1 bash` and run the command `gitlab-ctl reconfigure`

5. Get the generate password for the root use in the file ./gitlab/config/initial_root_password

6. In the browser access the address https://localhost:8433 with root user

7. In gitlab
    * Create a project
    * Create an configure a flag and an environment

8. Run the program

**Windows**

Open a command prompt on directory build a run the following command:
`feature-flag-demo.exe gitlab --config settings.yaml`

**Linux**

Open a command prompt on directory build a run the following command:
`feature-flag-demo gitlab --config settings.yaml`

**Run with go**
`go run . gitlab --config settings.yaml`
