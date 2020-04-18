# kimama-golang
My sandbox

## Hello world

How to run:

```bash
cd src/hello-world/
go run ./hello-world.go 
```

result:

```bash
$ cd src/hello-world/
$ go run ./hello-world.go 
Hello world
```

## Standard I/O

How to run:

```bash
cd src/try-stdio/
go run ./try-stdio.go 
```

result:

```bash
$ cd src/try-stdio/
$ go run ./try-stdio.go 
ABC
ABC
```

## Num calc

How to run:

```bash
cd src/try-num-calc/
go run ./try-num-calc.go 
```

result:

```bash
$ go run ./try-num-calc.go 
Input intA:
3
Input intB:
5
Result:
8
```

result (with error):

```bash
$ go run ./try-num-calc.go 
Input intA:
4
Input intB:
ABC
Cannot convert to int.
```