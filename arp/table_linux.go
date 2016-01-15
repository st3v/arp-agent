// +build linux
package arp

type table struct{}

func NewTable() Table {
	return new(table)
}

func (t table) Read(ip string) (string, error) {
	return "", nil
}

func (t table) Insert(ip, mac string) error {
	return nil
}
