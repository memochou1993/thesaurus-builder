package thesaurus

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
	"time"
)

func NewProgressBar(max int, step string, description string) *progressbar.ProgressBar {
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
	return progressbar.NewOptions(max, options...)
}

func StartPermanentProgress(bar *progressbar.ProgressBar) {
	for {
		if err := bar.Add(1); err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func FinishPermanentProgress(bar *progressbar.ProgressBar) {
	if err := bar.Finish(); err != nil {
		log.Println(err)
	}
}
