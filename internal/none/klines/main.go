// klines
package getklines

import "github.com/yakkhini/fultocapital/internal/pkg/pull"

func Getklines() string {
	return pull.Pull("/fapi/v1/klines")
}
