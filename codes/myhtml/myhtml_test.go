package myhtml

import (
	"Price_Scraper/myhttp"
	"testing"
)

func Test_myhtml(t *testing.T) {
	//input: body io.ReadCloser, price *string
	//output

	body, err := myhttp.Myhttp(`https://bkaghaz.ir/product/%da%a9%d8%a7%d8%ba%d8%b0-%d8%af%d8%a8%d9%84-%d8%a2-a4/`)
	if err != nil {
		t.Skipf("Skipping test - network error: %v", err)
		return
	}
	defer body.Close()
	price := "100,000"

	Myhtml(body, &price)
	println("myhtml function passed ")

}
