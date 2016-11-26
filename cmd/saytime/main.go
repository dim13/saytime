package main

import (
	"fmt"
	"time"

	"dim13.org/saytime"
)

func main() {
	fmt.Println(saytime.New(time.Now()))
}
