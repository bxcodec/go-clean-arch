# Internal Directory Guidelines

This directory is designated for internal processes. Golang operates on a package system, where each directory is effectively treated as a package. These packages can be imported into other projects as libraries. To prevent the exposure of detailed implementations, such as database handling or cache management, to the public, these are encapsulated within the internal package.

The concept of the internal folder is inspired directly by the Golang compiler. This approach is detailed in the [Release Notes Go 1.4](https://golang.org/doc/go1.4#internalpackages). Consequently, functions, structures, or interfaces within this directory are inaccessible for importation by external projects but remain available for use within this project.

An "external project" refers to any project distinct from the current one. For instance, if there is an authentication service written in Go and this project, the authentication service can incorporate this project as a module/library. However, it will not have visibility into the specific implementations housed within the /internal directory.
