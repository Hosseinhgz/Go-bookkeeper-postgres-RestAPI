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
