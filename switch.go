package main

import (
	"fmt"
	// "runtime"
	"time"
)

func main() {
	// fmt.Print("Go runs on ")

	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Println("s%.", os)
	// }

	/*fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	case today + 3:
		fmt.Println("In three days")
	case today + 4:
		fmt.Println("In four days")
	default:
		fmt.Println("Too far away")
	}
	*/

	// 没有条件的 switch
	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("Good morning")
	case t.Hour() > 0 && t.Hour() < 12:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}