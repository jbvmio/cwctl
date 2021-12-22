package cmd

import (
	"encoding/base64"
	"fmt"
)

const (
	encPwshFileCMD = "[System.IO.File]::WriteAllBytes(\"%s\", [System.Convert]::FromBase64String(\"%s\"))"
	encBashFileCMD = "echo \"%s\" | base64 -d > \"%s\""
)

func encodedFileCommand(file []byte, filename string, pwsh bool) (fileCmd string) {
	enc := base64.StdEncoding.EncodeToString(file)
	scrFmt := encBashFileCMD
	switch {
	case pwsh:
		scrFmt = encPwshFileCMD
		fileCmd = fmt.Sprintf(scrFmt, filename, enc)
	default:
		fileCmd = fmt.Sprintf(scrFmt, enc, filename)
	}
	return
}
