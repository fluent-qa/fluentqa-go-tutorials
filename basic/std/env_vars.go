package std

import (
	"fmt"
	"os"
)

func EnvDemo() {
	key := os.Getenv("API_KEY")

	fmt.Println(key)
}
