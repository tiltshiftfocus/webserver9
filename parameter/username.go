package parameter
import (
	"regexp"
)

type Username struct { 
	Value string
	Valid bool
}
func (self *Username) Set(s *string) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]+$", *s); ok {
		self.Valid = true
		self.Value = *s
	}
}

