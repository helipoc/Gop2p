package sender

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
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

	fileInfo, fStateErr := os.Stat(path.Join(".", f))

	if fStateErr != nil {
		log.Fatal(fStateErr.Error())
	}

	file, fileErr := os.ReadFile(path.Join(".", f))

	if fileErr != nil {
		log.Fatal(fileErr.Error())
	}

	w := bufio.NewWriter(c)
	r, _ := w.Write([]byte(f))

	if r < 16 {
		for i := 0; i < 16-r; i++ {
			w.WriteByte(0)
		}

	}

	sizeBlock := make([]byte, binary.MaxVarintLen64)                  // buffer for the file size
	binary.LittleEndian.PutUint64(sizeBlock, uint64(fileInfo.Size())) // packing file size as int64
	w.Write(sizeBlock)                                                // writing datasize
	w.Write(file)                                                     // writing data
	w.Flush()                                                         // send to receiver

	// Getting Reply back
	reply, _ := io.ReadAll(c)
	if len(reply) == 0 {
		fmt.Println("[-] Client Aborted !")
	} else {
		fmt.Print("Client Reply : ", string(reply))

	}
	c.Close()

}
