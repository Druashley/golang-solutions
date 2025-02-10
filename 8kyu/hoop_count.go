package eightkyu

import "fmt"

func HoopCount(n int) string {
	if n > 9 {
		fmt.Printf("In here 1")
		return "Great, now move on to tricks"
	} else {
		fmt.Printf("In here 2")
		return "Keep at it until you get it"
	}
}
