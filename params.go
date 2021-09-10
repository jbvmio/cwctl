package cwctl

import "net/url"

// Parameters define the available parameters for use when making an endpoint call to the CW API.
type Parameters struct {
	PageSize      string
	Page          string
	OrderBy       string
	IncludeFields string
	Ids           string
	Expand        string
	ExcludeFields string
	Condition     string
}

// Build contructs and returns Params based the values contained within.
func (P *Parameters) Build() Params {
	params := make(Params, len(totalParams))
	params[pidPageSize] = totalParams[pidPageSize](P.PageSize)
	params[pidPage] = totalParams[pidPage](P.Page)
	params[pidOrderBy] = totalParams[pidOrderBy](P.OrderBy)
	params[pidIncludeFields] = totalParams[pidIncludeFields](P.IncludeFields)
	params[pidIDs] = totalParams[pidIDs](P.Ids)
	params[pidExpand] = totalParams[pidExpand](P.Expand)
	params[pidExcludeFields] = totalParams[pidExcludeFields](P.ExcludeFields)
	params[pidCondition] = totalParams[pidCondition](P.Condition)
	return params
}

// Params is collection of ParamFuncs.
type Params []ParamFn

// Encode returns the collection of parameters url encoded.
func (p Params) Encode() string {
	params := url.Values{}
	for i := 0; i < len(p); i++ {
		k, v := p[i]()
		if v != "" {
			params.Add(k, v)
		}
	}
	e := params.Encode()
	if e != "" {
		e = `?` + e
	}
	return e
}

const (
	pidPageSize = iota
	pidPage
	pidOrderBy
	pidIncludeFields
	pidIDs
	pidExpand
	pidExcludeFields
	pidCondition
)

var totalParams = [...]makeParamFn{
	PageSize,
	Page,
	OrderBy,
	IncludeFields,
	Ids,
	Expand,
	ExcludeFields,
	Condition,
}

type makeParamFn func(string) ParamFn

// ParamFn returns a key value pair used as parameters for an endpoint call to the CW API.
type ParamFn func() (string, string)

func PageSize(v string) ParamFn {
	return func() (string, string) { return `pagesize`, v }
}

func Page(v string) ParamFn {
	return func() (string, string) { return `page`, v }
}

func OrderBy(v string) ParamFn {
	return func() (string, string) { return `orderby`, v }
}

func IncludeFields(v string) ParamFn {
	return func() (string, string) { return `includefields`, v }
}

func Ids(v string) ParamFn {
	return func() (string, string) { return `ids`, v }
}

func Expand(v string) ParamFn {
	return func() (string, string) { return `expand`, v }
}

func ExcludeFields(v string) ParamFn {
	return func() (string, string) { return `excludefields`, v }
}

func Condition(v string) ParamFn {
	return func() (string, string) { return `condition`, v }
}
