package main
import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
			case <- boom:
				fmt.Println("BOOM!")
				return
			case <- tick:
				fmt.Println("tick.")
			default:
				fmt.Println("     .")
				time.Sleep(5e7)
		}
	}
}
