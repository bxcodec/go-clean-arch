# Internal

This folder should be for your internal process.
As we know, Golang using kind of package-system. Every folder we can say as a package. And it can be imported to another projects as a library.
So to avoid detail implementations like handling database, handling cache exported publicly. We put all of it in the internal package.

The idea for this folder is comes from the Golang compiler itself. As described here: [Release Notes Go 1.4](https://golang.org/doc/go1.4#internalpackages).
So all of function, struct or even interface here won't be able imported to outside of this projects (external-projects). But it's still can be imported in this projects.

\*External project means other projects other than this. Let say, we have auth service written in Go. And we have this projects. The Auth service will able to import this project as a module/library and used the functionality. And upon imports all the details like very specific implementations that stored in `/internal` won't be able visible to auth service.
