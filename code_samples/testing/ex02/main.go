package main

import (
	"fmt"

	"github.com/utkarshpunjab/forexample/code_samples/testing/ex02/quote"
	"github.com/utkarshpunjab/forexample/code_samples/testing/ex02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
