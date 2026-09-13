package encoding

import (
	"encoding/json"
)

func Encoding_data(uRl string, price string) []byte {
	type data struct {
		Site  string `json:"site"`
		Price string `json:"price"`
	}
	data1 := data{uRl, price}
	encodedData, err := json.Marshal(data1)

	if err != nil {
		print("some errors happen in encoding_data function \n")
		return nil
	}
	return encodedData
}
