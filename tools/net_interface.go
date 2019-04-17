/*
Copyright 2019 tomoncle.

Licensed under the GNU General Public License, Version 3 (the "License")
*/

package tools

import (
	"errors"
	"fmt"
	"net"
)

func GetInterface() (string, error) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "unknown", err
	}

	for _, address := range addresses {
		// 检查ip地址判断是否回环地址
		if ipInfo, ok := address.(*net.IPNet); ok &&
			!ipInfo.IP.IsLoopback() &&
			ipInfo.IP.IsGlobalUnicast() {

			if ipInfo.IP.To4() != nil {
				return ipInfo.IP.String(), nil
			}
		}
	}
	return "unknown", errors.New("can not find 1 available ip address")
}
