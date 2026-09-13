package myhtml

import (
	"golang.org/x/net/html"
	"io"
)

//how to use
// for use of this funciton:
// we edit loop for or every IF: like in a page that calsses diffrent
// with this in here we change "sc-c1554bc0-0 eWrlhi coin-item-name"

func WPHtml(body io.ReadCloser, price *string) {
	tokenizer := html.NewTokenizer(body)
	if tokenizer.Err() != nil {
		print("error in html.NewTokenizer \n")
	}

	for {
		if tokenizer.Next() == html.ErrorToken {
			break
		}
		token1 := tokenizer.Token()
		if token1.Type == html.StartTagToken {
			if token1.Data == "bdi" {
				tokenizer.Next()
				token1 := tokenizer.Token()
				*price = token1.Data
			}
		}
	}

}
func WPHtml2(body io.ReadCloser, price *string) {
	tokenizer := html.NewTokenizer(body)
	if tokenizer.Err() != nil {
		print("error in html.NewTokenizer \n")
	}

	for {
		if tokenizer.Next() == html.ErrorToken {
			break
		}
		token1 := tokenizer.Token()
		if token1.Type == html.StartTagToken {
			if token1.Data == "bdi" {
				tokenizer.Next()
				tokenizer.Next()
				tokenizer.Next()
				tokenizer.Next()
				token1 := tokenizer.Token()
				*price = token1.Data
			}
		}
	}

}

func WPHtml3(body io.ReadCloser, price *string) {
	tokenizer := html.NewTokenizer(body)
	if tokenizer.Err() != nil {
		print("error in html.NewTokenizer \n")
	}

	for {
		if tokenizer.Next() == html.ErrorToken {
			break
		}
		token1 := tokenizer.Token()
		if token1.Type == html.StartTagToken {
			if token1.Data == "bdi" {
				tokenizer.Next()
				token1 := tokenizer.Token()
				*price = token1.Data
				return
			}
		}
	}

}

func TimisHtml(body io.ReadCloser, prices *string) {
	tokenizer := html.NewTokenizer(body)
	if tokenizer.Err() != nil {
		print("error in html.NewTokenizer \n")
	}

	for {
		if tokenizer.Next() == html.ErrorToken {
			break
		}
		token1 := tokenizer.Token()
		if token1.Type == html.StartTagToken {
			if token1.Data == "meta" {
				for _, value := range token1.Attr {
					if value.Key == "itemprop" {
						if value.Val == "price" {
							*prices = token1.Attr[1].Val
						}
					}
				}

			}
		}
	}

}

func TechnoHtml(body io.ReadCloser, prices *string) {
	tokenizer := html.NewTokenizer(body)
	if tokenizer.Err() != nil {
		print("error in html.NewTokenizer \n")

	}

	for {
		if tokenizer.Next() == html.ErrorToken {
			break
		}
		token1 := tokenizer.Token()
		if token1.Type == html.StartTagToken {
			if token1.Data == "p" {
				for _, value := range token1.Attr {
					if value.Val == "text-[19px] font-semiBold !leading-5 xl:text-[22px] text-primary-shade-1" {
						tokenizer.Next()
						token1 = tokenizer.Token()
						*prices = token1.Data
					}
				}

			}
		}
	}

}

func MahanHtml(body io.ReadCloser, prices *string) {
	tokenizer := html.NewTokenizer(body)
	if tokenizer.Err() != nil {
		print("error in html.NewTokenizer \n")
	}

	for {
		if tokenizer.Next() == html.ErrorToken {
			break
		}
		token1 := tokenizer.Token()
		if token1.Type == html.StartTagToken {
			if token1.Data == "span" {
				for _, value := range token1.Attr {
					if value.Val == "esprice-word" {
						tokenizer.Next()
						tokenizer.Next()
						tokenizer.Next()
						token1 = tokenizer.Token()
						*prices = token1.Attr[0].Val

					}
				}

			}
		}
	}

}
