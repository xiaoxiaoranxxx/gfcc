package main

import "fmt"

// import "time"

func main() {
	// fmt.Println("Hello word")
	//GOROOT：简单来说就是GO的安装目录，这个影响不大。
	//GOPATH：表示GO的工作目录，出现找不到包的情况大概率就是这里出了问题。
	// var a int=100
	a := "sss"
	fmt.Println(a)
	fmt.Printf("%T", a)
}
