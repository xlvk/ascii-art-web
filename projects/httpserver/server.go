package main

import (
	// "errors"
	"fmt"
	// "io"
	// "io/ioutil"
	// "log"
	"bufio"
	"html/template"
	"net/http"
	"os"

	// "time"
	"net/url"
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
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			// w.Write([]byte("This is the main content."))
			sere.ExecuteTemplate(w, "index.html", nil)
			return
			} else  if r.URL.Path != "/ascii-art" {
				http.ServeFile(w, r, "template/400Error.html")
			
		} else  if r.URL.Path != "/StrFile.txt" {
			http.ServeFile(w, r, "template/404Error.html")
		}

	} else {
		http.ServeFile(w, r, "template/405Error.html")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

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
	u, err := url.Parse("http://localhost:2003")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Listening and serving on: ")
	fmt.Printf("%+v", u)
	fmt.Println()
	http.ListenAndServe(":2003", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if r.URL.Path != "/" && r.URL.Path != "/StrFile.txt" && r.URL.Path != "/ascii-art"{

			// tmp, _ := template.ParseFiles("template/404Error.html")
			// tmp.Execute(w, nil)
			http.ServeFile(w, r, "template/404Error.html")
			w.WriteHeader(http.StatusNotFound)
		}
		http.ServeFile(w, r, "template/405Error.html")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
// 	} else  if r.URL.Path != "/ascii-art" {
// 		http.ServeFile(w, r, "template/400Error.html")
	
}
	font := r.FormValue("asciiBanner")
	text := r.FormValue("asciiText")
	if len(text) >= 255 {
		http.ServeFile(w, r, "template/400Error.html")
		w.WriteHeader(http.StatusBadRequest)
	}
	// ress := Str()
	// Read the content of the file
	// fmt.Println(text)
	text = strings.ReplaceAll(text, "\\t", "   ")
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	for i := 0; i < len(argsArr); i++ {
		word := argsArr[i]
		word = word[:len(word)-1]
		if !IsValid(word) {
			fmt.Println(argsArr[i])
			http.ServeFile(w, r, "template/400Error.html")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	arr := []string{}
	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()

	if err != nil {
		// ErrorPage(w)
		http.ServeFile(w, r, "template/500Error.html")
		w.WriteHeader(http.StatusInternalServerError)
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
	// for i:=0; i <len(argsArr); i++ {
	var f []string
	// check := false
	wee := ""
	for _, arg := range argsArr {
		for _, arg2 := range arg {
			if arg2 == 13 {
				// check = true
			} else {

				wee += string(arg2)
			}
		}
		// if check {
		// 	f = append(f, "")
		// } else {
		f = append(f, wee)
		// }
		// check = false
		wee = ""
		// fmt.Println(f)
	}
	// fmt.Println(f)
	for i := 0; i < len(f); i++ {
		if f[i] != "" {
			if !IsValid(f[i]) {
				http.ServeFile(w, r, "template/400Error.html")
				w.WriteHeader(http.StatusBadRequest)
			}
		}

	}
	ress := StrArr(f, arr)
	ress = Check(ress)
	// if ress
	nn := ""
	for i := 0; i < len(ress); i++ {
		if ress[i] != "" {
			nn = nn + ress[i] + "\n"
		} else {
			nn = nn + "\n"
		}
		// fmt.Println(ress[i])
	}

	// http.HandleFunc("/StrFile.txt", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w,r,"template/StrFile.txt")
	// })

	// fmt.Println(nn)
	d := struct {
		Result string
		// Banner string
	}{
		Result: nn,
	}
	// fmt.Println(d.Result)
	tmp, _ := template.ParseFiles("template/processor.html")
	// if err != nil {

	// }
	tmp.Execute(w, d)
	// sere.ExecuteTemplate(w, "template/processor.html", d)
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
