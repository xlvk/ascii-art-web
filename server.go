package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	fileLen = 855
)

var sere *template.Template

func init() {
	sere = template.Must(template.ParseGlob("template/index.html"))
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/template/css/style.css", Style)
	http.HandleFunc("/template/css/ErrorsStyle.css", StyleError)
	http.HandleFunc("/ascii-art", processor)
	http.HandleFunc("/download", DownLoad)
	u, err := url.Parse("http://localhost:2003")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Listening and serving on: ")
	fmt.Printf("%+v", u)
	fmt.Println()
	log.Fatal(http.ListenAndServe(":2003", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			sere.ExecuteTemplate(w, "index.html", nil)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "template/404Error.html")
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "template/405Error.html")
	}

}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if r.URL.Path != "/" && r.URL.Path != "/StrFile.txt" && r.URL.Path != "/ascii-art" {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "template/404Error.html")
		}
		// w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "template/405Error.html")
	}
	font := r.FormValue("asciiBanner")
	text := r.FormValue("asciiText")
	if len(text) >= 1000 {
		// w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "template/400Error.html")
		return
	}
	text = strings.ReplaceAll(text, "\\t", "   ")
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	for i := 0; i < len(argsArr); i++ {
		word := argsArr[i]
		if argsArr[i] != "" {
			// continue

			word = word[:len(word)-1]
		}
		if !IsValid(word) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "template/400Error.html")
			return
		}
	}
	arr := []string{}
	readFile, err := os.Open("fonts/" + font + ".txt")
	if err != nil {
		log.Println(err)
	}
	defer readFile.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "template/500Error.html")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "template/500Error.html")
		fmt.Println("File is corrupted.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	larg := len(argsArr)

	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	var f []string
	wee := ""
	for _, arg := range argsArr {
		for _, arg2 := range arg {
			if arg2 == 13 {
			} else {

				wee += string(arg2)
			}
		}
		f = append(f, wee)
		wee = ""
	}
	for i := 0; i < len(f); i++ {
		if f[i] != "" {
			if !IsValid(f[i]) {
				w.WriteHeader(http.StatusBadRequest)
				http.ServeFile(w, r, "template/400Error.html")
				return
			}
		}
	}
	ress := StrArr(f, arr)
	ress = Check(ress)
	nn := ""
	for i := 0; i < len(ress); i++ {
		if ress[i] != "" {
			nn = nn + ress[i] + "\n"
		} else {
			nn = nn + "\n"
		}
	}
	d := struct {
		Result string
	}{
		Result: nn,
	}
	tmp, _ := template.ParseFiles("template/processor.html")

	tmp.Execute(w, d)

}

func Style(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/css/style.css")
}

func StyleError(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/template/css/ErrorsStyle.css")
}

func ErrorPage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	template.Must(template.ParseGlob("template/500Error.html"))
	sere.ExecuteTemplate(w, "template/500Error.html", nil)
}

func DownLoad(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		file, err := os.Open("template/StrFile.txt")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to open file: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition", "attachment; filename=template/StrFile.txt")

		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to write file to response: %s", err.Error()), http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "template/405Error.html")
	}
}
