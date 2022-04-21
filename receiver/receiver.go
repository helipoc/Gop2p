package receiver

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"os"
	"path"
)

func HandlRec() {
	s, r := net.Listen("tcp", ":7331")

	if r != nil {
		fmt.Print(r.Error())
		os.Exit(1)
	}

	for {

		con, _ := s.Accept()
		// handling connection on a new goroutine
		go conHandler(con)
	}
}

func errHandler(c net.Conn, e error) {
	c.Write([]byte("[-] Error : " + e.Error()))
	c.Close()
	fmt.Println("[-] ", e.Error())
	os.Exit(1)
}

func conHandler(c net.Conn) {

	defer c.Close()

	stream := bufio.NewReader(c)

	fileName := ""
	fileSizeBlock := make([]byte, 0, 10)

	//reading file name block
	for i := 1; i <= 16; i++ {
		b, _ := stream.ReadByte()
		if b == 0 {
			continue
		}
		fileName += string(b)
	}

	//reading data size block
	for i := 1; i <= 10; i++ {
		b, _ := stream.ReadByte()
		fileSizeBlock = append(fileSizeBlock, b)
	}

	//parsing datasize block value
	dataSize := binary.LittleEndian.Uint64(fileSizeBlock)

	outputFile, errF := os.OpenFile(path.Join(".", fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if errF != nil {
		errHandler(c, errF)
	}

	defer outputFile.Close()

	//writing {datasize} byte from connection to file {filename}

	block := math.Ceil(float64(dataSize) / 100)
	fmt.Println("Receiving ... ", fileName)
	fmt.Print("[")

	for i := 1; i <= int(dataSize); i++ {
		b, _ := stream.ReadByte()

		outputFile.Write([]byte{b})
		if i%int(block) == 0 {
			fmt.Print("#")
		}

	}
	fmt.Println("]")
	c.Write([]byte("[+] File Received ! "))
	fmt.Println("[+] Received " + fileName)
	c.Close()
}
