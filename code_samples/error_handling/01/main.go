package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	nf, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer nf.Close()
	//io.Copy(nf, strings.NewReader("James Bond"))
	bs, err := ioutil.ReadAll(nf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
	_, err = sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return 0, nil
}
