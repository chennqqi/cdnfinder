package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chennqqi/cdnfinder"
)

func main() {
	var phantomjsbin string
	var server = flag.String("server", "8.8.8.8:53", "dns server for resolution")
	var full = flag.String("full", "", "URL for full finder")
	var hostname = flag.String("host", "", "hostname for single hostname finder")
	var verbose = flag.Bool("verbose", false, "Post verbose logs to stderr")
	flag.StringVar(&phantomjsbin, "phantomjsbin", "", "path to phantomjs, if blank tmp dir is used")

	flag.Parse()
	cdnfinder.Init(phantomjsbin, "")

	if *full != "" {
		out, err := cdnfinder.FullFinder(*full, *server, time.Minute, *verbose)
		if err != nil {
			log.Fatal(err)
		}
		d, err := json.MarshalIndent(out, " ", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(d))
	} else if *hostname != "" {
		c, _, err := cdnfinder.HostnametoCDN(*hostname, *server)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c)
	} else {
		log.Fatal("full or host needs to be specified")
	}
	//
}
