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