package main

import (
	"fmt"
	"github.com/tomoncle/linux_service_exporter/tools"
	"testing"
)

func TestGetInterface(t *testing.T) {
	ip, err := tools.GetInterface()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ip)
	}
}
