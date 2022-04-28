package receiver

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"os"
	"os/signal"
	"path"
	"syscall"
)

func HandlRec() {

	s, r := net.Listen("tcp", ":7331")

	cancel := make(chan os.Signal, 1)

	signal.Notify(cancel, os.Interrupt, syscall.SIGTERM)

	pendingFiles := make(map[string]uint8)

	if r != nil {
		fmt.Print(r.Error())
		os.Exit(1)
	}

	//deleting not fully transferred files on cancelation
	go func() {
		<-cancel
		fmt.Println("Aborting ...")
		fmt.Println("Deleting corrupted files ...")
		for f := range pendingFiles {
			os.Remove(f)
		}
		fmt.Println("Bye !")
		os.Exit(0)
	}()

	for {

		con, _ := s.Accept()
		// handling connection on a new goroutine
		go conHandler(con, pendingFiles)
	}
}

func errHandler(c net.Conn, e error, msg string) {
	c.Write([]byte("[-] Error : " + msg))
	c.Close()
	fmt.Println("[-] ", msg)
}

func conHandler(c net.Conn, m map[string]uint8) {

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

	// checking if file with same name&format already exists
	if _, exists := os.Stat(fileName); exists == nil {
		errHandler(c, exists, "File already exists .")
		return
	}

	m[fileName] = 1 // adding filename to the queue

	outputFile, errF := os.OpenFile(path.Join(".", fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if errF != nil {
		errHandler(c, errF, errF.Error())
		return
	}

	defer outputFile.Close()

	//writing {datasize} byte from connection to file {filename}

	block := math.Floor(float64(dataSize) / 100)
	blockNum := 1

	for i := 1; i <= int(dataSize); i++ {
		b, _ := stream.ReadByte()

		outputFile.Write([]byte{b})
		if i%int(block) == 0 {

			fmt.Println("Received ", fmt.Sprintf("%-3d / 100 blocks of ", blockNum), fileName, " \xE2\x9C\x93")
			blockNum++

		}

	}
	c.Write([]byte("[+] File Received ! "))
	fmt.Println("[+] Received " + fileName)
	delete(m, fileName) //removing file name from the queue
	c.Close()
}
