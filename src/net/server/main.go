package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	serverV2()
}

func server() {
	fmt.Println("starting the server ...")

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}

		go doServerWork(conn)
	}
}

func doServerWork(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Receive data: %v \n", string(buf[:len]))
	}
}

func serverV2() {
	flag.Parse()

	if flag.NArg() != 2 {
		panic("useage: host port")
	}

	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))

	listener := initServer(hostAndPort)

	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}

func initServer(address string) net.Listener {
	serverAddr, err := net.ResolveTCPAddr("tcp", address)
	checkError(err, "Resolving address:port failed '"+address+"'")
	listener, err := net.Listen("tcp", serverAddr.String())
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())

	return listener
}

func checkError(err error, msg string) {
	if err != nil {
		panic("ERROR: " + msg + " " + err.Error())
	}
}

func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("connection from: ", connFrom)
	sayHello(conn)

	for {
		ibuf := make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT

		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}

func sayHello(conn net.Conn) {
	str := "Let's go!\n"
	obuf := []byte(str)
	wrote, err := conn.Write(obuf)
	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">\n")
	}
}
