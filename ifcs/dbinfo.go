package ifcs

import strct "github.com/narsilworks/servicecore/strct"

type IDatabaseInfo interface {
	ConnectionString() string
	DriverName() string
	HelperID() string
	Schema() string
	ConnectionInfo() (maxOpen, maxIdle, maxConLt, maxConnIt int)
	ParameterInfo() (placeHolder string, inSequence bool)
	SQLStringInfo() (enclosing, escape, rwEscape string)
	Sequence() strct.SequenceGeneratorInfo
}
