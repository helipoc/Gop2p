package main

import (
	"fmt"
	"os"

	"github.com/helipoc/Gop2p/receiver"
	"github.com/helipoc/Gop2p/sender"
)

func main() {

	var t string

	x, _ := os.Stat("xx")
	fmt.Println(x.Size())
	fmt.Println("1 - Send a file")
	fmt.Println("2 - Recieve a file")
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
