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
	var err error

	if len(os.Args) > 1 {
		ipRange = os.Args[1]
	} else {
		reader := bufio.NewReader(os.Stdin)
		ipRange, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		ipRange = strings.TrimSpace(ipRange)
	}

	addr, network, err := net.ParseCIDR(ipRange)
	if err != nil {
		log.Fatal(err)
	}

	var ips []string
	for ip := addr.Mask(network.Mask); network.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "CIDR Range:\t%s\n", network)
	fmt.Fprintf(w, "Netmask:\t%d.%d.%d.%d\n", network.Mask[0], network.Mask[1], network.Mask[2], network.Mask[3])
	fmt.Fprintf(w, "First IP:\t%s\n", ips[0])
	fmt.Fprintf(w, "Last IP:\t%s\n", ips[len(ips)-1])
	fmt.Fprintf(w, "Addresses:\t%d\n", len(ips))
	w.Flush()
}
