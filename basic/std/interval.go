package std

import (
	"fmt"
	"time"
)

func callbackInternal(i int) {
	fmt.Println("called", i)
}

func main() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for range ticker.C {
		callbackInternal(i)

		if i == 3 {
			ticker.Stop()
			break
		}

		i++
	}
}
