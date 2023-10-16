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
	// fmt.Println(len(banners))
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				if word != "" {
					// fmt.Println(word, 0)
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
			// fmt.Println(ch)
			for _, j := range ch {
				// n := '0'
				n := (j-32)*9 + 1
				// word = ""
				// fmt.Println(string(j))
				// if int(n)+i < len(arr) && int(n)+i >= 0 {
				word = word + arr[int(n)+i]
				// fmt.Print(word)
				fmt.Fprint(writer, arr[int(n)+i])
				// } else {
				// 	if word != "" {
				// 		// fmt.Println(word, 0)
				// 		// word = word + "\n"
				// 		ReturnArr = append(ReturnArr, word)
				// 		word = ""
				// 	} else {
				// 		ReturnArr = append(ReturnArr, "")
				// 	}
				// }

			}
			ReturnArr = append(ReturnArr, word)
			word = ""
			// fmt.Println()
			fmt.Fprintln(writer, "")
		}
	}
	writer.Flush()
	return ReturnArr
}
