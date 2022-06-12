// recent trades list
package gettrades

import "github.com/yakkhini/fultocapital/internal/pkg/pull"

func Gettrades() string {
	return pull.Pull("/fapi/v1/trades")
}
