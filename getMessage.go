package mara

import (
	"strings"
	"time"
)

func (mara *Mara) getMessage(level LogLevel, service string, text string) string {

	// Check default values
	mara.checkDefaultValues()

	// Composing message according to line format
	var message = strings.Replace(mara.LineFormat, "{DATETIME}", time.Now().Format(mara.TimeFormat), -1)
	message = strings.Replace(message, "{SERVICE}", service, -1)
	message = strings.Replace(message, "{LEVEL}", string(level), -1)
	message = strings.Replace(message, "{MESSAGE}", text, -1)
	return message
}
