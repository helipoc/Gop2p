package receiver

import (
	"fmt"
	"io"
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

	defer s.Close()

	for {

		con, _ := s.Accept()
		// handling connection on a new goroutine
		go conHandler(con)
	}
}

func conHandler(c net.Conn) {

	defer c.Close()

	data, _ := io.ReadAll(c)
	if len(data) < 16 {
		//data size is less than file name area
		return
	}

	fileNameBlock := data[:16] // file name & save format
	fileSizeBlock := data[16:27]
	fmt.Print(fileSizeBlock)
	fileDataChunk := data[27:] // data
	fileName := ""

	for _, ch := range fileNameBlock {
		if ch == 0 {
			break
		}
		fileName += string(ch)
	}
	errSaving := os.WriteFile(path.Join(".", "test"), fileDataChunk, 0664)

	if errSaving != nil {
		fmt.Print("[-] Error while Saving file ... ")
		fmt.Print(errSaving)
		os.Exit(1)
	}

	_, rr := c.Write([]byte("ok"))

	print(rr)
	fmt.Println("[+] Received " + fileName)
	c.Close()
}
