package sites

import (
	"Price_Scraper/codes/compare"
	//"Price_Scraper/codes/myhtml"
	"Price_Scraper/codes/myhttp"
	//"Price_Scraper/codes/sites"
	//"Price_Scraper/codes/sendmail"
	"Price_Scraper/codes/save"
	"io"
)

type Site struct {
	Url      string
	Product  string
	Filename string
}

func (S Site) Extractor(product string, url string, filename string, myhtml func(body io.ReadCloser, prices *string)) {
	body, _ := myhttp.Myhttp(url)
	defer body.Close()

	//extract price from html page
	var price string
	myhtml(body, &price)

	//compare extracted price vs pervious price
	compare.Compare(price, filename)
	//if changes bigger than 5% send mail
	//if ok == true {
	//sendmail.SendMail(price, url, product)
	//}

	//save new data
	save.SaveOnFile(url, price, filename)

}
