package mara

func (mara *Mara) Error(service string, text string) {

	// Getting formatted message
	var message = mara.getMessage(Error, service, text)

	// Printing message to console
	if mara.FlagPrint {
		mara.printToConsole(message)
	}

	// Writing message to file
	if mara.FlagWrite {
		mara.writeToFile(message)
	}
}
