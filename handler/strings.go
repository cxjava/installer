package handler

import (
	"strings"
)

func getOS(s string) string {
	s = strings.ToLower(s)
	var result string
	switch {
	case osReDarwin.MatchString(s):
		result = "darwin"
	case osReDragonfly.MatchString(s):
		result = "dragonfly"
	case osReWindows.MatchString(s):
		result = "windows"
	case osReMisc.MatchString(s):
		// return the first capturing group (contains only the alphanumeric characters)
		matches := osReMisc.FindStringSubmatch(s)
		if len(matches) > 1 {
			result = matches[1]
		}
	default:
		result = ""
	}
	return result
}

func getArch(s string) string {
	s = strings.ToLower(s)
	var result string
	switch {
	case archReLoong64.MatchString(s):
		result = "loong64"
	case archRePPC64.MatchString(s):
		result = "ppc64"
	case archRePPC64LE.MatchString(s):
		result = "ppc64le"
	case archReRiscv64.MatchString(s):
		result = "riscv64"
	case archReArm64.MatchString(s):
		result = "arm64"
	case archReAmd64.MatchString(s):
		result = "amd64"
	case archReArm.MatchString(s):
		result = "arm"
	case archRe386.MatchString(s):
		result = "386"
	case archReMisc.MatchString(s):
		matches := archReMisc.FindStringSubmatch(s)
		if len(matches) > 1 {
			result = matches[1]
		}
	// fuzz match 'x?64(bit)?'
	case fuzzArchAmd64.MatchString(s):
		result = "amd64"
	// fuzz match 'x?32(bit)?'
	case fuzzArch386.MatchString(s):
		result = "386"
	default:
		result = ""
	}
	return result
}

func getFileExt(s string) string {
	return fileExtRe.FindString(s)
}

func splitHalf(s, by string) (string, string) {
	i := strings.Index(s, by)
	if i == -1 {
		return s, ""
	}
	return s[:i], s[i+len(by):]
}
