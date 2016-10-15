package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

func RunHost(ip string) {
	ipAndPort := fmt.Sprintf("%v:%v", ip, port)
	listener, err := net.Listen("tcp", ipAndPort)
	if err != nil {
		log.Fatal("net listen error: ", err)
	}
	fmt.Println("listening on", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}
	fmt.Println("A connection is accepted")

	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Println("Message recevied: ", message)

	fmt.Print("reply message: ")
	stdReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := stdReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("reply err: ", replyErr)
	}
	fmt.Fprint(conn, replyMessage)
}

func RunGuest(ip string) {
	ipAndPort := fmt.Sprintf("%v:%v", ip, port)
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Dialing error: ", dialErr)
	}
	for {
		handleGuest(conn)
	}
}

func handleGuest(conn net.Conn) {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Read stdin error: ", readErr)
	}
	fmt.Fprint(conn, message)

	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("reply err: ", replyErr)
	}
	log.Println("message got back: ", replyMessage)
}
