package testclient

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(ctx context.Context, timeout time.Duration, url string) ([]byte, error) {

	client := &http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
