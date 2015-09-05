# OnlineJudge
Judge built on Golang for hosting competetive programming contests and problems. Real time compilation and execution
of programs against testcases. Safe compilation and execution of the user program carried out in Linux containers
protecting host against malicious code. Languages supported currently are
* C++
* C
* Python (3.4 & 2.7)
* Java
* Golang (1.4)

## Requirements
* GoLang (1.4 or higher)
* mysql

## Installation
```
go get github.com/JRonak/OnlineJudge
```
Run the firstsetup.bash to setup Linux containers
In the project directory 
```
bee run
```
Make sure $GOPATH and git are added to env. Also setup dependencies

## Dependancy
* Beego Framework
```
go get github.com/astaxie/beego
```
* Bcrypt
```
golang.org/x/crypto/bcrypt
```
* Go-Sql-Driver
```
go get github.com/go-sql-driver/mysql
```
