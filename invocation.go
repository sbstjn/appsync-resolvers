package resolvers

import "encoding/json"

type context struct {
	Arguments json.RawMessage `json:"arguments"`
	Source    json.RawMessage `json:"source"`
	Identity  *json.RawMessage `json:"identity"`
}

type invocation struct {
	Resolve string  `json:"resolve"`
	Context context `json:"context"`
}

func (in invocation) isRoot() bool {
	return in.Context.Source == nil || string(in.Context.Source) == "null"
}

func (in invocation) payload() json.RawMessage {
	if in.isRoot() {
		return in.Context.Arguments
	}

	return in.Context.Source
}

func (in invocation) identity() *json.RawMessage {
	return in.Context.Identity
}
