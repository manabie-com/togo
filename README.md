## Environment
```
Ubuntu 18.04.2 LTS

$ java -version
openjdk version "11.0.11" 2021-04-20
OpenJDK Runtime Environment AdoptOpenJDK-11.0.11+9 (build 11.0.11+9)
OpenJDK 64-Bit Server VM AdoptOpenJDK-11.0.11+9 (build 11.0.11+9, mixed mode)

$ ansible --version
ansible 2.9.16
```

## Compile
```
Update url, username, password in src/main/resources/application.properties if any
manabie$ mvn clean install
```

## Run
```
manabie/target$ java -jar manabie-0.0.1-SNAPSHOT.jar
```

## Test
```
Update ansible_user, ansible_ssh_pass, ansible_sudo_pass, app_url in manabie/ansible/hosts.yml if any
manabie/ansible$ ansible-playbook -i hosts.yml test.yml
```

## Deploy
```
Update ansible_user, ansible_ssh_pass, ansible_sudo_pass, app_url, install_path, mvn_path, mysql in manabie/ansible/hosts.yml if any
manabie/ansible$ ansible-playbook -i hosts.yml deploy.yml
```

## Examples
```
curl -XPOST 'http://localhost:8080/task/add?userId=1' -H 'Content-Type: application/json' -d '{"taskName":"Task 1"}'
curl -XPOST 'http://localhost:8080/task/add?userId=1&taskLimit=2' -H 'Content-Type: application/json' -d '{"taskName":"Task 1"}'
```

## ToDo
```
Add more REST APIs
Update Unit Test
Apply Frontend
Apply authentication/authorization so that only specified users have permission to use
Refactor
```