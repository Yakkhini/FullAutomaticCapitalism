// action and strategy.
package action

import (
	"fmt"
	"github.com/yakkhini/fultocapital/internal/time"
)

func Action() {
	data := time.Getservertime()
	fmt.Println(data)
}
