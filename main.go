package main

import (
	"context"
	"log/slog"
	"os"

	cmd "github.com/usememos/memos/cmd/memos"
	"github.com/usememos/memos/desktop"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func cleanup() {
	cmd.ShutdownMemosServer()
	<-cmd.ServerShutdown
	desktop.CloseLogFile()
}

func main() {
	desktop.OpenLogFile()
	defer desktop.CloseLogFile()

	cmd.StartMemosServer()

	// Create application with options
	err := wails.Run(&options.App{
		Title:    "memos",
		Width:    1024,
		Height:   768,
		MinWidth: 375,
		AssetServer: &assetserver.Options{
			Middleware: desktop.WailsServerMiddleware(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			// wait for memos server ready
			<-cmd.ServerReady
		},
		OnShutdown: func(ctx context.Context) {
			slog.Info("[wails] App shutting down")
			cleanup()
			os.Exit(1)
		},
	})

	if err != nil {
		slog.Error("[wails] App cause error:", err.Error())
	}
}
