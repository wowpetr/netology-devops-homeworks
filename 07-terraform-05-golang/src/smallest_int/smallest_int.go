package main

import (
	"errors"
	"fmt"
)

func Min(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty array")
	}
	var min int = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min, nil
}

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	min, err := Min(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(min)
	}
}
