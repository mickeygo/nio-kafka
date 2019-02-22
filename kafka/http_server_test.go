package kafka

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCallHTTPServer_Https_Get(t *testing.T) {
	NewCallHTTPServer(func(resp *http.Response) {
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		t.Logf("StatusCode: %d \n", resp.StatusCode)
		t.Logf("Body: %s \n", string(data))

		if 400 <= resp.StatusCode && resp.StatusCode <= 505 {
			t.Fail()
		}
	}, nil).Get("https://www.baidu.com")
}
