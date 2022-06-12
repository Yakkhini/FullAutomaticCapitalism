// pull package is used to get klines from market.
package pull

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Pull() string {
	res, err := http.Get("https://api.binance.com/api/v3/time")
	checkError(err)

	data, err := ioutil.ReadAll(res.Body)
	checkError(err)

	return string(data)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
