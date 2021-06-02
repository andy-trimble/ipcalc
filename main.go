package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
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

	ipRange := kingpin.Arg("range", "A CIDR range.").String()
	kingpin.Parse()

	if *ipRange == "" {
		reader := bufio.NewReader(os.Stdin)
		r, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		r = strings.TrimSpace(r)
		ipRange = &r
	}

	addr, network, err := net.ParseCIDR(*ipRange)
	if err != nil {
		log.Fatal(err)
	}

	var ips []string
	for ip := addr.Mask(network.Mask); network.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	log.Printf("CIDR Range: %s", network)
	log.Printf("Netmask: %s", network.Mask.String())
	log.Printf("First IP: %s", ips[0])
	log.Printf("Last IP: %s", ips[len(ips)-1])
	log.Printf("Addresses: %d", len(ips))
}
