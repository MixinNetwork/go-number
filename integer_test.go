package number

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInteger(t *testing.T) {
	assert := assert.New(t)
	i := NewInteger(123, 4)
	assert.Equal("0.0123", i.Persist())
	a := FromString("60000").Integer(4)
	assert.Equal("60000", a.Persist())
	b := FromString("6").Integer(8)
	assert.Equal("6", b.Persist())
	assert.Equal("0.0001", b.Div(a).Persist())
	assert.Equal(big.NewInt(600000000), a.value)
	assert.Equal(big.NewInt(600000000), b.value)
}
