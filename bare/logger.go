package bare

import (
	"log"

	"github.com/narsilworks/servicecore/ifcs"
)

type StdOutLog struct {
}

func (s *StdOutLog) Log(lt ifcs.LogType, msg string) {
	lts := string(lt)
	bs := "\b"
	if lts != "" {
		lts = "[" + lts + "]"
		bs = ""
	}
	log.Printf("%s%s: %s\r\n", bs, lts, msg)
}
