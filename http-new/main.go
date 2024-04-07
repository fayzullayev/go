package main

func main() {

	server := NewAPIServer(":1234")

	server.Run()

}
