package main

import (
	"bufio"
	"io"
	"strings"
)

// func main() {
// 	str := []string{}
// 	fmt.Println(countMax(str))
// }

// func countMax(upRight []string) int {
// 	return 0
// }

// func slowestKey(keyTimes [][]int32) string {
// 	var highest int32
// 	var char int32

// 	charMap := map[int]string{
// 		0:  "a",
// 		1:  "b",
// 		2:  "c",
// 		3:  "d",
// 		4:  "e",
// 		5:  "f",
// 		6:  "g",
// 		7:  "h",
// 		8:  "i",
// 		9:  "j",
// 		10: "k",
// 		11: "l",
// 		12: "m",
// 		13: "n",
// 		14: "o",
// 		15: "p",
// 		16: "q",
// 		17: "r",
// 		18: "s",
// 		19: "t",
// 		20: "u",
// 		21: "v",
// 		22: "w",
// 		23: "x",
// 		24: "y",
// 		25: "z",
// 	}

// 	for i := 0; i < len(keyTimes); i++ {
// 		if i > 0 {
// 			if keyTimes[i][1]-keyTimes[i-1][1] > highest {
// 				highest = keyTimes[i][1] - keyTimes[i-1][1]
// 				char = keyTimes[i][0]
// 			}
// 		} else {
// 			if keyTimes[i][1] > highest {
// 				highest = keyTimes[i][1]
// 				char = keyTimes[i][0]
// 			}
// 		}
// 	}

// 	return charMap[int(char)]
// }

// func main() {
// 	fmt.Println(slowestKey([][]int32{{0, 2}, {1, 5}, {0, 9}, {2, 15}}))
// }

// func ways(total int32, k int32) int32 {
// 	dp := make([]uint32, total+1)
// 	dp[0] = 1

// 	var weight uint32 = 1
// 	for weight <= uint32(k) {
// 		for i := weight; i <= uint32(total); i++ {
// 			dp[i] += dp[i-weight]
// 		}
// 		weight++
// 	}

// 	fmt.Println(dp[total])
// 	return int32(dp[total])
// }

// func main() {
// 	fmt.Println(ways(int32(842), int32(91)))

// 	//expectednya ADALAH
// }

// func minimumGroup(predators []int) int {
// 	var currentSpecies []int
// 	predator := []int{-1}

// 	var count int

// 	for {
// 		for i := 0; i < len(predator); i++ {
// 			for _, j := range predators {
// 				if predator[i] == j {
// 					currentSpecies = append(currentSpecies, i)
// 				}
// 			}
// 		}

// 		if len(currentSpecies) > 0 {
// 			predator = make([]int, len(currentSpecies))
// 			copy(predator, currentSpecies)
// 			currentSpecies = []int{}
// 			count++
// 		} else {
// 			break
// 		}
// 	}

// 	return count
// }

// func main() {
// 	data := []int{-1, 8, 6, 0, 7, 3, 8, 9, -1, 6}
// 	fmt.Print(minimumGroup(data))
// }

// func perfectSubstring(s string, k int32) int32 {
// 	var count int32

// 	for i := 0; i < len(s); i++ {
// 		countMap := make(map[rune]int32)
// 		for j := i; j < len(s); j++ {
// 			countMap[rune(s[j])]++
// 			if isPerfect(countMap, k) {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

// func isPerfect(countMap map[rune]int32, k int32) bool {
// 	for _, v := range countMap {
// 		if v != k {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	a := "1102021222"
// 	k := int32(2)
// 	fmt.Println(perfectSubstring(a, k))
// }

/*
 * Complete the 'perfectSubstring' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func perfectSubstring(s string, k int32) int32 {
	var count int32

	for i := 0; i < len(s); i++ {
		countMap := make(map[rune]int32)
		for j := i; j < len(s); j++ {
			countMap[rune(s[j])]++
			if isPerfect(countMap, k) {
				count++
			}
			// If the length of the substring exceeds k * unique characters, move the window
			if int32(j-i+1) > k*int32(len(countMap)) {
				break
			}
		}
	}

	return count
}

func isPerfect(countMap map[rune]int32, k int32) bool {
	for _, v := range countMap {
		if v != k {
			return false
		}
	}
	return true
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	s := readLine(reader)

// 	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	k := int32(kTemp)

// 	result := perfectSubstring(s, k)

// 	fmt.Fprintf(writer, "%d\n", result)

// 	writer.Flush()
// }

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
