package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/headfirstgo/general"
)

func getSignatures(filename string) (signatures []string) {
	f, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	general.HandleErr(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		signatures = append(signatures, sc.Text())
	}
	general.HandleErr(sc.Err())
	return
}

func saveSignature(signature string) {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	f, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	general.HandleErr(err)
	_, err = fmt.Fprintln(f, signature)
	general.HandleErr(err)
	err = f.Close()
	general.HandleErr(err)
}

//Guestbook defines API of a guestbook
type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func serveHTML(w http.ResponseWriter, htmlFile string, data interface{}) {
	html, err := template.ParseFiles(htmlFile)
	general.HandleErr(err)
	err = html.Execute(w, data)
	general.HandleErr(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	saveSignature(r.FormValue("signature"))
	http.Redirect(w, r, "/guestbook", http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	signatures := getSignatures("signatures.txt")
	gb := Guestbook{len(signatures), signatures}
	serveHTML(w, "view.html", gb)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	serveHTML(w, "new.html", nil)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:3000", nil)
	log.Fatal(err)
}
