package press_yes

import (
	"bufio"
	"fmt"
	"os"
)

func ListenYesFromTerminal(msg string) bool {
	yes := "y"
	println(msg, fmt.Sprintf(", press %+v to continue", yes))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			if text != yes {
				return false
			}
			return true
		}
	}
	return false
}
