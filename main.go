package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"strings"
	"math/rand"
	"time"
	"os"
	c "github.com/fatih/color"
	//spew "github.com/davecgh/go-spew/spew"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

/*
Generate random characters.
*/
func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

/*
Display help.
*/
func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tnspectr -url=DOMAIN")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("\tnspectr -cache=false -url=DOMAIN")
	fmt.Println("\tnspectr -url=DOMAIN age cache-control")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

/*
Returns a map array of all available headers.
@param string - URL given
@return map[string]interface{}
*/
func getURLHeaders (url string) map[string]interface {} {
	response, err := http.Head(url)
	if err != nil {
		log.Fatal("Error: Unable to download URL (", url, ") with error: ", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatal("Error: HTTP Status = ", response.Status)
	}

	headers := make(map[string]interface{})
	
	for k, v := range response.Header {
		headers[strings.ToLower(k)] = string(v[0])
	}
	
	return headers
}

/**
 *Returns the header value from a given header key, if available, else returns empty string.
@param string - URL given
@param string - Header key
@return string
*/
func getURLHeader (url string) string {
	headers := getURLHeaders(url)

	// Ways to access flag.Args()

	// Option 1
	//a := flag.Args()
	//fmt.Println(a[1])

	// Option 2
	//fmt.Println(flag.Args()[1])

	// Option 3
	//fmt.Println("arguments:", flag.Args())
	//for key, val := range flag.Args() {
	//	fmt.Println(key)
	//	fmt.Println(val)
    //}

	yellow := c.New(c.FgYellow).SprintFunc()
	white := c.New(c.FgWhite).SprintFunc()

	if flag.NArg() == 0 {
		for index, element := range headers{
			
			el, ok := element.(string)
			if ok {
				fmt.Printf("%s: %s\n", yellow(index), white(el))
			}
			//fmt.Print(c.Green(index))
			//fmt.Println(index, ": ", element)
		}
	} else {
	    for _, key := range flag.Args() {
			if value, ok := headers[key]; ok {
				fmt.Printf("%s: %s\n", yellow(key), white(value))
			}
    	}
	}
	
	return ""
}

func main () {

	rand.Seed(time.Now().UnixNano())

	///scope := flag.String("headers", "all", "Display the entire headers.")
	urlRaw := flag.String("url", "", "The URL to diagnose.")
	cache := flag.Bool("cache", false, "Enable or disable.")
	help := flag.Bool("h", false, "Display help information.")
	//age := flag.Bool("age", false, "age header.")

	// Specific flags
	/*
	age := flag.Bool("age", false, "age header.")
	cacheControl := flag.Bool("cache-control", false, "cache-control header.")
	xServedBy := flag.Bool("x-served-by", false, "x-served-by header.")
	hsts := flag.Bool("strict-transport-security", false, "strict-transport-security header.")
	eTag := flag.Bool("etag", false, "etag header.")
	setCookie := flag.Bool("set-cookie", false, "set-cookie header")
	*/

	flag.Parse()

	// Display help information.
	if flag.NArg() > 0 {
		if (*help || flag.Args()[0] == "help") {
			printHelp()
			os.Exit(0)
		}
	}
 
	if !*cache {
		queryParam := randSeq(10)
		url := fmt.Sprintf("%s?query-param=%s", *urlRaw, queryParam)
		getURLHeader(url)
	} else {
		getURLHeader(*urlRaw)
	}
	
    /*
	spew.Dump(scope)
	spew.Dump(cache)
	fmt.Println("scope:", *scope)
    fmt.Println("cache:", *cache)
	fmt.Println("arguments:", flag.Args())
	fmt.Println(rand.Int())
	*/
}