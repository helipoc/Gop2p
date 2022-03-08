package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	var t string

	fmt.Println("1 - Send a file")
	fmt.Println("2 - Recieve a file")
	fmt.Print("Choose option : ")
	fmt.Scan(&t)

	switch t {
	case "1":
		{
			handlSend()
		}
	case "2":
		{
			handlRec()
		}
	default:
		{
			os.Exit(0)
		}
	}

}

func handlSend() {
	var clt string
	var c net.Conn
	var e error
	con := false
	fmt.Print("[+] Client's IP : ")
	fmt.Scan(&clt)
	fmt.Println("[+] Waiting for Client .. ")
	for !con {
		c, e = net.Dial("tcp", fmt.Sprintf("%s:7331", clt))
		if e == nil {
			con = true
		}
	}
	fmt.Println("[+] Connected ")
	c.Close()
}

func handlRec() {
	s, r := net.Listen("tcp", ":7331")

	if r != nil {
		fmt.Print(r.Error())
		os.Exit(0)
	}

	for {
		c, _ := s.Accept()
		fmt.Print(c.RemoteAddr())

	}
}
