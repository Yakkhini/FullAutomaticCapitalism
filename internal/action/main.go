// action and strategy.
package action

import (
	"fmt"
	"github.com/yakkhini/fultocapital/internal/none/time"
)

func Action() {
	data := gettime.Getservertime()
	fmt.Println(data)
}
