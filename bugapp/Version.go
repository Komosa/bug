package bugapp

import (
	"fmt"
	"os"
)

func Version() {
	fmt.Printf("%s version 0.3\n", os.Args[0])

}
