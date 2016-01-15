package arp

type Table interface {
	Read(ip string) (string, error)
	Insert(ip, mac string) error
}
