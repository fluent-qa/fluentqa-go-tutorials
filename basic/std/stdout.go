package std

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "hello world\n")
}
