package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/BRUHItsABunny/bunnlog"
	gokhttp_download "github.com/BRUHItsABunny/gOkHttp-download"
	gokhttp_client "github.com/BRUHItsABunny/gOkHttp/client"
	chaturbate_client "github.com/BRUHItsABunny/go-chaturbate/client"
	"github.com/BRUHItsABunny/stringvarformatter"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var FMTVars = []string{
	":SITE",
	":MODEL_NAME", // name of the model
	// DATE will be yyyy-mm-dd
	":DATE_STREAM_STARTED",
	":DATE_DOWNLOAD_STARTED",
	// TIME hh.mm.ss
	":TIME_STREAM_STARTED",
	":TIME_DOWNLOAD_STARTED",
	// TS is a timestamp with maximum accuracy
	":TS_STREAM_STARTED",
	":TS_DOWNLOAD_STARTED",
}

type App struct {
	Cfg            *Config
	BLog           *bunnlog.BunnyLog
	Client         *chaturbate_client.ChaturbateClient
	DownloadClient *http.Client
	SessionFile    *os.File
	Stats          *gokhttp_download.GlobalDownloadTracker
	NameFormatter  *stringvarformatter.Formatter
}

func NewApp() (*App, error) {
	result := &App{}

	result.ParseCfg()
	err := result.SetupLogger()
	if err != nil {
		return nil, err
	}

	err = result.SetupHTTPClient()
	if err != nil {
		return nil, err
	}

	err = result.SetupClient()
	if err != nil {
		return nil, err
	}

	result.Stats = gokhttp_download.NewGlobalDownloadTracker(time.Duration(15) * time.Second)
	return result, nil
}

func (a *App) ParseCfg() {
	if a.Cfg == nil {
		a.Cfg = &Config{}
	}

	flag.StringVar(&a.Cfg.URL, "url", "", "A model URL or the model ID")
	flag.IntVar(&a.Cfg.TimeOut, "timeout", 15, "How many seconds the downloader will idle for before exiting")
	flag.IntVar(&a.Cfg.CutOff, "cutoff", 0, "How many minutes to record for and then exit.")
	flag.StringVar(&a.Cfg.NameFMT, "namefmt", fmt.Sprintf("[%s] %s at %s-%s", FMTVars[0], FMTVars[1], FMTVars[3], FMTVars[5]), "This can be used to format the final file name using the variables: "+strings.Join(FMTVars, ", "))

	flag.BoolVar(&a.Cfg.Debug, "debug", false, "This argument is for how verbose the logger will be")
	flag.BoolVar(&a.Cfg.Daemon, "daemon", false, "This argument is for how the UI feedback will be, if set to true it will print JSON")
	flag.BoolVar(&a.Cfg.Version, "version", false, "This argument will print the current version data and exit")
	flag.Parse()

	a.NameFormatter = stringvarformatter.NewFormatter(a.Cfg.NameFMT+".ts", FMTVars...)
}

func (a *App) SetupLogger() error {
	logFile, err := os.Create("ChaturbateDownloader.log")
	if err != nil {
		return err
	}
	var bLog bunnlog.BunnyLog
	if a.Cfg.Debug {
		bLog = bunnlog.GetBunnLog(true, bunnlog.VerbosityDEBUG, log.Ldate|log.Ltime)
	} else {
		bLog = bunnlog.GetBunnLog(false, bunnlog.VerbosityWARNING, log.Ldate|log.Ltime)
	}
	bLog.SetOutputFile(logFile)
	a.BLog = &bLog
	return nil
}

func (a *App) SetupHTTPClient() error {
	var err error
	a.DownloadClient, err = gokhttp_client.NewHTTPClient()
	if err != nil {
		return fmt.Errorf("client.NewHTTPClient: %w", err)
	}
	return nil
}

func (a *App) SetupClient() (err error) {
	a.Client, err = chaturbate_client.NewChaturbateClient(a.DownloadClient)
	return err
}

func (a *App) VersionRoutine() string {
	result := strings.Builder{}
	currentPrompt := CurrentCodeBase.PromptCurrentVersion(CurrentVersion)
	result.WriteString(currentPrompt.Output)
	latestVersion, err := CurrentCodeBase.GetLatestVersion(context.Background(), nil)
	if err != nil {
		if strings.Contains(err.Error(), "repository has no tags") {
			return result.String()
		}
		panic(fmt.Errorf("CurrentCodeBase.GetLatestVersion: %w", err))
	}
	isOutdated, latestPrompt := CurrentCodeBase.PromptLatestVersion(CurrentVersion, latestVersion)

	if isOutdated {
		result.WriteString("\n")
		result.WriteString(latestPrompt.Output)
		result.WriteString(fmt.Sprintf("You can find more here:\n%s\n", latestPrompt.UpdateURL))
	}
	return result.String()
}
