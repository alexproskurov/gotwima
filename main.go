package main

import (
	"net/http"

	"github.com/soranoba/googp"
)

type CustomOGP struct {
	Title       string   `googp:"og:title"`
	Description string   `googp:"og:description"`
	Image       string `googp:"og:image"`
	Video       string `googp:"og:video,og:video:url"`
}

func Parse(res *http.Response, opts ...googp.ParserOpts) (CustomOGP, error) {
	var ogp CustomOGP
	err := googp.Parse(res, &ogp, opts...)

	return ogp, err
}
