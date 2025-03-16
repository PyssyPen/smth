package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type logginRoundTripper struct { //
	logger io.Writer         //
	next   http.RoundTripper //
}

func (l logginRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { //
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL) //
	return l.next.RoundTrip(r)                                                            //
}

func main() {

	//jar, err := cookiejar.New(nil)
	//jar.SetCookies()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { //
			fmt.Println(req.Response.Status) //
			fmt.Println("REDIRECT")
			return nil
		},
		Transport: &logginRoundTripper{ //
			logger: os.Stdout,             //
			next:   http.DefaultTransport, //
		},
		Timeout: time.Second * 10,
		//Jar: jar,
	}

	resp, err := client.Get("http://www.udemy.com/course/golang-ninja.com") // если http, то будет редирект
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
