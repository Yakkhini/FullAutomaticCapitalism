// pull package is used to get klines from market.
package pull

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Pull(api string) string {
	resp, err := http.Get("https://testnet.binancefuture.com" + api)
	checkError(err)

	data, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(data)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
