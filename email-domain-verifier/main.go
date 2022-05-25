package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain, hasMX,hasSPF,SPFrecord,hasDMARC,DMARCrecord")
	for scanner.Scan() {
		domainVerifier(scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		log.Printf("%v\n", err)
	}
}

func domainVerifier(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var SPFrecord, DMARCrecord string
	mxdata, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("%v\n", err)
	}
	if len(mxdata) > 0 {
		hasMX = true
	}
	TXTrecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("%v\n", err)
	}
	for _, record := range TXTrecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFrecord = record
			break
		}
	}

	TXTrecords2, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("%v\n", err)
	}
	for _, record := range TXTrecords2 {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DMARCrecord = record
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v\n", domain, hasMX, hasSPF, SPFrecord, hasDMARC, DMARCrecord)

}
