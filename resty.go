package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	//fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	//fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	/* Output
	   Response Info:
	     Error      : <nil>
	     Status Code: 200
	     Status     : 200 OK
	     Proto      : HTTP/2.0
	     Time       : 457.034718ms
	     Received At: 2020-09-14 15:35:29.784681 -0700 PDT m=+0.458137045
	     Body       :
	     {
	       "args": {},
	       "headers": {
	         "Accept-Encoding": "gzip",
	         "Host": "httpbin.org",
	         "User-Agent": "go-resty/2.3.0-dev (https://github.com/go-resty/resty)",
	         "X-Amzn-Trace-Id": "Root=1-5f5ff031-000ff6292204aa6898e4de49"
	       },
	       "origin": "0.0.0.0",
	       "url": "https://httpbin.org/get"
	     }

	   Request Trace Info:
	     DNSLookup     : 4.074657ms
	     ConnTime      : 381.709936ms
	     TCPConnTime   : 77.428048ms
	     TLSHandshake  : 299.623597ms
	     ServerTime    : 75.414703ms
	     ResponseTime  : 79.337µs
	     TotalTime     : 457.034718ms
	     IsConnReused  : false
	     IsConnWasIdle : false
	     ConnIdleTime  : 0s
	     RequestAttempt: 1
	     RemoteAddr    : 3.221.81.55:443
	*/

	// Create a Resty Client
	//client := resty.New()

	resp, err = client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit":   "20",
			"sort":    "name",
			"order":   "asc",
			"random":  strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/search_result")

	// Sample of using Request.SetQueryString method
	resp, err = client.R().
		SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/show_product")

	// If necessary, you can force response content type to tell Resty to parse a JSON response into your struct
	//resp, err = client.R().
	//	SetResult(result).
	//	ForceContentType("application/json").
	//	Get("v2/alpine/manifests/latest")
}
