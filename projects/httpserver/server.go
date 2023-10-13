package main

import (
	// "errors"
	"fmt"
	// "io"
	// "io/ioutil"
	// "log"
	"bufio"
	"net/http"
	"os"
	"html/template"
	// "time"
	"strings"
)

const (
	fileLen = 855
)

// const serverPort = 2003

var sere *template.Template

func init() {
	sere = template.Must(template.ParseGlob("template/index.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {

	// Save path links in var.
	// filePath := "/home/fatabbas/ascii-art-web/projects/httpserver/index.html"
	// error404 := "/home/fatabbas/ascii-art-web/projects/httpserver/404Error.html"

	// // If the path link not found.
	// _, err := os.Stat(filePath)
	// if os.IsNotExist(err) {
	// 	http.ServeFile(w, r, error404)
	// 	return
	// }
	
	sere.ExecuteTemplate(w, "index.html", nil)
	
	// _, err1 := os.Stat(filePath)
	// if os.IsNotExist(err1) {
	// 	http.ServeFile(w, r, error404)
	// 	return
	// }
	// http.ServeFile(w, r, filePath)
	// http.ServeFile(w, r, "/projects/httpserver/index.html")



	//http.ServeFile(w, r, "/home/fatabbas/ascii-art-web/projects/httpserver/index.html")
}

func main() {
	http.HandleFunc("/", Index)
	http.Handle("/template/css/", http.StripPrefix("/template/css/", http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/ascii-art", processor)
	http.ListenAndServe(":2003", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	font := r.FormValue("asciiBanner")
	text := r.FormValue("asciiText")

	// ress := Str()
	// Read the content of the file
	text = strings.ReplaceAll(text, "\\t", "   ")
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()

	if err != nil {
		ErrorPage(w)
		// log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		fmt.Println("File is corrupted.")
		return
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}

	ress := StrArr(argsArr, arr)
	nn := ""
	for i := 0; i < len(ress); i++ {
		if ress[i] == "" {
			nn = nn + ress[i] + "\n"
		} else {
			nn = nn + "\n"
		}
		// fmt.Println(ress[i])
	}
	// fmt.Println(nn)
	d := struct{
		result string
		// Banner string
	}{
		result: nn,
	}
	sere.ExecuteTemplate(w, "template/processor.html", d)
}

func ErrorPage(w http.ResponseWriter) {
	template.Must(template.ParseGlob("template/500Error.html"))
	sere.ExecuteTemplate(w, "template/500Error.html", nil)
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

// func main() {
	// go func() {
	// 	mux := http.NewServeMux()
	// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Printf("server: %s /\n", r.Method)
	// 	})
	// 	server := http.Server{
	// 		Addr:    fmt.Sprintf(":%d", serverPort),
	// 		Handler: mux,
	// 	}
	// 	if err := server.ListenAndServe(); err != nil {
	// 		if !errors.Is(err, http.ErrServerClosed) {
	// 			fmt.Printf("error running http server: %s\n", err)
	// 		}
	// 	}
	// }()

	// time.Sleep(100 * time.Millisecond)
	// requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	// req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	// if err != nil {
	// 	fmt.Printf("client: could not create request: %s\n", err)
	// 	os.Exit(1)
	// }

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("error making http request: %s\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)

	// resBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Printf("client: could not read response body: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("client: response body: %s\n", resBody)
// }
