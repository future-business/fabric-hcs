package hedera

import (
	"fmt"
	"math"
)

type Hbar struct {
	tinybar int64
}

var MaxHbar = Hbar{math.MaxInt64}

var MinHbar = Hbar{math.MinInt64}

var ZeroHbar = Hbar{0}

// HbarFrom creates a representation of Hbar in tinybar on the unit provided
func HbarFrom(bars float64, unit HbarUnit) Hbar {
	return HbarFromTinybar(int64(bars * float64(unit.numberOfTinybar())))
}

// HbarFromTinybar creates a representation of Hbar in tinybars
func HbarFromTinybar(tinybar int64) Hbar {
	return Hbar{tinybar}
}

func NewHbar(hbar float64) Hbar {
	return HbarFrom(hbar, HbarUnits.Hbar)
}

func (hbar Hbar) AsTinybar() int64 {
	return hbar.tinybar
}

func (hbar Hbar) As(unit HbarUnit) int64 {
	return hbar.tinybar * unit.numberOfTinybar()
}

func (hbar Hbar) String() string {
	// Format the string as tinybar if the value is 1000 tinybar or less
	if hbar.tinybar <= 1000 {
		return fmt.Sprintf("%v tħ", hbar.tinybar)
	}

	return fmt.Sprintf("%v ℏ", float64(hbar.tinybar)/float64(HbarUnits.Hbar.numberOfTinybar()))
}

func (hbar Hbar) negated() Hbar {
	return Hbar{
		tinybar: -hbar.tinybar,
	}
}
