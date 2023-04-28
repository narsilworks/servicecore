package ifcs

import (
	sg "github.com/narsilworks/servicecore/str"
)

type IDatabaseInfo interface {
	ConnectionString() string
	DriverName() string
	HelperID() string
	Schema() string
	ConnectionInfo() (maxOpen, maxIdle, maxConLt, maxConnIt int)
	ParameterInfo() (placeHolder string, inSequence bool)
	SQLStringInfo() (enclosing, escape, rwEscape string)
	Sequence() sg.SequenceGeneratorInfo
}
