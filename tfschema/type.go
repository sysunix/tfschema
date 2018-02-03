package tfschema

import (
	"fmt"
	"reflect"
	"strings"
)

type Type struct {
	// T is an instance of github.com/hashicorp/terraform/vendor/github.com/zclconf/go-cty.Type
	ctyType interface{}
}

func NewType(t interface{}) *Type {
	return &Type{
		ctyType: t,
	}
}

func (t *Type) MarshalJSON() ([]byte, error) {
	v := reflect.ValueOf(t.ctyType).MethodByName("GoString")
	if !v.IsValid() {
		return nil, fmt.Errorf("Faild to find GoString(): %#v", t)
	}

	nv := v.Call([]reflect.Value{})
	if len(nv) == 0 {
		return nil, fmt.Errorf("Faild to call GoString(): %#v", v)
	}

	name := nv[0].String()
	pretty := strings.ToLower(strings.Replace(name, "cty.", "", -1))

	return []byte(`"` + pretty + `"`), nil
}