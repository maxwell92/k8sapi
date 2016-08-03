package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	str := "2016-07-04T04:16:05Z"

	started, _ := time.Parse(time.RFC3339, str)
	fmt.Println(started)

	tn := time.Now()
	fmt.Println(tn.Format(time.RFC3339))
	fmt.Println(time.Parse(time.RFC3339, tn.Format(time.RFC3339)))

	fmt.Println(tn.Sub(started))

	elapsedDuration := tn.Sub(started)

	fmt.Printf("%d\n", int(elapsedDuration.Hours()/24))
}
