package dag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_warpMapValueToStr(t *testing.T) {
	res, err := warpMapValueToStr(DagContext{
		"size": map[string]any{
			"key1": "value1",
			"key2": 2,
		},
		"value": 100,
		"pi":    3.14,
	}, "this is a ${{size}},  ${{value}} > ${{pi}}")
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, `this is a {"key1":"value1","key2":2},  100 > 3.14`, res)
}
