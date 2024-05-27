package main

func main() {
	server := StartServer()
	if err := server.Start(":8000"); err != nil {
		return
	}
}