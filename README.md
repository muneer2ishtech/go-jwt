# go-jwt
JWT Auth Provider using Go Lang

## Versions
- go 1.22.1
- mysql - 8.x

## Setup DB
### MySQL Docker
```
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql
```
- You can customize docker and db propeties
    - change mysql container name from 'some-mysql' to your own choice
    - put a valid and secure password for 'my-secret-pw'
    - If you want to MySQL to be accessed on any other port, e.g. 3316 then you can set like `-p 3316:3306`
    - You can set / change other MySQL properties as per your choice
### Create DB items
See [db/mysql.sql](./db/mysql.sql)
