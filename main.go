package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	var ipRange string

	if len(os.Args) > 1 {
		// The range was passed in as an argument
		ipRange = os.Args[1]
	} else {
		// Read from standard input
		reader := bufio.NewReader(os.Stdin)
		ipRange, _ = reader.ReadString('\n')
	}

	ipRange = strings.TrimSpace(ipRange)
	addr, network, err := net.ParseCIDR(ipRange)
	if err != nil {
		log.Fatal(err)
	}

	ipv4 := addr.To4()
	if ipv4 == nil {
		log.Fatal("built for IPv4 only")
	}

	subnetMask := network.Mask
	firstIp := make([]byte, len(ipv4))
	lastIp := make([]byte, len(ipv4))
	numIPs := 0

	for i := range ipv4 {
		firstIp[i] = ipv4[i] & subnetMask[i]
		lastIp[i] = ipv4[i] | (subnetMask[i] ^ 0xff)
		numIPs <<= 8
		numIPs += int(subnetMask[i] ^ 0xff)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "CIDR Range:\t%s\n", network)
	fmt.Fprintf(w, "Netmask:\t%d.%d.%d.%d\n", network.Mask[0], network.Mask[1], network.Mask[2], network.Mask[3])
	fmt.Fprintf(w, "First IP:\t%s\n", net.IP(firstIp))
	fmt.Fprintf(w, "Last IP:\t%s\n", net.IP(lastIp))
	fmt.Fprintf(w, "Addresses:\t%d\n", numIPs+1)
	w.Flush()
}
