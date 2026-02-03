package app

import (
	"resty.dev/v3"
)

var UserAgent = "swords_to_poll_shares/0.1 (https://github.com/DevYukine/swords_to_poll_shares)"

// ProvideHTTPClient returns a pre-configured Resty client for DI.
func ProvideHTTPClient() *resty.Client {
	return resty.New().SetHeader("User-Agent", UserAgent)
}
