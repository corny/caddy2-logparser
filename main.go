package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	filterRemoteAddr := flag.String("remoteAddr", "", "Filter by remote address")
	filterUserAgent := flag.String("userAgent", "", "Filter by user agent")
	filterPath := flag.String("path", "", "Filter by path")
	filterStatus := flag.Int("status", 0, "Filter by status")
	groupBy := flag.String("groupBy", "", "Group by remoteAddr or userAgent")
	flag.Parse()

	userAgents := counterMap{}
	remoteAddrs := counterMap{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var entry LogEntry

		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			log.Println(err)
			continue
		}

		remoteAddr := entry.RemoteAddr()
		userAgent := entry.UserAgent()

		if *filterRemoteAddr != "" && *filterRemoteAddr != remoteAddr {
			continue
		}
		if *filterUserAgent != "" && !strings.Contains(userAgent, *filterUserAgent) {
			continue
		}
		if *filterPath != "" && !strings.Contains(entry.Request.URI, *filterPath) {
			continue
		}
		if *filterStatus > 0 && *filterStatus != entry.Status {
			continue
		}

		switch *groupBy {
		case "userAgent":
			userAgents.Inc(userAgent)
		case "remoteAddr":
			remoteAddrs.Inc(remoteAddr)
		case "":
			fmt.Printf("%v\t%v\t%v\t%v\t%v\n", time.Unix(int64(entry.Ts), 0).Format(time.RFC3339), remoteAddr, entry.Status, entry.Request.URI, userAgent)
		default:
			log.Fatalln("invalid groupBy:", *groupBy)
		}
	}

	switch *groupBy {
	case "userAgent":
		userAgents.PrintSorted()
	case "remoteAddr":
		for _, entry := range remoteAddrs.Sorted() {
			ptrNames, _ := net.LookupAddr(entry.name)
			fmt.Println(entry.count, entry.name, ptrNames)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Panicln(err)
	}
}
