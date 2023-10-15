package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"io/ioutil"
)

const serverPort = 2003


func HelloHandler1(w http.ResponseWriter, r *http.Request) {

	// Save path links in var.
	filePath := "/home/fatabbas/ascii-art-web/projects/httpserver/index.html"
	error404 := "/home/fatabbas/ascii-art-web/projects/httpserver/404Error.html"

	// If the path link not found.
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, error404)
		return
	}

	
	_, err1 := os.Stat(filePath)
	if os.IsNotExist(err1) {
		http.ServeFile(w, r, error404)
		return
	}
	http.ServeFile(w, r, filePath)



	//http.ServeFile(w, r, "/home/fatabbas/ascii-art-web/projects/httpserver/index.html")
}

// func main() {
// 	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "/home/fatabbas/ascii-art-web/projects/httpserver/style.css")
// 	})
// 	http.HandleFunc("/", HelloHandler)
// 	http.ListenAndServe(":2003", nil)
// 	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
// 	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
// 	if err != nil {
// 		fmt.Printf("client: could not create request: %s\n", err)
// 		os.Exit(1)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Printf("client: error making http request: %s\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("client: got response!\n")
// 	fmt.Printf("client: status code: %d\n", res.StatusCode)

// 	resBody, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Printf("client: could not read response body: %s\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("client: response body: %s\n", resBody)
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("server: %s /\n", r.Method)
// 		fmt.Fprintf(w, `{"message": "hello!"}`)
// 	})
// }

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/projects/httpserver/index.html")
}



// Example of parsing HTML and extracting data using goquery
// package main

// import (
// 	"fmt"
// 	"github.com/PuerkitoBio/goquery" // Import goquery package
// 	"log"
// 	"net/http"
// )

// func main() {
// 	url := "https://example.com" // Replace with your desired URL
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()

// 	// Parse the HTML
// 	doc, err := goquery.NewDocumentFromReader(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Use Find to extract specific data
// 	doc.Find("h1").Each(func(index int, element *goquery.Selection) {
// 		fmt.Println("Title:", element.Text())
// 	})
// }


// import (
// 	"asciiArtWeb"
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"net/http"
// )

// const (
// 	fileLen = 855
// )

// func main() {
// 	args := os.Args[1:]
// 	ArgsLen := len(args)
// 	font := "standard"
// 	StrFlagArr := []string{"--output=", "--color="}
// 	outputFile := ""
// 	ColorColor := ""
// 	text := ""
// 	// var Str []string
// 	Str := ""
// 	haha := os.Args[1:]
// 	// var color []string
// 	if ArgsLen < 1 {
// 		fmt.Println(len(os.Args), "is Not a valid amount of arguments.\n")
// 		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
// 		return
// 	} else if ArgsLen > 0 && ArgsLen < 3 {
// 		num := 0
// 		help := 0
// 		for i := 0; i < ArgsLen; i++ {
// 			num = num + 1
// 			if strings.Contains(args[i], StrFlagArr[0]) {
// 				outputFile = outputFileCheck(args[i])
// 				if ColorColor == "" {
// 					if i+1 < ArgsLen {
// 						haha = append(haha[:num-1], haha[num:]...)
// 						num = num - 1
// 					} else {
// 						haha = haha[:num-1]
// 					}

// 					help = help + 1
// 				} else {
// 					outputFile = ""
// 					continue
// 				}
// 			} else if strings.Contains(args[i], StrFlagArr[1]) {
// 				ColorColor = ColorColorCheck(args[i])
// 				if outputFile == "" {
// 					if i+1 < ArgsLen {
// 						haha = append(haha[:num-1], haha[num:]...)
// 						num = num - 1
// 					} else {
// 						haha = haha[:num-1]
// 					}
// 					help = help + 1
// 				} else {
// 					continue
// 				}
// 			}
// 		}
// 		if help == 0 {

