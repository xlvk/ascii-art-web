package main

import (
	"bufio"
	"fmt"
	"os"
)

// Print the full outcome in the triminal
func StrArr(banners, arr []string) []string {
	num := 0
	outputFileName := "template/StrFile.txt"
	word := ""
	var ReturnArr []string
	file, errs := os.Create(outputFileName)
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		os.Exit(2)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				if word != "" {
					fmt.Println(word)
					ReturnArr = append(ReturnArr, word)
					word = ""
				} else {
					ReturnArr = append(ReturnArr, "")
				}
				// fmt.Println()
				fmt.Fprintln(writer, "")
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			word = ""
			for _, j := range ch {
				n := (j-32)*9 + 1
				word = word + arr[int(n)+1]
				fmt.Println(word)
				// fmt.Print(arr[int(n)+i])
				fmt.Fprint(writer, arr[int(n)+i])

			}
			ReturnArr = append(ReturnArr, word)
			word = ""
			// fmt.Println()
			fmt.Fprintln(writer, "")
		}
	}
	writer.Flush()
	// fmt.Println(ReturnArr)
	// nn := ""
	// for i := 0; i < len(ReturnArr); i++ {
	// 	if ReturnArr[i] == "" {
	// 		nn = nn + ReturnArr[i] + "\n"
	// 	} else {
	// 		nn = nn + "\n"
	// 	}
	// 	fmt.Println(ReturnArr[i], len(ReturnArr))
	// }
	return ReturnArr
}
