package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/vishvananda/netlink"
)

func main() {
	l, _ := netlink.LinkByName("lo")
	routes, _ := netlink.RouteList(l, netlink.FAMILY_V4)
	defaultGW := routes[0].Gw
	fmt.Println(defaultGW)

	for k, v := range routes {
		fmt.Println(k, v)
		fmt.Println(v.Src, v.Dst)
	}

	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()

		for _, addr := range addrs {
			addrString := addr.String()
			family := -1
			if strings.Contains(addrString, ".") {
				family = netlink.FAMILY_V4
			} else if strings.Contains(addrString, ":") {
				family = netlink.FAMILY_V6
			}
			ip, ipnet, _ := net.ParseCIDR(addrString)
			fmt.Println(iface.Name, family, ip, ipnet)
		}
	}

}
