package main

//telnet客户端

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	IAC  = byte(255)
	DONT = byte(254)
	DO   = byte(253)
	WONT = byte(252)
	WILL = byte(251)
)

func readData(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "readData conn.Read Fatal error: %s", err.Error())
			return
		}
		if n > 0 {
			fmt.Fprintf(os.Stdout, "%s", string(buffer[:n]))
		}
	}
}

func writeData(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		var command string
		//fmt.Scanf("%s", &command)
		data, _, _ := reader.ReadLine()
		command = string(data)
		command += "\r\n" //这里人手加上\r\n，不然win平台识别不了
		fmt.Fprintf(os.Stdout, "write:%s", command)
		buffer := []byte(command)

		if len(buffer) > 0 {
			_, err := conn.Write(buffer)
			if err != nil {
				fmt.Fprintf(os.Stderr, "writeData conn.Write Fatal error: %s", err.Error())
				return
			}
		}

	}
}

func handleNegotiation(conn net.Conn) {
	buffer := make([]byte, 3)
	for {
		n, err := conn.Read(buffer)
		fmt.Printf("handle data : %v \n", buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "handleNegotiation conn.Read Fatal error: %s", err.Error())
			return
		}
		if n != 3 {
			fmt.Fprintf(os.Stderr, "handleNegotiation conn.Read buffer size is not 3 : %s", err.Error())
			return
		}
		if buffer[0] != IAC {
			fmt.Println("handleNegotiation success")
			fmt.Printf("%s", string(buffer))
			buffer = make([]byte, 3)
			return
		} else if buffer[1] == DO {
			buffer[1] = WONT
			n, err := conn.Write(buffer)
			if n != 3 {
				fmt.Fprintf(os.Stderr, "handleNegotiation conn.Write buffer size is not 3 : %s", err.Error())
				return
			}
			fmt.Printf("return data : %v \n", buffer)
		}
	}
}

func main() {

	server := "192.168.0.109:2333"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ResolveTCPAddr Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "DialTCP Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("connect success")

	handleNegotiation(conn)

	go readData(conn)
	writeData(conn)

}
