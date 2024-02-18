# go-clean-arch

## Changelog

- **v1**: checkout to the [v1 branch](https://github.com/bxcodec/go-clean-arch/tree/v1) <br>
  Proposed on 2017, archived to v1 branch on 2018 <br>
  Desc: Initial proposal by me. The story can be read here: https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

- **v2**: checkout to the [v2 branch](https://github.com/bxcodec/go-clean-arch/tree/v2) <br>
  Proposed on 2018, archived to v2 branch on 2020 <br>
  Desc: Improvement from v1. The story can be read here: https://medium.com/@imantumorang/trying-clean-architecture-on-golang-2-44d615bf8fdf

- **v3**: checkout to the [v3 branch](https://github.com/bxcodec/go-clean-arch/tree/v3) <br>
  Proposed on 2019, merged to master on 2020. <br>
  Desc: Introducing Domain package, the details can be seen on this PR [#21](https://github.com/bxcodec/go-clean-arch/pull/21)

- **v4**: master branch
  Proposed on 2024, merged to master on 2024. <br>
  Desc:

  - Declare Interfaces to the consuming side,
  - Introduce `internal` package
  - Introduce `Service-focused` package.

  Details can be seen in this PR [#88](https://github.com/bxcodec/go-clean-arch/pull/88).<br>
  You may notice it diverges from the structures seen in previous versions. I encourage you to explore the branches for each version to select the structure that appeals to you the most. In my recent projects, the code structure has progressed to version 4. However, I do not strictly advocate for one version over another. You may encounter alternative examples on the internet that align more closely with your preferences. Rest assured, the foundational concept will remain consistent or at least bear resemblance. The differences are primarily in the arrangement of directories or the integration of advanced tools directly into the setup.

## Description

This is an example of implementation of Clean Architecture in Go (Golang) projects.

Rule of Clean Architecture by Uncle Bob

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has 4 Domain layer :

- Models Layer
- Repository Layer
- Usecase Layer
- Delivery Layer

#### The diagram:

![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

The original explanation about this project's structure can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047.
It may be different already, but the concept still the same in application level, also you can see the change log from v1 to current version in Master.

### How To Run This Project

> Make Sure you have run the article.sql in your mysql

Since the project is already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make tests
```

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into your workspace
$ git clone https://github.com/bxcodec/go-clean-arch.git

#move to project
$ cd go-clean-arch

# copy the example.env to .env
$ cp example.env .env

# Run the application
$ make up

# The hot reload will running

# Execute the call in another terminal
$ curl localhost:9090/articles
```

### Tools Used:

In this project, I use some tools listed below. But you can use any similar library that have the same purposes. But, well, different library will have different implementation type. Just be creative and use anything that you really need.

- All libraries listed in [`go.mod`](https://github.com/bxcodec/go-clean-arch/blob/master/go.mod)
- ["github.com/vektra/mockery".](https://github.com/vektra/mockery) To Generate Mocks for testing needs.
