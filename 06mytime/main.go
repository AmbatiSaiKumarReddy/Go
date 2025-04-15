package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()

	formattedDate := presentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println(formattedDate)

	createdDate := time.Date(1997, time.February, 13, 1, 2, 3, 4, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
