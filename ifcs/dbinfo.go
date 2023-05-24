package ifcs

import "github.com/narsilworks/servicecore/dto"

type IDatabaseInfo interface {
	ConnectionString() string
	DriverName() string
	HelperID() string
	Schema() string
	ConnectionInfo() (maxOpen, maxIdle, maxConLt, maxConnIt int)
	ParameterInfo() (placeHolder string, inSequence bool)
	SQLStringInfo() (enclosing, escape, rwEscape string)
	Sequence() dto.SequenceGeneratorInfo
}
