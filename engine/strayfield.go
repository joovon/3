package engine

// Calculation of stray field

import (
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
	"github.com/mumax/3/mag"
)

// stray variables
var (
	B_strayfield     = NewVectorField("B_strayfield", "T", "Magnetostatic stray field", SetStrayField)
	StrayFieldLift    inputValue
)

func init() {

	StrayFieldLift = numParam(50e-9, "StrayFieldLift", "m", reinitmfmconv)
	DeclLValue("StrayFieldLift", &StrayFieldLift, "Stray field lift height")
}

// Sets dst to the current demag field
func SetStrayField(dst *data.Slice) {
	if EnableDemag {
		msat := Msat.MSlice()
		defer msat.Recycle()
		if NoDemagSpins.isZero() {
			// Normal demag, everywhere
			strayfieldConv().Exec(dst, M.Buffer(), geometry.Gpu(), msat)
		} else {
			setMaskedStrayField(dst, msat)
		}
	} else {
		cuda.Zero(dst) // will ADD other terms to it
	}
}

// Sets dst to the demag field, but cells where NoDemagSpins != 0 do not generate nor recieve field.
func setMaskedStrayField(dst *data.Slice, msat cuda.MSlice) {
	// No-demag spins: mask-out geometry with zeros where NoDemagSpins is set,
	// so these spins do not generate a field

	buf := cuda.Buffer(SCALAR, geometry.Gpu().Size()) // masked-out geometry
	defer cuda.Recycle(buf)

	// obtain a copy of the geometry mask, which we can overwrite
	geom, r := geometry.Slice()
	if r {
		defer cuda.Recycle(geom)
	}
	data.Copy(buf, geom)

	// mask-out
	cuda.ZeroMask(buf, NoDemagSpins.gpuLUT1(), regions.Gpu())

	// convolution with masked-out cells.
	strayfieldConv().Exec(dst, M.Buffer(), buf, msat)

	// After convolution, mask-out the field in the NoDemagSpins cells
	// so they don't feel the field generated by others.
	cuda.ZeroMask(dst, NoDemagSpins.gpuLUT1(), regions.Gpu())
}


// returns stray field convolution, making sure it's initialized
func strayfieldConv() *cuda.StrayFieldConvolution {
	if conv_ == nil {
		SetBusy(true)
		defer SetBusy(false)
		kernel := mag.StrayFieldKernel(Mesh().Size(), Mesh().PBC(), Mesh().CellSize(), StrayFieldLift, *Flag_cachedir)
		conv_ = cuda.NewStrayField(Mesh().Size(), Mesh().PBC(), kernel, *Flag_selftest)
	}
	return conv_
}