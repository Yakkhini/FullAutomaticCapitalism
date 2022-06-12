// ping server
package getping

import "github.com/yakkhini/fultocapital/internal/pkg/pull"

func Getping() string {
	return pull.Pull("/fapi/v1/time")
}
