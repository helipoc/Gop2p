package receiver

import (
	"fmt"
	"io"
	"net"
	"os"
)

func HandlRec() {
	s, r := net.Listen("tcp", ":7331")

	if r != nil {
		fmt.Print(r.Error())
		os.Exit(1)
	}

	for {
		c, _ := s.Accept()
		data, _ := io.ReadAll(c)
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
		c.Write([]byte("ok"))
		fmt.Println("[+] Recieved " + fileName)
	}
}
