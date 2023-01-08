package main

import (
	"flag"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var hostname string
var port int

func main() {

    var hostnameURI = os.Getenv("HOSTNAME_URI")
    if len(hostnameURI) == 0 {
        hostnameURI = "http://localhost:26657"
    }
    var proxyPort, err = strconv.Atoi(os.Getenv("PROXY_PORT"))

    if err != nil{
        //executes if there is any error
        fmt.Println(err)
        proxyPort = 1889
      }
   if proxyPort == 0 {
          proxyPort = 1889
      }
	// flags declaration using flag package
	flag.StringVar(&hostname, "H", hostnameURI, "Specify hostname")
	flag.IntVar(&port, "p", proxyPort, "Specify port")

	flag.Parse() // after declaring flags we
	http.HandleFunc("/", serveCorsProxy)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// Serve a reverse proxy for a given url
func serveCorsProxy(res http.ResponseWriter, req *http.Request) {
	proxyRequest, err := http.NewRequest(req.Method, hostname, req.Body)
	proxyRequest.URL.Path = req.URL.Path
	proxyRequest.URL.RawQuery = req.URL.RawQuery
	if err != nil {
		fmt.Printf("create request error: %v", err)
		return
	}
	response, err := http.DefaultClient.Do(proxyRequest)
	if err != nil {
		fmt.Printf("proxy request error: %v", err)
		return
	}
	setHeaders(response, &res)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("response read error: %v", err)
		return
	}
	res.WriteHeader(response.StatusCode)
	_, _ = res.Write(body)
}

func setHeaders(src *http.Response, dest *http.ResponseWriter) {
	header := (*dest).Header()
	for name, values := range (*src).Header {
		for _, value := range values {
			header.Set(name, value)
		}
	}
	header.Set("access-control-allow-headers", "Accept,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With")
	header.Set("access-control-allow-methods", "GET, POST, OPTIONS")
	header.Set("access-control-allow-origin", "*")
	header.Set("access-control-expose-headers", "Content-Length,Content-Range")
	header.Set("access-control-max-age", "1728000")
}
