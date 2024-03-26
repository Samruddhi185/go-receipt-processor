package main

func main() {
	server := CreateServer(":8000")
	server.Run()
}
