package logger

func CheckError(err error) {
	if err != nil {
		Logger.Fatal(err.Error())
	}
}

func CheckWarning(err error) {
	if err != nil {
		Logger.Info(err.Error())
	}
}
