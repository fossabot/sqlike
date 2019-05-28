package mysql

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Format :
func (ms MySQL) Format(it interface{}) (val string) {
	switch vi := it.(type) {
	case []byte:
		val = strconv.Quote(string(vi))
	case string:
		val = strconv.Quote(vi)
	case bool:
		val = "0"
		if vi {
			val = "1"
		}
	case int64:
		val = strconv.FormatInt(vi, 10)
	case uint64:
		val = strconv.FormatUint(vi, 10)
	case float64:
		val = strconv.FormatFloat(vi, 'e', -1, 64)
	case time.Time:
		val = vi.Format(`'2006-01-02 15:04:05.999999'`)
	case json.RawMessage:
		val = strconv.Quote(string(vi))
	case nil:
		val = "NULL"
	case fmt.Stringer:
		val = vi.String()
	}
	return
}
