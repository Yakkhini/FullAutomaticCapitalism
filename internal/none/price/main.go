// latest price
package getprice

import "github.com/yakkhini/fultocapital/internal/pkg/pull"

func Getprice() string {
	return pull.Pull("/fapi/v1/ticker/price")
}
