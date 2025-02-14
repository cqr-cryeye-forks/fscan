package main

import (
	"fmt"
	"github.com/shadow1ng/fscan/Plugins"
	"github.com/shadow1ng/fscan/common"
	"time"
)

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	t := time.Now().Sub(start)
	fmt.Printf("[*] The scan is over, time consuming: %s", t)
}
