package gracer

import (
	"io/ioutil"
	"net/http"
)

fung SetupRoutes() {
	http.HandleFunc("google", google)
	http.HandleFunc("bing", bing)
	http.HandleFunc("yahoo", yahoo)
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am ready to be created!"))
}

func google(w http.ResponseWriter, r *http.Request) {
	u := "https://www.google.com/search?q=" + r.URL.Query().Get("q")
	w.Write(query(u))
}

func bing(w http.ResponseWriter, r *http.Request) {
	u := "https://www.bing.com/search?q=" + r.URL.Query().Get("q")
	w.Write(query(u))
}

func yahoo(w http.ResponseWriter, r *http.Request) {
	u := "https://search.yahoo.com/search?p=" + r.URL.Query().Get("q")
	w.Write(query(u))
}

func query(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		return []byte("The race is flawed;", u, " errored:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("The race if flawed; could not read bytes from body: " + err)
	}
}
