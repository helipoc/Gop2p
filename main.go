package main

import (
	"bufio"
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
	var clt, f string
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

	fmt.Println("[+] Client Connected ")
	fmt.Print("[+] File to send : ")
	fmt.Scan(&f)
	if len(f) > 16 {
		fmt.Println("[-] File name too large ")
		os.Exit(0)
	}
	w := bufio.NewWriter(c)
	//TODO Add filename to TCP Header
	r, _ := w.Write([]byte(f))
	w.Write([]byte("Reaaal dataaaa bla bla bla"))
	w.Flush()
	fmt.Print("Wrote : ", r, " bytes !")
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
		rd := bufio.NewReader(c)

		fileNm, _ := rd.Peek(16)

		fmt.Print(string(fileNm) + "\n")
	}
}
