package client

import "net/http"

func Auth(token string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {
			r.Header.Set("Authorization", token)

			return c.Do(r)
		})
	}
}
