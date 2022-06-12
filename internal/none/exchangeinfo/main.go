// ping server
package getexchangeinfo

import "github.com/yakkhini/fultocapital/internal/pkg/pull"

func Getping() string {
	return pull.Pull("/fapi/v1/exchangeinfo")
}
