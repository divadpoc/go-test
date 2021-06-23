package main

import (
	"flag"
	"fmt"
	"gotest/libcommon"
)

func main() {

	configPath := flag.String("configPath", "./lic-test.yaml", "path to the configuration file")

	fmt.Print("configPath: ", configPath)
	flag.Parse()

	libcommon.InitLogging("info", "stderr", "./logs/")

	libcommon.Logger.Info().Msg("version=" + libcommon.GetVersion() + ", Build=" + libcommon.Build)

}
