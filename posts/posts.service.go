package posts

import (
	"github.com/go-resty/resty/v2"
)

func NewServicePost() (string, error) {
	client := resty.New()
	res, err := client.R().SetHeader("Accept", "application/json").
		Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}

	return res.String(), nil
}
