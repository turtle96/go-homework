package sparse_binary

import (
	"fmt"
	"strconv"
)

func GetSparseNumber(input int) int {
	binaryArr := decimalToBinary(input)

	fmt.Println("input", input, binaryArr)

	sparse, index := isSparse(binaryArr)
	for !sparse {
		binaryArr = addOneToIndex(binaryArr, index)
		fmt.Println(binaryArr, binaryToDecimal(binaryArr))

		sparse, index = isSparse(binaryArr)
	}

	return binaryToDecimal(binaryArr)
}

func decimalToBinary(n int) []int {
	binaryStr := strconv.FormatInt(int64(n), 2)

	binaryArr := []int{}
	for _, char := range binaryStr {
		if string(char) == "1" {
			binaryArr = append(binaryArr, 1)
		} else if string(char) == "0" {
			binaryArr = append(binaryArr, 0)
		}
	}

	return binaryArr
}

func binaryToDecimal(binary []int) int {
	str := ""
	for _, b := range binary {
		if b == 1 {
			str = str + "1"
		} else if b == 0 {
			str = str + "0"
		}
	}

	i, _ := strconv.ParseInt(str, 2, 64)
	return int(i)
}

func isSparse(binary []int) (bool, int) {
	for i := len(binary) - 1; i > 0; i-- {
		if binary[i] == 1 && binary[i-1] == 1 {
			return false, i
		}
	}

	return true, -1
}

func addOneToIndex(binary []int, i int) []int {
	binary[i]++
	if binary[i] == 2 {
		binary[i] = 0

		if i != 0 {
			binary = addOneToIndex(binary, i-1)
		} else {
			binary = append([]int{1}, binary...)
		}
	}

	return binary
}
