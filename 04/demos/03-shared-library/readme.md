# 04 / demos / 03-shared-library
$ go build -buildmode c-shared hello.go 

sysadmin@ub22:~/Documents/code/go/go-playbook-cli/04/demos/03-shared-library (main)$ 
tree
```
.
├── go.mod
├── hello
├── hello.c
├── hello.go
└── hello.h

0 directories, 5 files
```

$ gcc hello.c ./hello 

sysadmin@ub22:~/Documents/code/go/go-playbook-cli/04/demos/03-shared-library (main)$ tree
```
.
├── a.out
├── go.mod
├── hello
├── hello.c
├── hello.go
├── hello.h
└── readme.md

0 directories, 7 files
```

sysadmin@ub22:~/Documents/code/go/go-playbook-cli/04/demos/03-shared-library (main)$ ./a.out 
```
Hello
```