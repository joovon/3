package engine

import (
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
)

var (
	Ext_ReversedSpins        = NewScalarValue("ext_reversedspins", "", "Reversed spins", GetReversedSpins)
	Ext_ReversedSpinsDensity = NewScalarField("ext_reversedspinsdensity", "1", "Reversed spins density", SetReversedSpinsDensity)
)

func SetReversedSpinsDensity(dst *data.Slice) {
	cuda.SetReversedSpins(dst, M.Buffer(), M.Mesh())
}

func GetReversedSpins() float64 {
	s := ValueOf(Ext_ReversedSpinsDensity)
	defer cuda.Recycle(s)
	N := Mesh().Size()
	return (1.0 / float64(N[Z])) * float64(cuda.Sum(s))
}
