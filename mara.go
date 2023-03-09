package mara

type LogLevel string

const (
	Info              LogLevel = "Info"
	Debug             LogLevel = "Debug"
	Trace             LogLevel = "Trace"
	Error             LogLevel = "Error"
	defaultLineFormat          = "{DATETIME} {SERVICE} {LEVEL} {MESSAGE}"
	defaultTimeFormat          = "02-01-2006 15:04:05"
	defaultPath                = "/logs"
	defaultLoglevel   LogLevel = Info
)

type Mara struct {
	LineFormat string   //"{DATETIME} {SERVICE} {LEVEL} {MESSAGE}"
	TimeFormat string   // "02-01-2006 15:04:05"
	LogLevel   LogLevel // "Define current logging level"
	Path       string   // "C:\ProgramData\RSL\"
	FlagPrint  bool     // "Enable to print message to console"
	FlagWrite  bool     // "Enable to write message to file"
}
