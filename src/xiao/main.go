package main

func main() {
	// go build -o server main.go server.go

	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
