package reflext

import (
	"reflect"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

type dbStruct struct {
	Name  string `db:"name"`
	skip  bool
	Email *string
}

func TestMapper(t *testing.T) {
	var (
		mapper = NewMapperFunc("db", func(str string) string {
			return str
		})
		ok bool
	)

	tmp := dbStruct{Name: "John"}
	v := reflect.ValueOf(&tmp)
	fv := mapper.FieldByName(v, "name")
	require.NotNil(t, fv)

	require.Panics(t, func() {
		mapper.FieldByName(reflect.ValueOf(0), "unknown")
		mapper.FieldByName(reflect.ValueOf(""), "unknown")
	})
	require.Panics(t, func() {
		mapper.FieldByName(v, "unknown")
	})

	// FieldByIndexesReadOnly will not initialise the field even if it's nil
	{
		fv := mapper.FieldByIndexesReadOnly(v, []int{0})
		require.Equal(t, reflect.String, fv.Kind())
		require.Equal(t, "John", fv.Interface().(string))

		fv = mapper.FieldByIndexesReadOnly(v, []int{2})
		require.Nil(t, fv.Interface())

		require.Panics(t, func() {
			mapper.FieldByIndexesReadOnly(v, []int{1000000})
		})
	}

	// FieldByIndexes will initialise if the field is nil
	{
		fv := mapper.FieldByIndexes(v, []int{2})
		require.NotNil(t, fv.Interface())
		require.Equal(t, "", fv.Elem().Interface().(string))

		require.Panics(t, func() {
			mapper.FieldByIndexes(v, []int{1000000})
		})
	}

	{
		fv, ok := mapper.LookUpFieldByName(v, "name")
		require.True(t, ok)
		require.Equal(t, "John", fv.Interface().(string))

		fv, ok = mapper.LookUpFieldByName(v, "unknown")
		require.False(t, ok)
		require.Equal(t, v.Elem(), fv)
	}

	codec := mapper.CodecByType(v.Type())

	// lookup an existed field
	{
		_, ok = codec.LookUpFieldByName("name")
		require.True(t, ok)
	}

	// lookup unexists field
	{
		_, ok = codec.LookUpFieldByName("Unknown")
		require.False(t, ok)
	}

	// lookup private field
	{
		_, ok = codec.LookUpFieldByName("skip")
		require.False(t, ok)
	}

	m := NewMapperFunc("sqlike", nil)
	// go func() {
	type A struct {
		Name    string
		Address string
	}
	type B struct{}
	type C struct{}
	type D struct{}
	type E struct{ ID string }
	type F struct{ ID string }
	type G struct{ ID string }
	type H struct{ ID string }
	type Z struct{ ID string }
	i := 0
	datatypes := []interface{}{
		time.Time{},
		language.Malay,
		language.English,
		civil.DateOf(time.Now()),
		civil.DateTime{},
		A{},
		B{},
		C{},
		D{},
		E{},
		F{},
		G{},
		H{},
		Z{},
	}

	for {

		if len(datatypes) <= i {
			i = 0
		}

		time.Sleep(time.Second * 2)
		typ := reflect.TypeOf(datatypes[i])

		m.CodecByType(typ)
		i++

	}

	// }()
}
