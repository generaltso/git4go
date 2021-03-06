// +build windows

package git4go

func guessSystemFile() []string {
	return []string{}
}

func guessGlobalFile() []string {
	return []string{}
}

func guessXDGFile() []string {
	return []string{}
}

func guessTemplateFile() []string {
	return []string{}
}

var defaultBoolConfig map[string]bool = map[string]bool{
	"core.symlinks":          true,
	"core.ignorecase":        true,
	"core.filemode":          true,
	"core.ignorestat":        false,
	"core.trustctime":        true,
	"core.abbrev":            true,
	"core.precomposeunicode": true,
	"core.logallrefupdates":  true,
	"core.protectHFS":        false,
	"core.protectNTFS":       true,
}

var defaultStringConfig map[string]string = map[string]string{
	"core.autocrlf": "false",
	"core.eol":      "crlf",
}
