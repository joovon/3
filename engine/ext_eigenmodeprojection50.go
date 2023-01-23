package engine

// **************************************** 
// Author(s): Joo-Von Kim, C2N, CNRS/Univ. Paris-Saclay
//
// This module projects the magnetization onto user-supplied transverse directions,
// delta_mx, delta_my (obtained, e.g., from the relaxed magnetization state), with
// a spatial convolution with the user-supplied masks psi_kn. It returns the amplitudes
// 
// 	a_kxn = int_dV {psi_kn (m . delta_mx)}
//	a_kyn = int_dV {psi_kn (m . delta_my)}
//
// with up to 51 modes (n = 0 to 50).
//
// The user-supplied masks/vector fields can be added in the source .mx3 file with
//	psi_kn.Add( LoadFile(("psi_kn_file.ovf"),1) )	
//	delta_mx.Add( LoadFile("delta_mx_file.ovf"), 1 )
//	etc.
//
// Acknowledgements:
// This work was supported by Horizon Europe Research and Innovation Programme of the
// European Commission under grant agreement No. 101070290 (NIMFEIA). 
//
// **************************************** 


import (
	"github.com/mumax/3/cuda"
)


var (
//  Transverse magnetization is already defined in ext_eigenmodeprojection.go
//	delta_mx	= NewExcitation("delta_mx", "", "Transverse magnetization 1")
//	delta_my	= NewExcitation("delta_my", "", "Transverse magnetization 2")

	psi_k00		= NewScalarExcitation("psi_k00", "", "Eigenmode spatial profile")
	psi_k01		= NewScalarExcitation("psi_k01", "", "Eigenmode spatial profile")
	psi_k02		= NewScalarExcitation("psi_k02", "", "Eigenmode spatial profile")
	psi_k03		= NewScalarExcitation("psi_k03", "", "Eigenmode spatial profile")
	psi_k04		= NewScalarExcitation("psi_k04", "", "Eigenmode spatial profile")
	psi_k05		= NewScalarExcitation("psi_k05", "", "Eigenmode spatial profile")
	psi_k06		= NewScalarExcitation("psi_k06", "", "Eigenmode spatial profile")
	psi_k07		= NewScalarExcitation("psi_k07", "", "Eigenmode spatial profile")
	psi_k08		= NewScalarExcitation("psi_k08", "", "Eigenmode spatial profile")
	psi_k09		= NewScalarExcitation("psi_k09", "", "Eigenmode spatial profile")
	psi_k10		= NewScalarExcitation("psi_k10", "", "Eigenmode spatial profile")
	psi_k11		= NewScalarExcitation("psi_k11", "", "Eigenmode spatial profile")
	psi_k12		= NewScalarExcitation("psi_k12", "", "Eigenmode spatial profile")
	psi_k13		= NewScalarExcitation("psi_k13", "", "Eigenmode spatial profile")
	psi_k14		= NewScalarExcitation("psi_k14", "", "Eigenmode spatial profile")
	psi_k15		= NewScalarExcitation("psi_k15", "", "Eigenmode spatial profile")
	psi_k16		= NewScalarExcitation("psi_k16", "", "Eigenmode spatial profile")
	psi_k17		= NewScalarExcitation("psi_k17", "", "Eigenmode spatial profile")
	psi_k18		= NewScalarExcitation("psi_k18", "", "Eigenmode spatial profile")
	psi_k19		= NewScalarExcitation("psi_k19", "", "Eigenmode spatial profile")
	psi_k20		= NewScalarExcitation("psi_k20", "", "Eigenmode spatial profile")
	psi_k21		= NewScalarExcitation("psi_k21", "", "Eigenmode spatial profile")
	psi_k22		= NewScalarExcitation("psi_k22", "", "Eigenmode spatial profile")
	psi_k23		= NewScalarExcitation("psi_k23", "", "Eigenmode spatial profile")
	psi_k24		= NewScalarExcitation("psi_k24", "", "Eigenmode spatial profile")
	psi_k25		= NewScalarExcitation("psi_k25", "", "Eigenmode spatial profile")
	psi_k26		= NewScalarExcitation("psi_k26", "", "Eigenmode spatial profile")
	psi_k27		= NewScalarExcitation("psi_k27", "", "Eigenmode spatial profile")
	psi_k28		= NewScalarExcitation("psi_k28", "", "Eigenmode spatial profile")
	psi_k29		= NewScalarExcitation("psi_k29", "", "Eigenmode spatial profile")
	psi_k30		= NewScalarExcitation("psi_k30", "", "Eigenmode spatial profile")
	psi_k31		= NewScalarExcitation("psi_k31", "", "Eigenmode spatial profile")
	psi_k32		= NewScalarExcitation("psi_k32", "", "Eigenmode spatial profile")
	psi_k33		= NewScalarExcitation("psi_k33", "", "Eigenmode spatial profile")
	psi_k34		= NewScalarExcitation("psi_k34", "", "Eigenmode spatial profile")
	psi_k35		= NewScalarExcitation("psi_k35", "", "Eigenmode spatial profile")
	psi_k36		= NewScalarExcitation("psi_k36", "", "Eigenmode spatial profile")
	psi_k37		= NewScalarExcitation("psi_k37", "", "Eigenmode spatial profile")
	psi_k38		= NewScalarExcitation("psi_k38", "", "Eigenmode spatial profile")
	psi_k39		= NewScalarExcitation("psi_k39", "", "Eigenmode spatial profile")
	psi_k40		= NewScalarExcitation("psi_k40", "", "Eigenmode spatial profile")
	psi_k41		= NewScalarExcitation("psi_k41", "", "Eigenmode spatial profile")
	psi_k42		= NewScalarExcitation("psi_k42", "", "Eigenmode spatial profile")
	psi_k43		= NewScalarExcitation("psi_k43", "", "Eigenmode spatial profile")
	psi_k44		= NewScalarExcitation("psi_k44", "", "Eigenmode spatial profile")
	psi_k45		= NewScalarExcitation("psi_k45", "", "Eigenmode spatial profile")
	psi_k46		= NewScalarExcitation("psi_k46", "", "Eigenmode spatial profile")
	psi_k47		= NewScalarExcitation("psi_k47", "", "Eigenmode spatial profile")
	psi_k48		= NewScalarExcitation("psi_k48", "", "Eigenmode spatial profile")
	psi_k49		= NewScalarExcitation("psi_k49", "", "Eigenmode spatial profile")
	psi_k50		= NewScalarExcitation("psi_k50", "", "Eigenmode spatial profile")
	
			
	a_k00		= NewVectorValue("a_k00", "", "delta_mx projection onto psi_k00", GetModeAmplitudek00 )
	a_k01		= NewVectorValue("a_k01", "", "delta_my projection onto psi_k01", GetModeAmplitudek01 )
	a_k02		= NewVectorValue("a_k02", "", "delta_my projection onto psi_k02", GetModeAmplitudek02 )
	a_k03		= NewVectorValue("a_k03", "", "delta_my projection onto psi_k03", GetModeAmplitudek03 )
	a_k04		= NewVectorValue("a_k04", "", "delta_my projection onto psi_k04", GetModeAmplitudek04 )
	a_k05		= NewVectorValue("a_k05", "", "delta_my projection onto psi_k05", GetModeAmplitudek05 )
	a_k06		= NewVectorValue("a_k06", "", "delta_my projection onto psi_k06", GetModeAmplitudek06 )
	a_k07		= NewVectorValue("a_k07", "", "delta_my projection onto psi_k07", GetModeAmplitudek07 )
	a_k08		= NewVectorValue("a_k08", "", "delta_my projection onto psi_k08", GetModeAmplitudek08 )
	a_k09		= NewVectorValue("a_k09", "", "delta_my projection onto psi_k09", GetModeAmplitudek09 )
	a_k10		= NewVectorValue("a_k10", "", "delta_my projection onto psi_k10", GetModeAmplitudek10 )
	a_k11		= NewVectorValue("a_k11", "", "delta_my projection onto psi_k11", GetModeAmplitudek11 )
	a_k12		= NewVectorValue("a_k12", "", "delta_my projection onto psi_k12", GetModeAmplitudek12 )
	a_k13		= NewVectorValue("a_k13", "", "delta_my projection onto psi_k13", GetModeAmplitudek13 )
	a_k14		= NewVectorValue("a_k14", "", "delta_my projection onto psi_k14", GetModeAmplitudek14 )
	a_k15		= NewVectorValue("a_k15", "", "delta_my projection onto psi_k15", GetModeAmplitudek15 )
	a_k16		= NewVectorValue("a_k16", "", "delta_my projection onto psi_k16", GetModeAmplitudek16 )
	a_k17		= NewVectorValue("a_k17", "", "delta_my projection onto psi_k17", GetModeAmplitudek17 )
	a_k18		= NewVectorValue("a_k18", "", "delta_my projection onto psi_k18", GetModeAmplitudek18 )
	a_k19		= NewVectorValue("a_k19", "", "delta_my projection onto psi_k19", GetModeAmplitudek19 )
	a_k20		= NewVectorValue("a_k20", "", "delta_my projection onto psi_k20", GetModeAmplitudek20 )
	a_k21		= NewVectorValue("a_k21", "", "delta_my projection onto psi_k01", GetModeAmplitudek21 )
	a_k22		= NewVectorValue("a_k22", "", "delta_my projection onto psi_k02", GetModeAmplitudek22 )
	a_k23		= NewVectorValue("a_k23", "", "delta_my projection onto psi_k03", GetModeAmplitudek23 )
	a_k24		= NewVectorValue("a_k24", "", "delta_my projection onto psi_k04", GetModeAmplitudek24 )
	a_k25		= NewVectorValue("a_k25", "", "delta_my projection onto psi_k05", GetModeAmplitudek25 )
	a_k26		= NewVectorValue("a_k26", "", "delta_my projection onto psi_k06", GetModeAmplitudek26 )
	a_k27		= NewVectorValue("a_k27", "", "delta_my projection onto psi_k07", GetModeAmplitudek27 )
	a_k28		= NewVectorValue("a_k28", "", "delta_my projection onto psi_k08", GetModeAmplitudek28 )
	a_k29		= NewVectorValue("a_k29", "", "delta_my projection onto psi_k09", GetModeAmplitudek29 )
	a_k30		= NewVectorValue("a_k30", "", "delta_my projection onto psi_k10", GetModeAmplitudek30 )
	a_k31		= NewVectorValue("a_k31", "", "delta_my projection onto psi_k11", GetModeAmplitudek31 )
	a_k32		= NewVectorValue("a_k32", "", "delta_my projection onto psi_k12", GetModeAmplitudek32 )
	a_k33		= NewVectorValue("a_k33", "", "delta_my projection onto psi_k13", GetModeAmplitudek33 )
	a_k34		= NewVectorValue("a_k34", "", "delta_my projection onto psi_k14", GetModeAmplitudek34 )
	a_k35		= NewVectorValue("a_k35", "", "delta_my projection onto psi_k15", GetModeAmplitudek35 )
	a_k36		= NewVectorValue("a_k36", "", "delta_my projection onto psi_k16", GetModeAmplitudek36 )
	a_k37		= NewVectorValue("a_k37", "", "delta_my projection onto psi_k17", GetModeAmplitudek37 )
	a_k38		= NewVectorValue("a_k38", "", "delta_my projection onto psi_k18", GetModeAmplitudek38 )
	a_k39		= NewVectorValue("a_k39", "", "delta_my projection onto psi_k19", GetModeAmplitudek39 )
	a_k40		= NewVectorValue("a_k40", "", "delta_my projection onto psi_k20", GetModeAmplitudek40 )
	a_k41		= NewVectorValue("a_k41", "", "delta_my projection onto psi_k11", GetModeAmplitudek41 )
	a_k42		= NewVectorValue("a_k42", "", "delta_my projection onto psi_k12", GetModeAmplitudek42 )
	a_k43		= NewVectorValue("a_k43", "", "delta_my projection onto psi_k13", GetModeAmplitudek43 )
	a_k44		= NewVectorValue("a_k44", "", "delta_my projection onto psi_k14", GetModeAmplitudek44 )
	a_k45		= NewVectorValue("a_k45", "", "delta_my projection onto psi_k15", GetModeAmplitudek45 )
	a_k46		= NewVectorValue("a_k46", "", "delta_my projection onto psi_k16", GetModeAmplitudek46 )
	a_k47		= NewVectorValue("a_k47", "", "delta_my projection onto psi_k17", GetModeAmplitudek47 )
	a_k48		= NewVectorValue("a_k48", "", "delta_my projection onto psi_k18", GetModeAmplitudek48 )
	a_k49		= NewVectorValue("a_k49", "", "delta_my projection onto psi_k19", GetModeAmplitudek49 )
	a_k50		= NewVectorValue("a_k50", "", "delta_my projection onto psi_k20", GetModeAmplitudek50 )
	
)
	
	
func GetModeAmplitudek00() []float64 {

	sx := Mul(psi_k00, Dot(&M, delta_mx) )
	sy := Mul(psi_k00, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek01() []float64 {

	sx := Mul(psi_k01, Dot(&M, delta_mx) )
	sy := Mul(psi_k01, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek02() []float64 {

	sx := Mul(psi_k02, Dot(&M, delta_mx) )
	sy := Mul(psi_k02, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek03() []float64 {

	sx := Mul(psi_k03, Dot(&M, delta_mx) )
	sy := Mul(psi_k03, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek04() []float64 {

	sx := Mul(psi_k04, Dot(&M, delta_mx) )
	sy := Mul(psi_k04, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek05() []float64 {

	sx := Mul(psi_k05, Dot(&M, delta_mx) )
	sy := Mul(psi_k05, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek06() []float64 {

	sx := Mul(psi_k06, Dot(&M, delta_mx) )
	sy := Mul(psi_k06, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek07() []float64 {

	sx := Mul(psi_k07, Dot(&M, delta_mx) )
	sy := Mul(psi_k07, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek08() []float64 {

	sx := Mul(psi_k08, Dot(&M, delta_mx) )
	sy := Mul(psi_k08, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek09() []float64 {

	sx := Mul(psi_k09, Dot(&M, delta_mx) )
	sy := Mul(psi_k09, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek10() []float64 {

	sx := Mul(psi_k10, Dot(&M, delta_mx) )
	sy := Mul(psi_k10, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek11() []float64 {

	sx := Mul(psi_k11, Dot(&M, delta_mx) )
	sy := Mul(psi_k11, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek12() []float64 {

	sx := Mul(psi_k12, Dot(&M, delta_mx) )
	sy := Mul(psi_k12, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek13() []float64 {

	sx := Mul(psi_k13, Dot(&M, delta_mx) )
	sy := Mul(psi_k13, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek14() []float64 {

	sx := Mul(psi_k14, Dot(&M, delta_mx) )
	sy := Mul(psi_k14, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek15() []float64 {

	sx := Mul(psi_k15, Dot(&M, delta_mx) )
	sy := Mul(psi_k15, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek16() []float64 {

	sx := Mul(psi_k16, Dot(&M, delta_mx) )
	sy := Mul(psi_k16, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek17() []float64 {

	sx := Mul(psi_k17, Dot(&M, delta_mx) )
	sy := Mul(psi_k17, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek18() []float64 {

	sx := Mul(psi_k18, Dot(&M, delta_mx) )
	sy := Mul(psi_k18, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek19() []float64 {

	sx := Mul(psi_k19, Dot(&M, delta_mx) )
	sy := Mul(psi_k19, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek20() []float64 {

	sx := Mul(psi_k20, Dot(&M, delta_mx) )
	sy := Mul(psi_k20, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek21() []float64 {

	sx := Mul(psi_k21, Dot(&M, delta_mx) )
	sy := Mul(psi_k21, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek22() []float64 {

	sx := Mul(psi_k22, Dot(&M, delta_mx) )
	sy := Mul(psi_k22, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek23() []float64 {

	sx := Mul(psi_k23, Dot(&M, delta_mx) )
	sy := Mul(psi_k23, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek24() []float64 {

	sx := Mul(psi_k24, Dot(&M, delta_mx) )
	sy := Mul(psi_k24, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek25() []float64 {

	sx := Mul(psi_k25, Dot(&M, delta_mx) )
	sy := Mul(psi_k25, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek26() []float64 {

	sx := Mul(psi_k26, Dot(&M, delta_mx) )
	sy := Mul(psi_k26, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek27() []float64 {

	sx := Mul(psi_k27, Dot(&M, delta_mx) )
	sy := Mul(psi_k27, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek28() []float64 {

	sx := Mul(psi_k28, Dot(&M, delta_mx) )
	sy := Mul(psi_k28, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek29() []float64 {

	sx := Mul(psi_k29, Dot(&M, delta_mx) )
	sy := Mul(psi_k29, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek30() []float64 {

	sx := Mul(psi_k30, Dot(&M, delta_mx) )
	sy := Mul(psi_k30, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek31() []float64 {

	sx := Mul(psi_k31, Dot(&M, delta_mx) )
	sy := Mul(psi_k31, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek32() []float64 {

	sx := Mul(psi_k32, Dot(&M, delta_mx) )
	sy := Mul(psi_k32, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek33() []float64 {

	sx := Mul(psi_k33, Dot(&M, delta_mx) )
	sy := Mul(psi_k33, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek34() []float64 {

	sx := Mul(psi_k34, Dot(&M, delta_mx) )
	sy := Mul(psi_k34, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek35() []float64 {

	sx := Mul(psi_k35, Dot(&M, delta_mx) )
	sy := Mul(psi_k35, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek36() []float64 {

	sx := Mul(psi_k36, Dot(&M, delta_mx) )
	sy := Mul(psi_k36, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek37() []float64 {

	sx := Mul(psi_k37, Dot(&M, delta_mx) )
	sy := Mul(psi_k37, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek38() []float64 {

	sx := Mul(psi_k38, Dot(&M, delta_mx) )
	sy := Mul(psi_k38, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek39() []float64 {

	sx := Mul(psi_k39, Dot(&M, delta_mx) )
	sy := Mul(psi_k39, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek40() []float64 {

	sx := Mul(psi_k40, Dot(&M, delta_mx) )
	sy := Mul(psi_k40, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek41() []float64 {

	sx := Mul(psi_k41, Dot(&M, delta_mx) )
	sy := Mul(psi_k41, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek42() []float64 {

	sx := Mul(psi_k42, Dot(&M, delta_mx) )
	sy := Mul(psi_k42, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek43() []float64 {

	sx := Mul(psi_k43, Dot(&M, delta_mx) )
	sy := Mul(psi_k43, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek44() []float64 {

	sx := Mul(psi_k44, Dot(&M, delta_mx) )
	sy := Mul(psi_k44, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek45() []float64 {

	sx := Mul(psi_k45, Dot(&M, delta_mx) )
	sy := Mul(psi_k45, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek46() []float64 {

	sx := Mul(psi_k46, Dot(&M, delta_mx) )
	sy := Mul(psi_k46, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek47() []float64 {

	sx := Mul(psi_k47, Dot(&M, delta_mx) )
	sy := Mul(psi_k47, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek48() []float64 {

	sx := Mul(psi_k48, Dot(&M, delta_mx) )
	sy := Mul(psi_k48, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek49() []float64 {

	sx := Mul(psi_k49, Dot(&M, delta_mx) )
	sy := Mul(psi_k49, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek50() []float64 {

	sx := Mul(psi_k50, Dot(&M, delta_mx) )
	sy := Mul(psi_k50, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}