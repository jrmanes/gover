# GoVer

Is a repository that contains differents ways to build your Go app and set the version on it.

---
## Const
Splitted in different files in which you will define the variable values.

Non prod build: 

[main_dev.go](https://github.com/jrmanes/gover/blob/master/const/main_dev.go)
```go
// +build !prod

// Package name -> P.E
package main

// const which you will set with that value -> P.E
const version = "development"
```

[main_prod.go](https://github.com/jrmanes/gover/blob/master/const/main_prod.go)
```go
//  +build prod

// Package name -> P.E
package main

// const which you will set with that value -> P.E
const version = "production"
```
---
## Vars
Build go with tags for NON pro
```go
go build
```

Build go with tags for pro
```go
go build -tags prod
```
---
## Ldflags
Build go with git commit id:

```go
go build -ldflags "-X main.version=$(git rev-parse HEAD)"
```
---