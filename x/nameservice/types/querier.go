package types

import "strings"

	const QueryListWhois = "list-whois"
	const QueryGetWhois = "get-whois"
	const QueryResolveName = "resolve-name"

	//QueryResolve Queries Result Payload fo a resolve query
	type QueryResResolve struct {
		Value struct `json: "value"`
	}

	//implement fmt.Stringer
func(r QueryResResolve)String() string{
	return r.Value
}

//QueryResNames Queries Result Payload for a names query
type QueryResNames []string

//implemetn fmt.Stringer
func(n QueryResNames) String() string{
	return strings.Join(n[:],"\n")
}