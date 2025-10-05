package utils

import "github.com/schollz/progressbar/v3"

func NewProgressBar(length int, description string) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(length,
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetWidth(50),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetElapsedTime(false),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionShowTotalBytes(false),
		progressbar.OptionClearOnFinish(),
	)
	return bar
}

func NewSpinner(description string) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetWidth(30),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetElapsedTime(false),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionShowTotalBytes(false),
		progressbar.OptionClearOnFinish(),
	)
	return bar
}
