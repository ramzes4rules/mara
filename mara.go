package mara

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LogLevel string

const (
	Info  LogLevel = "Info"
	Debug LogLevel = "Debug"
	Trace LogLevel = "Trace"
	Error LogLevel = "Error"
)

type Mara struct {
	LineFormat string //"{DATETIME} {SERVICE} {LEVEL} {MESSAGE}"
	TimeFormat string //"02-01-2006 15:04:05"
	LogLevel   LogLevel
	Path       string // "C:\ProgramData\RSL\"
	FlagPrint  bool
	FlagWrite  bool
}

func (mara Mara) Save(level LogLevel, service string, text string) {

	//
	if mara.LogLevel == Error && level != Error {
		return
	}
	if mara.LogLevel == Info && (level == Trace || level == Debug) {
		return
	}
	if mara.LogLevel == Debug && level == Trace {
		return
	}

	// Ьфлштп
	var message = strings.Replace(mara.LineFormat, "{DATETIME}", time.Now().Format(mara.TimeFormat), -1)
	message = strings.Replace(message, "{SERVICE}", service, -1)
	message = strings.Replace(message, "{LEVEL}", string(level), -1)
	message = strings.Replace(message, "{MESSAGE}", text, -1)

	//
	if mara.FlagPrint {
		mara.PrintToConsole(message)
	}

	//
	if mara.FlagWrite {
		mara.WriteToFile(message)
	}
}

func (mara Mara) WriteToFile(msg string) {

	//
	//fmt.Println(filepath.Base(mara.Path))
	//fmt.Println(filepath.Dir(mara.Path))

	//
	var paths = strings.Split(filepath.Dir(mara.Path), "\\")

	//
	var f = fmt.Sprintf("%s", paths[0])
	for i := 1; i < len(paths); i++ {
		var dir = paths[i]
		dir = strings.Replace(dir, "{MONTH}", fmt.Sprintf("%s", time.Now().Month()), -1)
		dir = strings.Replace(dir, "{YYYY}", fmt.Sprintf("%d", time.Now().Year()), -1)
		dir = strings.Replace(dir, "{MM}", fmt.Sprintf("%02d", time.Now().Month()), -1)
		dir = strings.Replace(dir, "{DD}", fmt.Sprintf("%02d", time.Now().Day()), -1)
		f = fmt.Sprintf("%s\\%s", f, dir)
		//fmt.Println(f)
		_ = os.Mkdir(f, 0666)
	}

	//
	//fmt.Printf("%s\\%s", f, filepath.Base(mara.Path))
	file, err := os.OpenFile(fmt.Sprintf("%s\\%s", f, filepath.Base(mara.Path)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	//
	_, err = file.WriteString(fmt.Sprintf("%s\n", msg))
	if err != nil {
		return
	}

	//
	err = file.Close()
	if err != nil {
		return
	}
}

func (mara Mara) PrintToConsole(message string) {
	fmt.Println(message)
}
