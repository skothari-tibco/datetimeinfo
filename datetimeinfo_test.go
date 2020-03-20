package datetimeinfo

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestDay(t *testing.T) {

	d := DateTimeInfo{}
	num, err := d.Eval("10:11:05.00000", "hour")
	assert.Nil(t, err)
	assert.Equal(t, 10, num)

	num, err = d.Eval("2017/08/02", "day")
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

}
