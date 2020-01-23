package sparse_number

import (
	"strconv"
)

func bruteForce(input int) int {
	binaryArr := decimalToBinary(input)

	//fmt.Println("input", input, binaryArr)

	sparse := isSparse(binaryArr)
	for !sparse {
		binaryArr = addOneToIndex(binaryArr, len(binaryArr)-1)
		//fmt.Println(binaryArr, binaryToDecimal(binaryArr))
		sparse = isSparse(binaryArr)
	}

	output := binaryToDecimal(binaryArr)
	//fmt.Println("output", output)
	return output
}

func GetSparseNumber(input int) int {
	binaryArr := decimalToBinary(input)

	//fmt.Println("input", input, binaryArr)

	if !isSparse(binaryArr) {
		newNum := roundUp(len(binaryArr))
		return binaryStrToDecimal(newNum)
	}

	return input
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

func binaryStrToDecimal(binary string) int {
	i, _ := strconv.ParseInt(binary, 2, 64)
	return int(i)
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

	return binaryStrToDecimal(str)
}

func isSparse(binary []int) bool {
	for i := len(binary) - 1; i > 0; i-- {
		if binary[i] == 1 && binary[i-1] == 1 {
			return false
		}
	}

	return true
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

func roundUp(len int) string {
	str := "1"
	for i := 0; i < len; i++ {
		str = str + "0"
	}

	return str
}
