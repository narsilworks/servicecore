package ifcs

type LogType string

const (
	LOGTYPE_APP = LogType(``)
	LOGTYPE_INF = LogType(`INF`)
	LOGTYPE_WRN = LogType(`WRN`)
	LOGTYPE_ERR = LogType(`ERR`)
)
