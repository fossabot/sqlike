package expr

import (
	"github.com/si3nloong/sqlike/sqlike/primitive"
)

// JSONQuote :
func JSONQuote(value string) (fc primitive.Func) {
	fc.Type = primitive.JSONQuote
	fc.Arguments = append(fc.Arguments, wrapColumn(value))
	return
}

// JSONContains :
func JSONContains(target, candidate interface{}, paths ...string) (jc primitive.JC) {
	var path *string
	if len(paths) > 0 {
		path = &paths[0]
	}
	switch vi := target.(type) {
	case string:
		jc.Target = primitive.Column{Name: vi}
	case primitive.Column:
		jc.Target = vi
	default:
		jc.Target = primitive.Value{Raw: vi}
	}
	jc.Candidate = wrapJSONColumn(candidate)
	jc.Path = path
	return
}

func wrapJSONColumn(it interface{}) interface{} {
	switch vi := it.(type) {
	case primitive.Column:
		return primitive.CastAs{
			Value:    vi,
			DataType: primitive.JSON,
		}
	case primitive.Func:
		return vi
	default:
		return primitive.Value{Raw: vi}
	}
}
