package gettime

import "github.com/yakkhini/fultocapital/internal/pkg/pull"

// get server time.
func Getservertime() string {
	return pull.Pull("/fapi/v1/time")
}
