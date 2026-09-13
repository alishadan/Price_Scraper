package myhttp

import (
	"io"
	"testing"
)

func Test_myhttp(t *testing.T) {
	//input: string
	//output: io.ReadCloser, error

	body, err := Myhttp(`https://bkaghaz.ir/product/%da%a9%d8%a7%d8%ba%d8%b0-%d8%af%d8%a8%d9%84-%d8%a2-a4/`)
	if err != nil {
		t.Skipf("Skipping test - network error: %v", err)
		return
	}
	defer body.Close()
	_, err = io.ReadAll(body)
	if err == nil {
		println("myhtpp passed ")
	} else {
		println("error exist in myhttp function")
	}

}
