package xnet

import (
	"fmt"
	"net"
)

//InterfaceFilter func returns true if interface match
type InterfaceFilter func(net.Interface) bool

//IsBroadcast check if interface has flag net.FlagBroadcast
func IsBroadcast(i net.Interface) bool {
	return i.Flags&net.FlagBroadcast == net.FlagBroadcast
}

//HasAddr check if interface contains Addr
func HasAddr(i net.Interface) bool {
	addrs, err := i.Addrs()
	return err == nil && len(addrs) > 0
}

//Filter returns interfaces matching InterfaceFilter
func Filter(interfaces []net.Interface, f InterfaceFilter) []net.Interface {
	finterfaces := make([]net.Interface, 0)
	for _, in := range interfaces {
		if f(in) {
			finterfaces = append(finterfaces, in)
		}
	}
	return finterfaces
}

//First returns first interface matching InterfaceFilter
func First(interfaces []net.Interface, f InterfaceFilter) (net.Interface, error) {
	for _, in := range interfaces {
		if f(in) {
			return in, nil
		}
	}
	return net.Interface{}, fmt.Errorf("interface not found")
}
