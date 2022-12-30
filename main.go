package main

import (
	"log"

	"github.com/tawesoft/golib/v2/dialog"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("Hi")
	ok := dialog.Info("HI")
	log.Printf("ok: %v\n", ok)
}
