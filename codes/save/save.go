package save

import (
	"Price_Scraper/codes/encoding"
	"os"
)

func SaveOnFile(uRl string, price string, filename string) error {
	encodedData := encoding.Encoding_data(uRl, price)
	//create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		print("error in opening file \n")
		return err
	}
	defer file.Close()

	//write data to the file
	_, err = file.Write(encodedData)
	if err != nil {
		print("error in file.WriteString() function \n")
		return err
	}

	print("data save on", filename, " successfully \n")
	return nil
}
