package number

import (
	"log"
	"math/big"
)

type Integer struct {
	value    *big.Int
	decimals uint8
}

func NewInteger(value int64, decimals uint8) Integer {
	return Integer{big.NewInt(value), decimals}
}

func (i Integer) Value() int64 {
	if !i.value.IsInt64() {
		panic(i.value.String())
	}
	return i.value.Int64()
}

func (i Integer) Precision() uint8 {
	return i.decimals
}

func (i Integer) Zero() Integer {
	return Integer{new(big.Int), i.decimals}
}

func (i Integer) Decimal() Decimal {
	return FromString(i.value.String()).Mul(NewDecimal(1, int32(i.decimals)))
}

func (i Integer) Persist() string {
	return i.Decimal().Persist()
}

func (i Integer) MarshalJSON() ([]byte, error) {
	return []byte("\"" + i.Persist() + "\""), nil
}

func (a Integer) safeAdd(b Integer) (Integer, bool) {
	i := Integer{
		value:    new(big.Int).Add(a.value, b.value),
		decimals: a.decimals,
	}
	if a.decimals != b.decimals {
		return i, false
	}
	if i.value.Cmp(a.value) < 0 || i.value.Cmp(b.value) < 0 {
		return i, false
	}
	return i, true
}

func (a Integer) safeSub(b Integer) (Integer, bool) {
	i := Integer{
		value:    new(big.Int).Sub(a.value, b.value),
		decimals: a.decimals,
	}
	if a.decimals != b.decimals {
		return i, false
	}
	if b.value.Cmp(a.value) > 0 {
		return i, false
	}
	return i, true
}

func (a Integer) safeMul(b Integer) (Integer, bool) {
	i := Integer{
		value:    new(big.Int).Mul(a.value, b.value),
		decimals: a.decimals + b.decimals,
	}
	if i.decimals < a.decimals || i.decimals < b.decimals {
		return i, false
	}
	if a.value.Sign() != 0 && new(big.Int).Div(i.value, a.value).Cmp(b.value) != 0 {
		return i, false
	}
	return i, true
}

func (a Integer) safeDiv(b Integer) (Integer, bool) {
	i := Integer{
		value:    new(big.Int).Div(a.value, b.value),
		decimals: a.decimals - b.decimals,
	}
	if a.decimals < b.decimals {
		return i, false
	}
	return i, true
}

func (a Integer) safeCmp(b Integer) (int, bool) {
	if a.decimals != b.decimals {
		return 0, false
	}
	if a.value.Cmp(b.value) > 0 {
		return 1, true
	}
	if a.value.Cmp(b.value) < 0 {
		return -1, true
	}
	return 0, true
}

func (a Integer) Add(b Integer) Integer {
	i, ok := a.safeAdd(b)
	if !ok {
		log.Panicln(a, b)
	}
	return i
}

func (a Integer) Sub(b Integer) Integer {
	i, ok := a.safeSub(b)
	if !ok {
		log.Panicln(a, b)
	}
	return i
}

func (a Integer) Mul(b Integer) Integer {
	i, ok := a.safeMul(b)
	if !ok {
		log.Panicln(a, b)
	}
	return i
}

func (a Integer) Div(b Integer) Integer {
	i, ok := a.safeDiv(b)
	if !ok {
		log.Panicln(a, b)
	}
	return i
}

func (a Integer) Cmp(b Integer) int {
	i, ok := a.safeCmp(b)
	if !ok {
		log.Panicln(a, b)
	}
	return i
}

func (a Integer) IsZero() bool {
	return a.value.Sign() == 0
}

func (a Integer) IsPositive() bool {
	return a.value.Sign() > 0
}

func (a Integer) IsNegative() bool {
	return a.value.Sign() < 0
}
