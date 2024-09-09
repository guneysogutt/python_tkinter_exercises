package main

/*
import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	authenticatedKeys := []string{"1285fa9360f4b4ef4d123e7764597e8f5a0370943713302a3abed5d7e9d6da4e"}

	// Get the server ip address
	serverIP, err := getLocalIP()
	if err != nil {
		fmt.Println("Error retrieving server IP:", err)
		os.Exit(1)
	}

	fmt.Printf("Server is running on IP address: %s, port 8080\n", serverIP)

	// Listen on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("TCP Server listening on port 8080")

	for {
		// Accept a client connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// Handle client in a new goroutine
		go handleConnection(conn, authenticatedKeys)
	}
}

func handleConnection(conn net.Conn, keys []string) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read the data sent by the client
	message, err := reader.ReadString('\n') // Declare err here

	// verify the hashed key
	isMatched := false
	for _, key := range keys {
		isMatched = VerifyHash(message, key)
		if isMatched {
			break
		}
	}

	if isMatched {
		fmt.Println("keys are matched!!!!!!!!!!!!!!!!")
	} else {
		fmt.Println("keys are NOT matched!!!!!!!!!!!!!!!!")
	}

	// Send a response back to the client with \r\n
	response := "33\n"
	_, err = conn.Write([]byte(response)) // Send the response
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}

	fmt.Println("Response sent to client")
}

// Get the server local IP address
func getLocalIP() (string, error) {
	// Connect to an external address to get the local ip
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}


func HashInput(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// Compare the given input with a known sha256 hash
func VerifyHash(input, knownHash string) bool {
	computedHash := HashInput(input)
	return computedHash == knownHash
}
*/
