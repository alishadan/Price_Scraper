package compare

import (
	"testing"
)

func Test_parsePrice(t *testing.T) {
	extracted_price := "11,000,000"
	excepted := "11000000"
	result, _ := parsePrice(extracted_price)
	if result != excepted {
		print("error exist in parsePrice function \n")
	}
	if result == excepted {
		print("parsePrice passed \n")
	}

}
func Test_readPerviousPrice(t *testing.T) {
	//inpouts : empty
	//outputs: string,err
	excepted := "11000000"
	filename := "news.txt"
	result, _ := readPerviousPrice(filename)
	if excepted != result {
		print("error exist in readPerviousPrice function \n")

	}
	if result == excepted {
		print("readPerviousPrice passed \n")
	}
}

func Test_Compare(t *testing.T) {
	//inpouts : string
	//outputs: bool,err
	input := "11,000,000"
	filename := "news.txt"
	excepted := false
	result, _ := Compare(input, filename)
	if excepted != result {
		print("error exist in Compare function \n")

	}

	if result == excepted {
		print("Compare passed \n")
	}
}
