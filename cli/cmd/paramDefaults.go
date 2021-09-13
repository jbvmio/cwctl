package cmd

import "github.com/jbvmio/cwctl/connectwise"

type flags connectwise.Parameters

func (f *flags) merge(params *connectwise.Parameters) *connectwise.Parameters {
	F := connectwise.Parameters(*f)
	if params == nil {
		return &F
	}
	F.PageSize = params.PageSize
	F.Page = params.Page
	F.OrderBy = params.OrderBy
	F.IncludeFields = params.IncludeFields
	F.Ids = params.Ids
	F.Expand = params.Expand
	F.ExcludeFields = params.ExcludeFields
	F.Condition = params.Condition
	return &F
}

var paramsClient = flags{}

var paramsComputer = flags{
	ExcludeFields: `address,irq`,
}
