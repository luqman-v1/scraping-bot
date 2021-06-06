package curl

import (
	"net/http"
)

func Get(url string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//req.Header.Set("User-Agent", "XYZ/3.0")
	req.Header.Set("User-Agent", "*")
	req.Header.Set("Disallow", "/")
	res, _ := client.Do(req)
	return res
}
