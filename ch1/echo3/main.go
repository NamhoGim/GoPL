// Echo3 prints its command-line arguments.
package main

import (
	"os"
	"fmt"
	"strings"
)

func main () {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
