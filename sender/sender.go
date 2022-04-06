package sender

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func HandlSend() {
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
	fmt.Println("file sent")

}
