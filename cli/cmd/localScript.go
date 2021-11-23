package cmd

import (
	"encoding/base64"
	"fmt"
)

const (
	encPwshScriptCMD = "iex $([System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String(\"%s\")))"
	encBashScriptCMD = "echo \"%s\" | base64 -d | bash"
)

func encodedScriptCommand(script []byte, pwsh bool) (scriptCmd string) {
	scrFmt := encBashScriptCMD
	if pwsh {
		scrFmt = encPwshScriptCMD
	}
	enc := base64.StdEncoding.EncodeToString(script)
	scriptCmd = fmt.Sprintf(scrFmt, enc)
	return
}
