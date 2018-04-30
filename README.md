# go-clean-arch

## Looking for the old code ? 
If you are looking for the old code, you can checkout to the [v1 branch](https://github.com/bxcodec/go-clean-arch/tree/v1)

_Last Updated: April 30th 2018_

## Description
This is an example of implementation of Clean Architecture in Go (Golang) projects.

Rule of Clean Architecture by Uncle Bob
 * Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
 * Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
 * Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
 * Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
 * Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has  4 Domain layer :
 * Models Layer
 * Repository Layer
 * Usecase Layer  
 * Delivery Layer

The explanation about this project's structure  can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

### How To Run This Project

```bash
#move to directory
cd $GOPATH/src/github.com/bxcodec

# Clone into YOUR $GOPATH/src
git clone https://github.com/bxcodec/go-clean-arch.git

#move to project
cd go-clean-arch

# Install Dependencies
dep ensure

# Make File
make

# Run Project
go run main.go

```

Or

```bash
# GET WITH GO GET
go get github.com/bxcodec/go-clean-arch

# Go to directory

cd $GOPATH/src/github.com/bxcodec/go-clean-arch

# Install Dependencies
dep ensure

# Make File
make

# Run Project
go run main.go
```


> Make Sure you have run the article.sql in your mysql
