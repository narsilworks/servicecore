package ifcs

type ILogger interface {
	Log(LogType, string)
}
