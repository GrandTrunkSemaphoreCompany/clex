package main

import (
	"fmt"
	"github.com/GrandTrunkSemaphoreCompany/clex/clacks"

)

func main() {
	s := " . "

	// Clacks as an Image
	im := clacks.Image{
		Directory: "/tmp/clacks-1",
	}

	/*
		// Clacks as a Machine
			im := clacks.Machine{
				RPIPORT: 12
			}
	*/

	im.Write([]byte(s))

	// -----------------------------

	im = clacks.Image{
		Directory: "clacks/testdata",
	}

	/*
		// Clacks as a Camera
		im := clacks.Camera{
			Usb: 1
		}
	*/

	message := make([]byte, 100)

	chars, err := im.Read(message)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Message: %d %d :%s:", chars, len(message), message)
}