package main

import (
	"ChaturbateDownloader/app"
	"context"
	"fmt"
	"github.com/BRUHItsABunny/bunterm"
	gokhttp_download "github.com/BRUHItsABunny/gOkHttp-download"
	gokhttp_requests "github.com/BRUHItsABunny/gOkHttp/requests"
	chaturbate_api "github.com/BRUHItsABunny/go-chaturbate/api"
	chaturbate_constants "github.com/BRUHItsABunny/go-chaturbate/constants"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func main() {
	appData, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	versionOutput := appData.VersionRoutine()
	if appData.Cfg.Version {
		fmt.Println(versionOutput)
		os.Exit(0)
	}
	appData.BLog.Debug(versionOutput)

	notification := make(chan struct{}, 1)
	notification <- struct{}{}
	appData.Stats.TotalFiles.Store(1)

	modelId := appData.Cfg.URL
	if strings.Contains(modelId, "/") {
		urlSplit := strings.Split(modelId, "/")
		urlSplitLen := len(urlSplit)
		if urlSplitLen < 2 {
			// Invalid URL, error out
		}
		if len(urlSplit[urlSplitLen-1]) > 0 {
			modelId = urlSplit[urlSplitLen-1]
		} else {
			modelId = urlSplit[urlSplitLen-2]
		}
	}

	appData.BLog.Debug(fmt.Sprintf("MAIN THREAD: Current ModelId: %s", modelId))
	hlsManifest, err := appData.Client.GetHLSManifest(context.Background(), modelId)
	if err != nil {
		err = fmt.Errorf("appData.Client.GetHLSManifest: %w", err)
		appData.BLog.Error(fmt.Sprintf("MAIN Thread: %s", err.Error()))
		os.Exit(1)
	}

	if hlsManifest.RoomStatus != "public" {
		appData.BLog.Error(fmt.Sprintf("MAIN Thread: stream status: %s", hlsManifest.RoomStatus))
		fmt.Println(fmt.Sprintf("Model room is: %s", hlsManifest.RoomStatus))
		os.Exit(1)
	}

	nameVars := appData.NameFormatter.GetVarMap()
	nameVars[app.FMTVars[0]] = "chaturbate.com"
	nameVars[app.FMTVars[1]] = modelId
	nameVars[app.FMTVars[2]] = ""
	nameVars[app.FMTVars[3]] = time.Now().Format("2006-01-02")
	nameVars[app.FMTVars[4]] = ""
	nameVars[app.FMTVars[5]] = time.Now().Format("15-04-05")
	nameVars[app.FMTVars[6]] = ""
	nameVars[app.FMTVars[7]] = strconv.FormatInt(time.Now().UnixMilli(), 10)
	fileName := appData.NameFormatter.Format(4, nameVars) // should sanitize
	appData.BLog.Debug(fmt.Sprintf("Downloading file: %s", fileName))

	mediaHeaders := chaturbate_api.DefaultHeaders(chaturbate_constants.BaseURL)
	mediaHeaders["sec-fetch-site"] = []string{"cross-site"}
	mediaHeaders["origin"] = []string{chaturbate_constants.BaseURL[:len(chaturbate_constants.BaseURL)-1]}
	delete(mediaHeaders, "x-requested-with")
	reqOpts := []gokhttp_requests.Option{
		gokhttp_requests.NewHeaderOption(mediaHeaders),
	}

	newTask, err := gokhttp_download.NewStreamHLSTask(appData.Stats, appData.DownloadClient, hlsManifest.Url, fileName, false, reqOpts...)
	if err != nil {
		err = fmt.Errorf("gokhttp_download.NewStreamHLSTask: %w", err)
		appData.BLog.Warn(fmt.Sprintf("MAIN Thread: %s", err.Error()))
	}

	var ctxCancel context.CancelFunc
	ctx := context.Background()
	if appData.Cfg.CutOff > 0 {
		ctx, ctxCancel = context.WithDeadline(ctx, time.Now().Add(time.Minute*time.Duration(appData.Cfg.CutOff)))
		defer ctxCancel()
	}

	go func() {
		err = newTask.Download(ctx)
		if err != nil {
			err = fmt.Errorf("newTask.Download: %w", err)
			appData.BLog.Warn(fmt.Sprintf("DL Thread: %s", err.Error()))
		}

		appData.BLog.Debug("DL Thread: Waiting for threads to finish")
		appData.Stats.Stop()
	}()

	// UI
	appData.Stats.PollIP(appData.DownloadClient)
	appData.BLog.Debug("Starting the UI thread")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ticker := time.Tick(time.Second)

	term := bunterm.DefaultTerminal
	continueLoop := true
	for continueLoop {
		shouldStop := appData.Stats.GraceFulStop.Load()
		select {
		case <-c:
			appData.BLog.Debug("UI Thread: SIGTERM detected")
			shouldStop = true
			appData.Stats.Stop()
			break
		case <-ticker:
			if !appData.Cfg.Daemon {
				// Human-readable means we clear the spam
				term.ClearTerminal()
				term.MoveCursor(0, 0)
			}
			fmt.Println(appData.Stats.Tick(!appData.Cfg.Daemon))
			break
		}

		if shouldStop || appData.Stats.IdleTimeoutExceeded() {
			continueLoop = false
			if shouldStop {
				appData.BLog.Warn("UI Thread graceful stop")
			}
			appData.BLog.Warn("Downloaded all files")
			break
		}
	}
	appData.BLog.Debug("Stopping the UI thread")
}
