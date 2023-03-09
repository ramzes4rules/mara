package mara

func (mara *Mara) Trace(service string, text string) {

	// Getting formatted message
	var message = mara.getMessage(Trace, service, text)

	// Checking logging level
	if mara.LogLevel == Error || mara.LogLevel == Info || mara.LogLevel == Debug {
		return
	}

	// Printing message to console
	if mara.FlagPrint {
		mara.printToConsole(message)
	}

	// Writing message to file
	if mara.FlagWrite {
		mara.writeToFile(message)
	}
}
