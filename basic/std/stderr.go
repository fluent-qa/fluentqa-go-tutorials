package std

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "hello error\n")
}
