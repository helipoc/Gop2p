package main

import (
	"fmt"
	"os"

	"github.com/helipoc/Gop2p/receiver"
	"github.com/helipoc/Gop2p/sender"
)

func main() {

	var t string

	fmt.Println("1 - Send a file")
	fmt.Println("2 - Receive a file")

	fmt.Print("Choose option : ")
	fmt.Scan(&t)

	switch t {
	case "1":
		{
			sender.HandlSend()
		}
	case "2":
		{
			receiver.HandlRec()

		}
	default:
		{
			os.Exit(0)
		}
	}

}
