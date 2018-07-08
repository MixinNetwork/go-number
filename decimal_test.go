package number

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("0", Zero().Persist())
	assert.True(FromString("0.00000000000000100000000000000003").Equal(FromString("0.000000000000001000000000000000031")))
	assert.False(FromString("0.00000000000000100000000000000002").Equal(FromString("0.000000000000001000000000000000031")))

	a := FromString("1.0000000000000000300000")
	assert.Equal("1.00000000000000003", a.Persist())

	b := FromString("1.0000000000000000300000")
	assert.Equal("1.00000000000000003", b.Persist())
	assert.Equal("1", b.PresentFloor())
	assert.Equal("1.00000001", b.PresentCeil())

	c := NewDecimal(1e8, 7)
	assert.Equal("10", c.Persist())
	assert.Equal("10", c.PresentFloor())
	assert.Equal("10", c.PresentCeil())

	d := NewDecimal(1, 15)
	assert.Equal("0.000000000000001", d.Persist())
	assert.Equal(1, d.Cmp(Zero()))
	assert.True(d.Exhausted())
	assert.Equal("0.00000000000000100000000000000003", a.Mul(d).Persist())
	assert.Equal("0", a.Mul(d).PresentFloor())
	assert.Equal("0.00000001", a.Mul(d).PresentCeil())
	assert.True(a.Mul(d).Exhausted())
	assert.Equal("1000000000000000.03", a.Div(d).Persist())
	assert.False(a.Div(d).Exhausted())

	e, _ := new(big.Int).SetString("0x32748932FA632BFF8323282", 0)
	assert.Equal("975945861920638759748121218", e.Text(10))

	f := FromString(e.Text(10)).Mul(NewDecimal(1, 18))
	assert.Equal("975945861.920638759748121218", f.Persist())
	assert.Equal("975945861.92063875", f.PresentFloor())
	assert.Equal("975945861.92063876", f.PresentCeil())
	assert.False(f.Exhausted())

	g := NewDecimal(1e10, 0)
	assert.Equal("10000000000", g.Persist())

	h := FromString(e.Text(10)).Div(FromString("1000000000000000000"))
	assert.Equal("975945861.920638759748121218", h.Persist())
	assert.Equal("975945861.92063875", h.PresentFloor())
	assert.Equal("975945861.92063876", h.PresentCeil())
	assert.True(f.Equal(h))

	i := FromString("465.505437")
	assert.Equal("465.50543", i.RoundFloor(5).Persist())
	assert.Equal("465.505437", i.RoundFloor(8).Persist())

	j := FromString("0.00000000999")
	assert.Equal("0", j.RoundFloor(8).Persist())
	assert.True(j.Exhausted())
}
