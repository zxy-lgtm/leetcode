package main

import (
	"fmt"
	"time"
)

func main() {
	ExpireTime := 86400 * 30
	fmt.Println(time.Second * time.Duration(ExpireTime))
}
