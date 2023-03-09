package mara

func (mara *Mara) Info(service string, text string) {

	// Getting formatted message
	var message = mara.getMessage(Info, service, text)

	// Checking logging level
	if mara.LogLevel == Error {
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
