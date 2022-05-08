package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func calculateCost(arr []int32, threshold int32) int32 {
	var minCost int32
	for i, v := range arr {

	}
}

func main() {
	arr := []int32{1, 2, 4, 5, 2, 6}
	var t int32 = 3
	log.Println(calculateCost(arr, t))
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)

	// var arr []int32

	// for i := 0; i < int(arrCount); i++ {
	// 	arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// 	checkError(err)
	// 	arrItem := int32(arrItemTemp)
	// 	arr = append(arr, arrItem)
	// }

	// thresholdTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)
	// threshold := int32(thresholdTemp)

	// result := calculateCost(arr, threshold)

	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
}

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
