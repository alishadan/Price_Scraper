package sites

import (
	"Price_Scraper/codes/myhtml"
	"testing"
)

func Test_Extractor(t *testing.T) {
	//input: product string, url string, filename string,function myhtml
	//output: price string

	var exe Site
	exe.Product = "double a"
	exe.Url = "https://beraito.com/product/double-a-paper-a4/"
	exe.Filename = "example.txt"

	exe.Extractor(exe.Product, exe.Url, exe.Filename, myhtml.WPHtml)

	println("sites package passed")

}
