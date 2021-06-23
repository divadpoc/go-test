package libcommon

import (
    "fmt"
    "os"
    "time"

    "github.com/rs/zerolog"
)

// multiple outputs using multiwriter
// https://github.com/muya/zerolog/commit/0312c7655f615f046901e27aca1fb617cbf6ac36

const (
    defaultLogFileName string = "lic.log"
)

var (
    Logger    zerolog.Logger
    LogCommon zerolog.Logger
    LogOutput zerolog.Logger
    LogApi    zerolog.Logger
)

func InitLogging(logLevel, logOutput, outputDir string) {
    // use UTC - https://github.com/rs/zerolog/issues/77
    zerolog.TimestampFunc = func() time.Time {
        return time.Now().UTC()
    }

    switch logLevel {
    case "debug":
        zerolog.SetGlobalLevel(zerolog.DebugLevel)
    case "trace":
        zerolog.SetGlobalLevel(zerolog.TraceLevel)
    default:
        zerolog.SetGlobalLevel(zerolog.InfoLevel)
    }

    // zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
    switch logOutput {
    case "stderr":
        Logger = zerolog.New(os.Stderr)
    case "stdout":
        Logger = zerolog.New(os.Stdout)
    case "console":
        Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
    /*case "journald":
        if runtime.GOOS == "linux" {
            Logger = zerolog.New(journald.NewJournalDWriter())
        } else {
            fmt.Println("journald not supported, using stdout")
            Logger = zerolog.New(os.Stdout)
        } */
    case "file":
        // zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
        logfile, err := os.OpenFile(getLogFilePath(outputDir), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err == nil {
            Logger = zerolog.New(logfile)
            fmt.Printf("The log file is allocated at %s\n", logfile.Name())
        } else {
            fmt.Println("file logger failed, using stdout")
            Logger = zerolog.New(os.Stdout)
        }
    default:
        fmt.Println("no supported logging, using stdout")
        Logger = zerolog.New(os.Stdout)
    }
    Logger = Logger.With().Timestamp().Logger()
    Logger.Info().Msg("logging initialized")
    LogCommon = Logger.With().Str("marker", "common").Logger()
    LogOutput = Logger.With().Str("marker", "output").Logger()
    LogApi = Logger.With().Str("marker", "api   ").Logger()
}

func getLogFilePath(outputDir string) string {
    if "" != outputDir {
        if err := CreateDir(outputDir); err != nil {
            fmt.Println("failed creating log directory")
            return defaultLogFileName
        }
        return outputDir + "/" + defaultLogFileName
    }
    return defaultLogFileName
}
