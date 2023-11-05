# 03 / demos / 01-gcflags
#:~/go-playbook-cli/03/demos/01-gcflags (main)$ go build -ldflags '-v'
```
HEADER = -H5 -T0x401000 -R0x1000
95125 symbols, 18951 reachable
        41543 package symbols, 37040 hashed symbols, 12601 non-package symbols, 3941 external symbols
97935 liveness data
```

$ tree
```
.
├── bald-mountain_co.csv
├── demo
├── go.mod
├── main.go
└── results.go

0 directories, 5 files
```


# 03 / demos / 04-pgo

#:~/go-playbook-cli/03/demos/04-pgo (main)$ go run .
```
Elapsed:  215.771419ms
```


$ go run . -cpuprofile default.pgo
```
Elapsed:  192.98309ms
```

$ go run -pgo auto .
```
Elapsed:  210.260627ms
```

$ tree
```
.
├── bald-mountain_co.csv
├── default.pgo
├── go.mod
├── main.go
└── results.go
00 directories, 5 files
```

# 04 / demos / 01-go-plugin
$ go build -buildmode plugin ./plugin.go 


#:~/go-playbook-cli/04/demos/01-go-plugin/plugin (main)$ tree
```
.
├── go.mod
├── plugin.go
└── plugin.so

0 directories, 3 files
```


#:~/go-playbook-cli/04/demos/01-go-plugin/prog (main)$ 

$ go run . -plugin ../plugin/plugin.so 
```
Executing action...
2023/11/05 12:31:56 Did the thing
```

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

