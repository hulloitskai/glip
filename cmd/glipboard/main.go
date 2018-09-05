package main

import (
	"github.com/steven-xie/glip"
	"log"
	"os"
)

func init() {
	configureLog()
}

func main() {
	var err error

	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalln("Could not read os.Stdin info:", err)
	}

	b, err := glip.NewBoard()
	if err != nil {
		log.Fatalln("Failed to create Board:", err)
	}

	if (info.Mode() & os.ModeCharDevice) == 0 {
		Write(b)
	} else {
		Read(b)
	}
}

// Write writes data from os.Stdin to the system clipboard.
func Write(b glip.Board) {
	if _, err := b.ReadFrom(os.Stdin); err != nil {
		log.Fatalln("Failed to write from os.Stdin to system clipboard:", err)
	}
}

// Read transfers data from the system clipboard into os.Stdout.
func Read(b glip.Board) {
	if _, err := b.WriteTo(os.Stdout); err != nil {
		log.Fatalln("Failed to write clipboard contents into os.Stdout:", err)
	}
}

func configureLog() {
	log.SetFlags(0)
}
