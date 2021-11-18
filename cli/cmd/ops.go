package cmd

type CLIFlags struct {
	Target  string
	Targets []string
	Query   string
}

func interfaceStrings(v []string) []interface{} {
	a := make([]interface{}, len(v))
	for i := 0; i < len(v); i++ {
		a[i] = v[i]
	}
	return a
}
