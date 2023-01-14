package main

import (
	"errors"
	"fmt"
)

func ListModBy3(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("incorrect bounds for the range")
	}
	var res []int
	for i := start; i <= end; i++ {
		if i%3 == 0 {
			res = append(res, i)
		}
	}
	return res, nil
}

func main() {
	res, err := ListModBy3(1, 100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