// 		}
// 	} else if ArgsLen > 2 && ArgsLen < 6 {
// 		num := 0
// 		help := 0
// 		for i := 0; i < ArgsLen; i++ {
// 			num = num + 1
// 			if strings.Contains(args[i], StrFlagArr[0]) {
// 				if ColorColor == "" {
// 					outputFile = outputFileCheck(args[i])
// 					if i+1 < ArgsLen {
// 						haha = append(haha[:num-1], haha[num:]...)
// 					} else {
// 						haha = haha[:num-1]
// 					}
// 					num = num - 1
// 				} else {
// 					continue
// 				}
// 				help = (help + 1)
// 			} else if strings.Contains(args[i], StrFlagArr[1]) {
// 				ColorColor = ColorColorCheck(args[i])
// 				if outputFile == "" {
// 					haha = append(haha[:num-1], haha[num:]...)
// 					num = num - 1
// 					if i+1 < ArgsLen {
// 						Str = haha[num]
// 						haha = append(haha[:num], haha[num+1:]...)
// 						num = num - 1
// 					} else {
// 						haha = haha[:num]
// 					}
// 				} else if outputFile != "" {
// 					if i+1 < ArgsLen {
// 						haha = append(haha[:num], haha[num+1:]...)
// 					} else {
// 						haha = haha[:num]
// 					}
// 					num = num - 1
// 				} else if i+1 < ArgsLen {
// 					Str = args[i+1]
// 					if i+2 < ArgsLen {
// 						haha = append(haha[:num], haha[num+2:]...)
// 						num = num - 2
// 					} else {
// 						haha = haha[:num+1]
// 					}
// 				}
// 				help = help + 1
// 			}
// 		}
// 		if help == 0 {
// 			fmt.Println("Error: Invalid arguments.")
// 			fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
// 			return
// 		}
// 	} else if ArgsLen > 5 {
// 		fmt.Println("Error: Invalid arguments.")
// 		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
// 		return
// 	}
// 	ArgsLen = len(haha)
// 	// fmt.Println(haha, "4")
// 	if ArgsLen == 0 && Str != "" {
// 		haha = append(haha, Str)
// 		haha = append(haha, "standard")
// 		ArgsLen = len(haha)
// 		Str = ""
// 	} else if ArgsLen == 0 && Str == "" {
		
// 		haha = append(haha, os.Args[1])
// 		haha = append(haha, "standard")
// 		ArgsLen = len(haha)
// 	} else if ArgsLen == 1 {
// 		if Str != "" && Fonts(haha[ArgsLen-1]) {
// 			lol := haha[ArgsLen-1]
// 			haha[ArgsLen-1] = Str
// 			haha = append(haha, lol)
// 			Str = ""
// 			font = lol
// 		} else {
// 		haha = append(haha, "standard")
// 		}
		
// 		ArgsLen = len(haha)
// 	} else if ArgsLen == 2 {
// 		font = haha[ArgsLen-1]
// 		if Fonts(haha[ArgsLen-1]) {
// 			font = haha[ArgsLen-1]

// 		} else {
// 			fmt.Println("\""+font+"\"", "is Not a valid font.")
// 			os.Exit(0)
// 		}
// 	} else {
// 		fmt.Println("Error: Invalid arguments.", haha)
// 		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
// 		return
// 	}
	
// 	for i := 0; i < len(args); i++ {
// 		if !(asciiArtColor.IsValid(args[i])) {
// 			fmt.Println(args[i], "isn't a valid character/argument.")
// 			return
// 		}
// 	}

// 	text = haha[ArgsLen-2]
// 	if len(Str) > len(text) {
// 		fmt.Println("the \"", Str, "\" should be less or equal than \"", text, "\".")
// 		return
// 	} else if !(strings.Contains(text, Str)) {
// 		fmt.Println("the string", "'"+text+"'", "should contains", "'"+Str+"'", "to Print.")
// 		os.Exit(1)
// 	}

// 	// Read the content of the file
// 	text = strings.ReplaceAll(text, "\\t", "   ")
// 	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
// 	arr := []string{}
// 	readFile, err := os.Open("fonts/" + font + ".txt")
// 	defer readFile.Close()

// 	if err != nil {
// 		log.Fatalf("failed to open file: %s", err)
// 	}

// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)

// 	for fileScanner.Scan() {
// 		arr = append(arr, fileScanner.Text())
// 	}

// 	if len(arr) != fileLen {
// 		fmt.Println("File is corrupted.")
// 		return
// 	}
// 	larg := len(argsArr)
// 	if larg >= 2 {
// 		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
// 			argsArr = argsArr[:larg-1]
// 		}
// 	}

// 	if outputFile != "" {
// 		asciiArtColor.PrintBannersInFile(outputFile, argsArr, arr)
// 	} else if ColorColor != "" {
// 		asciiArtColor.PrintBannersWithColors(Str, ColorColor, argsArr, arr)
// 	} else {
// 		asciiArtColor.PrintBanners(argsArr, arr)
// 	}
// }

// func Fonts(argFont string) bool {
// 	// font := ""
// 	switch argFont {
// 	case "shadow":
// 		argFont = "shadow"
// 	case "thinkertoy":
// 		argFont = "thinkertoy"
// 	case "standard":
// 		argFont = "standard"
// 	default:
// 		return false
// 	}
// 	return true
// }

// // outputFile Error manegment.
// func outputFileCheck(output string) string {
// 	outputFile := strings.TrimPrefix(output, "--output=")
// 	if !strings.HasSuffix(outputFile, ".txt") {
// 		fmt.Println("Output file name should end with .txt")
// 		os.Exit(1)
// 	}
// 	return outputFile
// }

// func ColorColorCheck(color string) string {
// 	textColor := strings.TrimPrefix(color, "--color=")
// 	return textColor
// }

// func hello(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "hello\n")
// }