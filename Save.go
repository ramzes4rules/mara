package mara

func (mara *Mara) Save(level LogLevel, service string, text string) {

	// Getting formatted message
	var message = mara.getMessage(level, service, text)

	// Checking logging level
	if mara.LogLevel == Error && level != Error {
		return
	}
	if mara.LogLevel == Info && (level == Trace || level == Debug) {
		return
	}
	if mara.LogLevel == Debug && level == Trace {
		return
	}

	// Checking print flag
	if mara.FlagPrint {
		mara.printToConsole(message)
	}

	// Checking write flag
	if mara.FlagWrite {
		mara.writeToFile(message)
	}
}
