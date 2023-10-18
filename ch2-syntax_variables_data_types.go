package main

var c, python, java bool

func variables() {
	var i int
	println(i, c, python, java)
}

func declare() {
	var a bool = true
	var b int = 1
	var c float32 = 3.1415
	var d string = "Hello World!"

	println(a, b, c, d)
}

func main() {
	variables()
	declare()
}
