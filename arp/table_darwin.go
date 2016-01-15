// +build darwin
package arp

import (
	"fmt"
	"os/exec"
	"strings"
)

type table struct{}

func NewTable() Table {
	return new(table)
}

func (t table) Read(ip string) (string, error) {
	cmd := exec.Command("/usr/sbin/arp", "-a", "-n")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	rows := strings.Split(string(out), "\n")
	for _, s := range rows {
		cols := strings.Split(s, " ")
		if len(cols) != 7 {
			continue
		}

		ipFound, mac := strings.Trim(cols[1], "()"), strings.Trim(cols[3], "()")
		if ipFound == ip {
			if mac == "incomplete" {
				mac = ""
			}

			return mac, nil
		}
	}

	return "", nil
}

func (t table) Insert(ip, mac string) error {
	cmd := exec.Command("/usr/sbin/arp", "-s", ip, mac, "temp")
	log, err := cmd.CombinedOutput()
	fmt.Println(string(log))
	return err
}
