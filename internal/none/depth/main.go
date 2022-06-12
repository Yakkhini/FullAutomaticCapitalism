// depth information
package getdepth

import "github.com/yakkhini/fultocapital/internal/pkg/pull"

func Getdepth() string {
	return pull.Pull("/fapi/v1/depth")
}
