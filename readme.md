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


# /go-playbook-cli/03/demos/04-pgo (main)$ go run .
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