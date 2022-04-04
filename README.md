# Go-bookkeeper-postgres-RestAPI
This app is a simple book keeping application rest api that you can connect it with your Front end
This app is writen with Go.
The branch main is deployed version
## Test
this app is up and running on Heroku:
https://bookkeeper-postgres-go.herokuapp.com/

and postgres database is also up and running on Heroku which is connected with this app
you can test the app by sending request like these examples:
// 1- show all persons
https://bookkeeper-postgres-go.herokuapp.com/persons
// 2 -show all books
https://bookkeeper-postgres-go.herokuapp.com/books
// 3 -create book (this one is post request which need json in its body)
https://bookkeeper-postgres-go.herokuapp.com/create/book
example body JSON for this request : 
{
	"Title":"The Best Horror of the Year",
	"Author":"Ellen Datlow",
	"CallNumber":762235222
}
and more


## Installation and Commands
Commands to install postgres:

Linux/WSL:
sudo apt update
sudo apt install postgresql postgresql-contrib

// create new user called postgres with password
Sudo passwd postgres
sudo service postgresql start
// go inside postgres shell
sudo -u postgres psql
// change the "postgres" user password
ALTER USER postgres PASSWORD 'Hhh44974497'
// create database called book_keeper
CREATE DATABASE book_keeper
// go out of postgres shell
\q

// Dependencies and Instalations:
go mod init github.com/github.com/Hosseinhgz/{your-repository-name}

// for connection and interact with Database
go get "github.com/jinzhu/gorm"

// MYSQL inside the gorm
go get "github.com/jinzhu/gorm/dialects/postgres"

// Gorilla
go get "github.com/gorilla/mux"

## App details
Thsi rest API is listenning on port localhost:9010
You should have user and database created and ready in Postgres
