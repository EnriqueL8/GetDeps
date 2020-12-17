# GetDeps


# Instructions

Install:
```
go build && cp getdeps /usr/localbin
```


Use:
```
go mod graph > graph.txt
getdeps graph.txt <module-name>
```
