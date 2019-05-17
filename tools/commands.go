/*
Copyright 2019 tomoncle.

Licensed under the GNU General Public License, Version 3 (the "License")
*/

package tools

import (
	"github.com/prometheus/common/log"
	"io/ioutil"
	"os/exec"
	"strings"
)

// CentOSServiceActive 获取linux系统服务状态
// active --> true ; dead  --> false
func CentOSServiceActive(serviceName string) (int64, error) {
	cmd := exec.Command("systemctl", "status", serviceName)
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorln(err)
		return 0, err
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Errorln(err)
		return 0, nil
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Errorln(err)
		return 0, nil
	}
	data := string(opBytes)
	if strings.Contains(data, "active (running)") {
		return 1, nil
	}
	return 0, nil
}
