package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"io/ioutil"
	"encoding/json"

	"github.com/gozl/woocommerce"
	"github.com/gozl/woocommerce/models"
)

const (
	appName = "WoooCLI"
	appVer = "v1.0.0"
	appDesc = "a utility to fetch woocommerce orders"
)

var (
	baseURL string
	wcCred string
	page int
	outfile string
	raw bool
)

func init() {
	flag.StringVar(&baseURL, "u", "", "Wordpress base URL or local file (mandatory)")
	flag.StringVar(&wcCred, "c", "", "Credential as name:passwd (mandatory if -u is URL)")
	flag.IntVar(&page, "p", 1, "Paginate, ignored when -u is not URL")
	flag.StringVar(&outfile, "o", "", "Output to file (default stdout)")
	flag.BoolVar(&raw, "r", false, "Fetch from source unprocessed")

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "%s %s () %s\n", appName, appVer, appDesc)
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintf(os.Stdout, "Usage: %s -u /local/file.json [-r] [-o /local/outfile.json]\n", os.Args[0])
		fmt.Fprintf(os.Stdout, "       %s -u http://example.org -c <cred> [-r] [-p <n>] [-o /local/outfile.json]\n", os.Args[0])
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "Parameters:")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if baseURL != "" {
		dumpCmd()
		os.Exit(0)
	}

	fmt.Fprintln(os.Stderr, "invalid flags!")
	flag.PrintDefaults()
	os.Exit(1)
}

func dumpCmd() {
	var wc woocommerce.WooCommerce
	var credParts []string
	var resultJSON []byte
	var errJSON error

	if strings.HasPrefix(baseURL, "http://") || strings.HasPrefix(baseURL, "https://") {
		credParts = strings.Split(wcCred, ":")
		if len(credParts) != 2 {
			fmt.Println("format credential as name:passwd")
			os.Exit(1)
		}
		wc = woocommerce.NewWooCommerce(baseURL, credParts[0], credParts[1])

		if raw {
			resultJSON, errJSON = wc.GetOrdersJSON(page)
		} else {
			result, errQuery := wc.GetOrders(page)
			if errQuery != nil {
				panic(errQuery)
			}

			resultJSON, errJSON = json.Marshal(result)
		}
	} else {
		infile, errRead := ioutil.ReadFile(baseURL)
		if errRead != nil {
			panic(errRead)
		}

		if raw {
			resultJSON = infile
		} else {
			var result []models.Order
			errMarshal := json.Unmarshal(infile, &result)
			if errMarshal != nil {
				panic(errMarshal)
			}

			resultJSON, errJSON = json.Marshal(result)
		}
	}

	if errJSON != nil {
		panic(errJSON)
	}

	if len(resultJSON) == 2 && string(resultJSON) == "[]" {
		fmt.Println("no record!")
		os.Exit(0)
	}

	if outfile != "" && outfile != "-" {
		errWrite := ioutil.WriteFile(outfile, resultJSON, 0644)
		if errWrite != nil {
			panic(errWrite)
		}

		fmt.Println("write: " + outfile)
		os.Exit(0)
		return
	}

	fmt.Println(string(resultJSON))
}