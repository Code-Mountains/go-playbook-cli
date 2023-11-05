# 04 / demos / 02-static-library
$ go build -buildmode c-archive hello.go

sysadmin@ub22:~/Documents/code/go/go-playbook-cli/04/demos/02-static-library (main)$ tree
```
.
├── go.mod
├── hello.a
├── hello.c
├── hello.go
└── hello.h

0 directories, 5 files
```

$ go build -buildmode c-archive hello.go


sysadmin@ub22:~/Documents/code/go/go-playbook-cli/04/demos/02-static-library (main)$ tree
```
.
├── go.mod
├── hello.a
├── hello.c
├── hello.go
└── hello.h

0 directories, 5 files
```