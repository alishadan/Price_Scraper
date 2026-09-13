package main

import (
	//"Price_Scraper/codes/compare"
	"Price_Scraper/codes/myhtml"
	//"Price_Scraper/codes/myhttp"
	"Price_Scraper/codes/sites"
	//"Price_Scraper/codes/sendmail"
)

func main() {

	//#1:Paper A4 Double A
	var bkaghaz sites.Site
	bkaghaz.Product = "Paper A4 Double A"
	bkaghaz.Url = `https://bkaghaz.ir/product/%da%a9%d8%a7%d8%ba%d8%b0-%d8%af%d8%a8%d9%84-%d8%a2-a4/`
	bkaghaz.Filename = "files/bkaqhaz.txt"
	bkaghaz.Extractor(bkaghaz.Product, bkaghaz.Url, bkaghaz.Filename, myhtml.WPHtml)

	var beraito sites.Site
	beraito.Product = "Paper A4 Double A"
	beraito.Url = `https://beraito.com/Product/double-a-paper-a4/`
	beraito.Filename = "files/beraito.txt"
	beraito.Extractor(beraito.Product, beraito.Url, beraito.Filename, myhtml.WPHtml)

	var timis sites.Site
	timis.Product = "Paper A4 Double A"
	timis.Url = `https://timis.ir/%DA%A9%D8%A7%D8%BA%D8%B0-%D8%AF%D8%A8%D9%84-%D8%A2-80-%DA%AF%D8%B1%D9%85-500-%D8%A8%D8%B1%DA%AF%DB%8C-a4`
	timis.Filename = "files/timis.txt"
	timis.Extractor(timis.Product, timis.Url, timis.Filename, myhtml.TimisHtml)

	//#2: Canon i-Sensys MF237w Printer
	var technolife sites.Site
	technolife.Product = "Canon i-Sensys MF237w Printer"
	technolife.Url = `https://www.technolife.com/product-15376/%D9%BE%D8%B1%DB%8C%D9%86%D8%AA%D8%B1-%DA%A9%D8%A7%D9%86%D9%86-%D9%85%D8%AF%D9%84-i-sensys-mf237w-%D9%84%DB%8C%D8%B2%D8%B1%DB%8C-%DA%86%D9%86%D8%AF%DA%A9%D8%A7%D8%B1%D9%87-`
	technolife.Filename = "files/technolife.txt"
	technolife.Extractor(technolife.Product, technolife.Url, technolife.Filename, myhtml.TechnoHtml)

	var atlasprinter sites.Site
	atlasprinter.Product = "Canon i-Sensys MF237w Printer"
	atlasprinter.Url = `https://atlasprinter.com/product/%D9%BE%D8%B1%DB%8C%D9%86%D8%AA%D8%B1-%DA%86%D9%86%D8%AF%DA%A9%D8%A7%D8%B1%D9%87-%D9%84%DB%8C%D8%B2%D8%B1%DB%8C-%DA%A9%D8%A7%D9%86%D9%86-%D9%85%D8%AF%D9%84-i-sensys-mf237w/`
	atlasprinter.Filename = "files/atlasprinter.txt"
	atlasprinter.Extractor(atlasprinter.Product, atlasprinter.Url, atlasprinter.Filename, myhtml.WPHtml2)

	var avandprinter sites.Site
	avandprinter.Product = "Canon i-Sensys MF237w Printer"
	avandprinter.Url = `https://avandprinter.com/product/canon-i-sensys-mf237w/`
	avandprinter.Filename = "files/avandprinter.txt"
	avandprinter.Extractor(avandprinter.Product, avandprinter.Url, avandprinter.Filename, myhtml.WPHtml)

	//#3: HP LaserJet Proffesional CP5225dn
	var mahan sites.Site
	mahan.Product = "HP LaserJet Proffesional CP5225dn"
	mahan.Url = `https://mahanprinter.com/5873-%D9%BE%D8%B1%DB%8C%D9%86%D8%AA%D8%B1-%D9%84%DB%8C%D8%B2%D8%B1%DB%8C-%D8%B1%D9%86%DA%AF%DB%8C-%D8%A7%DA%86-%D9%BE%DB%8C-%D8%AA%DA%A9-%DA%A9%D8%A7%D8%B1%D9%87-hp-5225-dn-%D8%A2%DA%A9%D8%A8%D9%86%D8%AF-%D8%A8%D8%A7%DA%AF%D8%A7%D8%B1%D8%A7%D9%86%D8%AA%DB%8C-`
	mahan.Filename = "files/mahan.txt"
	mahan.Extractor(mahan.Product, mahan.Url, mahan.Filename, myhtml.MahanHtml)

	var avanhp sites.Site
	avanhp.Product = "HP LaserJet Proffesional CP5225dn"
	avanhp.Url = `https://avandprinter.com/product/hp-color-laserjet-professional-cp5225dn-printe/`
	avanhp.Filename = "files/avanhp.txt"
	avanhp.Extractor(avanhp.Product, avanhp.Url, avanhp.Filename, myhtml.WPHtml)

	var atlashp sites.Site
	atlashp.Product = "HP LaserJet Proffesional CP5225dn"
	atlashp.Url = `https://atlasprinter.com/product/%D9%BE%D8%B1%DB%8C%D9%86%D8%AA%D8%B1-%D9%84%DB%8C%D8%B2%D8%B1%DB%8C-%D8%B1%D9%86%DA%AF%DB%8C-%D8%A7%DA%86-%D9%BE%DB%8C-%D9%85%D8%AF%D9%84-laserjet-proffesional-cp5225dn/`
	atlashp.Filename = "files/atlashp.txt"
	atlashp.Extractor(atlashp.Product, atlashp.Url, atlashp.Filename, myhtml.WPHtml2)

	//#4 casio calculator fx 991es plus 2nd edition
	var avandCasio sites.Site
	avandCasio.Product = "casio calculator fx 991es plus 2nd edition"
	avandCasio.Url = `https://avandprinter.com/product/casio-calculator-fx-991es-plus-2nd-edition/`
	avandCasio.Filename = "files/avandCasio.txt"
	avandCasio.Extractor(avandCasio.Product, avandCasio.Url, avandCasio.Filename, myhtml.WPHtml)

	var royzkala sites.Site
	royzkala.Product = "casio calculator fx 991es plus 2nd edition"
	royzkala.Url = `https://royzkala.com/product/%D9%85%D8%A7%D8%B4%DB%8C%D9%86-%D8%AD%D8%B3%D8%A7%D8%A8-%D9%85%D9%87%D9%86%D8%AF%D8%B3%DB%8C-%DA%A9%D8%A7%D8%B3%DB%8C%D9%88-%D9%85%D8%AF%D9%84-fx-991es-plus-2nd/`
	royzkala.Filename = "files/royzkala.txt"
	royzkala.Extractor(royzkala.Product, royzkala.Url, royzkala.Filename, myhtml.WPHtml3)

	var edariara sites.Site
	edariara.Product = "casio calculator fx 991es plus 2nd edition"
	edariara.Url = `https://edariara.com/product/casio-calculator-model-fx-991es-plus-2nd-edition/`
	edariara.Filename = "files/edariara.txt"
	edariara.Extractor(edariara.Product, edariara.Url, edariara.Filename, myhtml.WPHtml3)

}
