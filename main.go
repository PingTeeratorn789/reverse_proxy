package main

import (
	"github.com/PingTeeratorn789/reverse_proxy/cmd"
	"fmt"
	"runtime"
	"time"
)

var (
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

func main() {
	fmt.Printf("\033[1;36mGO_VERSION = %s, BUILD_COMMIT = %s, BUILD_TIME = %s \033[0m\n",
		runtime.Version(), buildcommit, buildtime)
	cmd.Execute()
}