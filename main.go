package main

import (
	"fmt"
	"time"
)

func main() {
	var aTime string
	fmt.Scan(&aTime)

	t, err := time.Parse(time.RFC3339, aTime)
	if err != nil {
		fmt.Println("Error for rarsing", err)
		return
	}
	fmt.Println(t.Format(time.UnixDate))
}
