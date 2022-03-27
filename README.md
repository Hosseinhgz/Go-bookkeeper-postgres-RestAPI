# Go-bookkeeper-postgres-RestAPI
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
ALTER USER postgres PASSWORD 'Hh4497'
// create database called book_store
CREATE DATABASE book_store
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
