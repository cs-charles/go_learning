package date_time_package

import (
	"fmt"
	"time"
)

var now = time.Now()

var day = time.Now().Day()

var month = time.Now().Month()

var year = time.Now().Year()

func Time() {
	fmt.Printf("%d %d %d \n", day, month, year)
	fmt.Printf(now.Format(time.RFC822) + "\n")
}
