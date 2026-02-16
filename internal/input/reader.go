package input

import (
	"bufio"
	"os"
)

func ReadLine() (text string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	return
}
