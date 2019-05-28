package jsonb

import (
	"testing"

	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestEscapeString(t *testing.T) {
	str := `test		\"asdasjkd
	lkasd128378127#$%^&*()_)(*&^%$#@#~!@#$%`
	blr := &bytes.Buffer{}
	escapeString(blr, str)
	assert.Equal(t, `test\t\t\\\"asdasjkd\n\tlkasd128378127#$%^&*()_)(*&^%$#@#~!@#$%`, blr.String(), "unexpected result")
}
