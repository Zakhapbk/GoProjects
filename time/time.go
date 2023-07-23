package main

import (
	"fmt"
	"time"
)

func main() {
	var t string
	fmt.Scan(&t)
	fmt.Println(t)
	res, _ := time.Parse(time.RFC3339, t)
	fmt.Println(res.Format(time.UnixDate))

}
