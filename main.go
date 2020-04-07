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
	"net"
	c "github.com/fatih/color"
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
	yellow := c.New(c.FgYellow).SprintFunc()
	fmt.Println("Usage:")
	fmt.Println("\tnspectr -url=DOMAIN")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("\tnspectr -cache=false -url=DOMAIN")
	fmt.Println("\tnspectr -url=DOMAIN age cache-control")
	fmt.Println()
	fmt.Println("Additional Parameters:")
	fmt.Printf("\t%s\n", yellow("age"))
	fmt.Println("\t\tReturn age only.")
	fmt.Printf("\t%s\n", yellow("cache-control"))
	fmt.Println("\t\tReturn cache-control only.")
	fmt.Printf("\t%s\n", yellow("x-cache"))
	fmt.Println("\t\tReturn x-cache only.")
	fmt.Printf("\t%s\n", yellow("set-cookie"))
	fmt.Println("\t\tReturn set-cookie only.")
	fmt.Printf("\t%s\n", yellow("strict-transport-security"))
	fmt.Println("\t\tReturn strict-transport-security only.")
	fmt.Printf("\t%s\n", yellow("etag"))
	fmt.Println("\t\tReturn etag only.")
	fmt.Printf("\t%s\n", yellow("x-served-by"))
	fmt.Println("\t\tReturn x-served-by only.")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println()
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

/*
Returns the header value from a given header key, if available, else returns empty string.
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

/*
Get DNS information.
*/
func dnsQuery (url string) {
    p := c.New(c.FgYellow, c.Bold)
	fmt.Println()
	
	// Query IP address.
	p.Printf("Querying IP records...\n")
	iprecords, ipeErr := net.LookupIP(url)
	if ipeErr != nil {
		panic(ipeErr)
	}
	for _, ip := range iprecords {
		fmt.Println(ip)
	}
   
	fmt.Println()

	// Query CNAME records.
	p.Printf("Querying CNAME records...\n")
	cname, cnameErr := net.LookupCNAME(url)
	if cnameErr != nil {
		panic(cnameErr)
	}
	fmt.Printf("%s\n", cname)

	fmt.Println()

	// Query MX records.
	p.Printf("Querying MX records...\n")
	mxs, mxErr := net.LookupMX(url)
	if mxErr != nil {
		panic(mxErr)
	}
	for _, mx := range mxs {
		fmt.Printf("%s %v\n", mx.Host, mx.Pref)
	}

	fmt.Println()

	// Query NS records.
	p.Printf("Querying NS records...\n")
	nss, nsErr := net.LookupNS(url)
	if nsErr != nil {
		panic(nsErr)
	}
	if len(nss) == 0 {
		fmt.Printf("no record")
	}
	for _, ns := range nss {
		fmt.Printf("%s\n", ns.Host)
	}

	fmt.Println()

	// Query txt records.
	p.Printf("Querying TXT records...\n")
	txts, txtErr := net.LookupTXT(url)
	if txtErr != nil {
		panic(txtErr)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		fmt.Printf("%s\n", txt)
	}

	fmt.Println()
}

func main () {
    p := c.New(c.FgYellow, c.Bold)

	rand.Seed(time.Now().UnixNano())

	urlRaw := flag.String("url", "", "The URL to diagnose.")
	cache := flag.Bool("cache", false, "Enable or disable.")
	help := flag.Bool("h", false, "Display help information.")

	flag.Parse()

	// Display help information.
	if len(*urlRaw) < 1 {
		if (*help || flag.Args()[0] == "help") {
			printHelp()
			os.Exit(0)
		}
	}
 
	p.Printf("Querying HTTP headers...\n")

	if !*cache {
		queryParam := randSeq(10)
		url := fmt.Sprintf("http://%s?query-param=%s", *urlRaw, queryParam)
		getURLHeader(url)
	} else {
		url := fmt.Sprintf("http://%s", *urlRaw)
		getURLHeader(url)
	}

	// Check DNS records.
	dnsQuery(*urlRaw)
}