package mara

func (mara *Mara) checkDefaultValues() {

	if mara.LogLevel == "" {
		mara.LogLevel = defaultLoglevel
	}

	if mara.LineFormat == "" {
		mara.LineFormat = defaultLineFormat
	}

	if mara.TimeFormat == "" {
		mara.TimeFormat = defaultTimeFormat
	}

	// Setting default value
	if mara.Path == "" {
		mara.Path = defaultPath
	}

}
