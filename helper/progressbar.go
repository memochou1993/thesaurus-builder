package helper

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
	"time"
)

var (
	ProgressBar *progressbar.ProgressBar
)

func InitProgressBar(max int, step string, description string) {
	theme := progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}
	options := []progressbar.Option{
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetDescription(fmt.Sprintf("[cyan][%s][reset] %s", step, description)),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSetTheme(theme),
		progressbar.OptionOnCompletion(func() {
			_, _ = fmt.Fprintln(os.Stdout)
		}),
	}
	ProgressBar = progressbar.NewOptions(max, options...)
}

func StartPermanentProgress(max int, step string, description string) {
	InitProgressBar(max, step, description)
	for {
		if err := ProgressBar.Add(1); err != nil {
			log.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func FinishPermanentProgress() {
	if err := ProgressBar.Finish(); err != nil {
		log.Println(err)
	}
}
