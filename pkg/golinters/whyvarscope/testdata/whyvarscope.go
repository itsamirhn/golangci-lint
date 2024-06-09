//golangcitest:args -Ewhyvarscope
package testdata

import (
	"fmt"
	"math/rand"
)

func main() {
	if r := rand.Int(); r == 0 { // want "variable r can be removed and use assignee directly"
		fmt.Println("r is 0, Super Lucky!")
	}

	if r := rand.Int(); r != 0 {
		fmt.Printf("r is %d, Unlucky!", r)
	}
}
