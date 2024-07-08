package main

import (
	"net"
	"fmt"
	"os"
)

const (
	CONNECTION_STRING = "user=logan password=leauxgan1 dbname=jsframeworks sslmode=disable"
	SOCKET_HOST = "localhost"
	SOCKET_PORT = "5000"
	SERVER_TYPE = "tcp"
)

func processClient(connection net.Conn) {
	defer connection.Close()
	buffer := make([]byte,1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	message := string(buffer[:mLen])
	fmt.Println("Received: ", message)
	_, err = connection.Write([]byte("Got your message!: " + message))

}
func StartupWebsocket() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SOCKET_HOST+":"+SOCKET_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SOCKET_HOST + ":" + SOCKET_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}
