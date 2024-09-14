package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
}

type Message struct {
	from    string
	payload []byte
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
	}
}

var authenticatedKeys = []string{"1285fa9360f4b4ef4d123e7764597e8f5a0370943713302a3abed5d7e9d6da4e"}

func (s *Server) Start() error {
	// Start listening the port
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	// Set listener
	s.ln = ln
	// Call goroutine to make connections with clients
	go s.acceptLoop()

	<-s.quitch
	// Close the previous message channel to avoid broken pipe and inproper messages
	close(s.msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error : ", err)
			continue
		}

		fmt.Println("new connection to the server : ", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			continue
		}
		// Assign the read data over the channel
		message := buf[:n]
		s.msgch <- Message{
			from:    conn.RemoteAddr().String(),
			payload: message,
		}
		// Authenticate the recieved key
		isAuthenticated := AuthenticateKey(string(message), authenticatedKeys)
		// Set the result => 44 is not verified, 33 is verified
		response := "44"
		if isAuthenticated {
			response = "33"
		}
		// Send the server response to the client
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("ERROR : ", err)
		}
	}
}

func main() {
	// Initialize new server
	server := NewServer(":8080")
	// A goroutine to print the recieved messages
	go func() {
		for msg := range server.msgch {
			fmt.Printf("recieved message from connection (%s): %s\n", msg.from, string(msg.payload))
		}
	}()
	log.Fatal(server.Start())
}

// Hash the recieved key
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

// Authenticate the recieved key
func AuthenticateKey(recievedKey string, keys []string) bool {
	// Verify the hashed key
	isMatched := false
	for _, key := range keys {
		isMatched = VerifyHash(recievedKey, key)
		if isMatched {
			break
		}
	}
	return isMatched
}
