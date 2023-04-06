package main

import (
	"fmt"

	"github.com/amauryeuzebio/goexpert/M10-Eventos/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
