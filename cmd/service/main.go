package main

import (
    "flag"
    "fmt"
    "gotest/libcommon"
)

func main() {

    configPath := flag.String("configPath", "./lic-service.yaml", "path to the configuration file")
    // debug := flag.Bool("debug", false, "sets log level to debug")
    // trace := flag.Bool("trace", false, "sets log level to trace")

    flag.Parse()

    fmt.Print("configPath: ", configPath)

    libcommon.InitLogging("info", "stderr", "./logs/")

    //libcommon.Logger.Info().Msg("version=" + version + ", Build=" + Build)

    //libcommon.version = version
    //libcommon.Build = Build

    libcommon.Logger.Info().Msg("version=" + libcommon.GetVersion() + ", Build=" + libcommon.Build)

}
