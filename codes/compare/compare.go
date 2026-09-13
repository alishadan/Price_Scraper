package compare

import (
	"encoding/json"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

//this funciton compare between extrcted price and pervoius price

func Compare(extracted_price string, filename string) (ok bool, err error) {
	pervoiusPrice, _ := readPerviousPrice(filename)
	//compare between perviousPrice and extracted_price
	pervoiusPrice, _ = parsePrice(pervoiusPrice)
	extracted_price, _ = parsePrice(extracted_price)
	price1, _ := (strconv.Atoi(pervoiusPrice))
	price2, _ := strconv.Atoi(extracted_price)

	//5% price1
	price1_5 := (price1 * 5) / 100
	diffrence_p := int(math.Abs(float64(price1) - float64(price2)))

	if diffrence_p >= price1_5 {
		//sendmail.SendMail(extracted_price, url, product)
		ok = true
		return ok, nil
	}
	if diffrence_p < price1_5 {
		//print("changes is lower than 5% \n")
		ok = false
		return ok, nil
	}

	return true, nil

}
func readPerviousPrice(filename string) (string, error) {
	//check exist of file
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		print(filename, " dont exist \n")
		return "", nil
	}

	//opening file
	file, err1 := os.Open(filename)
	if err1 != nil {
		print("error in opening file \n")
		return "", err1
	}
	defer file.Close()

	//define type struct
	type data struct {
		Site  string `json:"site"`
		Price string `json:"price"`
	}
	var data1 data

	//read data from file
	byte_v, err2 := io.ReadAll(file)
	if err2 != nil {
		print("error in occured inf io.Readall funciton in compare \n")
		return "", err2
	}

	//decoding byte_v and save them in data1
	err = json.Unmarshal(byte_v, &data1)
	if err != nil {
		print("error in json.Unmarshal funciton in compare package \n")
		return "", err

	}

	return data1.Price, nil
}
func parsePrice(price string) (string, error) {

	clean_string := strings.TrimSpace(price)
	clean_string = strings.ReplaceAll(clean_string, ",", "")
	return clean_string, nil

}
