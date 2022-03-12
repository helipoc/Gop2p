package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
		os.Exit(1)
	}
	file, fileErr := os.ReadFile(f)

	if fileErr != nil {
		fmt.Print("[-] Can't Open file")
		os.Exit(1)
	}

	w := bufio.NewWriter(c)
	r, _ := w.Write([]byte(f))
	if r < 16 {
		for i := 0; i < 16-r; i++ {
			w.WriteByte(0)
		}

	}
	w.Write(file)
	w.Flush()
	c.Close()
}

func handlRec() {
	s, r := net.Listen("tcp", ":7331")

	if r != nil {
		fmt.Print(r.Error())
		os.Exit(1)
	}

	for {
		c, _ := s.Accept()
		data, _ := ioutil.ReadAll(c)
		if len(data) < 16 {
			continue
		}

		fileNameChunk := data[:16] // file name & save format
		fileDataChunk := data[16:] // data
		fileName := ""

		for _, c := range fileNameChunk {
			if c == 0 {
				break
			}
			fileName += string(c)
		}
		errSaving := os.WriteFile(fileName, fileDataChunk, 0664)
		if errSaving != nil {
			fmt.Print("[-] Error while Saving file ... ")
			fmt.Print(errSaving)
			os.Exit(1)
		}
		fmt.Println("[+] Recieved " + fileName)
	}
}
