package exepath

import "path/filepath"
import "strings"

// By default, contains the lowercase basename of Abs with any file extension stripped.
//
// This can be changed by any configuration code which knows better. It may be
// used e.g. as the application name for syslog.
var ProgramName string

// Used to track what set the program name.
var ProgramNameSetter = "default"

func initProgramName() {
	b := filepath.Base(Abs)
	ProgramName = strings.ToLower(b[0 : len(b)-len(filepath.Ext(b))])
}
