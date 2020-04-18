# go-clean-arch

## Changelog
- **v1**: checkout to the [v1 branch](https://github.com/bxcodec/go-clean-arch/tree/v1) <br>
  Proposed on 2017, archived to v1 branch on 2018 <br>
  Desc: Initial proposal by me. The story can be read here: https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

- **v2**: checkout to the [v2 branch](https://github.com/bxcodec/go-clean-arch/tree/v2) <br>
  Proposed on 2018, archived to v2 branch on 2020 <br>
  Desc: Improvement from v1. The story can be read here: https://medium.com/@imantumorang/trying-clean-architecture-on-golang-2-44d615bf8fdf

- **v3**: master branch <br>
  Proposed on 2019, merged to master on 2020. <br>
  Desc: Introducing Domain package, the details can be seen on this PR [#21](https://github.com/bxcodec/go-clean-arch/pull/21)

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

#### The diagram:

![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

The original explanation about this project's structure  can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047
It may different already, but the concept still the same in application level, also you can see the change log from v1 to current version in Master.

### How To Run This Project
> Make Sure you have run the article.sql in your mysql


Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make test
```

#### Run the Applications
Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/bxcodec/go-clean-arch.git

#move to project
$ cd go-clean-arch

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps

# Execute the call
$ curl localhost:9090/articles

# Stop
$ make stop
```


### Tools Used:
In this project, I use some tools listed below. But you can use any simmilar library that have the same purposes. But, well, different library will have different implementation type. Just be creative and use anything that you really need. 

- All libraries listed in [`go.mod`](https://github.com/bxcodec/go-clean-arch/blob/master/go.mod) 
- ["github.com/vektra/mockery".](https://github.com/vektra/mockery) To Generate Mocks for testing needs.


### Change log 
 - 2018-04-30 : [Move to new projects folder](https://github.com/bxcodec/go-clean-arch/pull/8)
 - 2018-05-09 : [Add Context](https://github.com/bxcodec/go-clean-arch/pull/9)
