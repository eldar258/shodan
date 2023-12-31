package main

import (
	"fmt"
	"log"
	"os"
	"shodan/shodan"
)

func main()  {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}

	s := shodan.New()
	info, err := s.Info()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Query Credits: %d, Scan Credits: %d\n\n", info.QueryCredits, info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
