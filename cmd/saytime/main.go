package main

import (
	"fmt"
	"time"

	"dim13.org/saytime"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.Kitchen), saytime.New(now))
}
