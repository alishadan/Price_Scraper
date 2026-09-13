package save

import (
	"testing"
)

func Test_Save(t *testing.T) {
	//input []byte
	//output error

	price := "100,000"
	url := "example.com"
	filename := "new.txt"
	if err := SaveOnFile(url, price, filename); err == nil {
		println("Save_on_file passed")
	} else {
		println("error exist in Save_on_file function")
	}
}
