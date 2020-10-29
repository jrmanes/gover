package main

import (
	"fmt"
	"time"
)

var version = "development"

func main() {
	fmt.Printf("production-%s", time.Now())
	fmt.Printf("Version %s\n", version)
}
