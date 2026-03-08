package desktop

import (
	"log"
	"log/slog"
	"os"
)

var logFile *os.File

func OpenLogFile() {
	file, err := os.OpenFile("memos-desktop.log", os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	os.Stdout = file
	os.Stderr = file

	logFile = file
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func CloseLogFile() {
	if logFile != nil {
		slog.Info("[Memos Server] close log file")
		err := logFile.Close()
		if err != nil {
			return
		}
	}
}
