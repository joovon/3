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
// with up to 11 modes (n = 0 to 10).
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
//	delta_mx	= NewExcitation("delta_mx", "", "Transverse magnetization 1")
//	delta_my	= NewExcitation("delta_my", "", "Transverse magnetization 2")

	psi_k0		= NewScalarExcitation("psi_k0", "", "Eigenmode spatial profile")
	psi_k1		= NewScalarExcitation("psi_k1", "", "Eigenmode spatial profile")
	psi_k2		= NewScalarExcitation("psi_k2", "", "Eigenmode spatial profile")
	psi_k3		= NewScalarExcitation("psi_k3", "", "Eigenmode spatial profile")
	psi_k4		= NewScalarExcitation("psi_k4", "", "Eigenmode spatial profile")
	psi_k5		= NewScalarExcitation("psi_k5", "", "Eigenmode spatial profile")
	psi_k6		= NewScalarExcitation("psi_k6", "", "Eigenmode spatial profile")
	psi_k7		= NewScalarExcitation("psi_k7", "", "Eigenmode spatial profile")
	psi_k8		= NewScalarExcitation("psi_k8", "", "Eigenmode spatial profile")
	psi_k9		= NewScalarExcitation("psi_k9", "", "Eigenmode spatial profile")
	psi_k10		= NewScalarExcitation("psi_k10", "", "Eigenmode spatial profile")
		
	a_k0		= NewVectorValue("a_k0", "", "delta_mx projection onto psi_k0", GetModeAmplitudek0 )
	a_k1		= NewVectorValue("a_k1", "", "delta_my projection onto psi_k1", GetModeAmplitudek1 )
	a_k2		= NewVectorValue("a_k2", "", "delta_my projection onto psi_k2", GetModeAmplitudek2 )
	a_k3		= NewVectorValue("a_k3", "", "delta_my projection onto psi_k3", GetModeAmplitudek3 )
	a_k4		= NewVectorValue("a_k4", "", "delta_my projection onto psi_k4", GetModeAmplitudek4 )
	a_k5		= NewVectorValue("a_k5", "", "delta_my projection onto psi_k5", GetModeAmplitudek5 )
	a_k6		= NewVectorValue("a_k6", "", "delta_my projection onto psi_k6", GetModeAmplitudek6 )
	a_k7		= NewVectorValue("a_k7", "", "delta_my projection onto psi_k7", GetModeAmplitudek7 )
	a_k8		= NewVectorValue("a_k8", "", "delta_my projection onto psi_k8", GetModeAmplitudek8 )
	a_k9		= NewVectorValue("a_k9", "", "delta_my projection onto psi_k9", GetModeAmplitudek9 )
	a_k10		= NewVectorValue("a_k10", "", "delta_my projection onto psi_k10", GetModeAmplitudek10 )
	
)
	
	
func GetModeAmplitudek0() []float64 {

	sx := Mul(psi_k0, Dot(&M, delta_mx) )
	sy := Mul(psi_k0, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek1() []float64 {

	sx := Mul(psi_k1, Dot(&M, delta_mx) )
	sy := Mul(psi_k1, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek2() []float64 {

	sx := Mul(psi_k2, Dot(&M, delta_mx) )
	sy := Mul(psi_k2, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek3() []float64 {

	sx := Mul(psi_k3, Dot(&M, delta_mx) )
	sy := Mul(psi_k3, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek4() []float64 {

	sx := Mul(psi_k4, Dot(&M, delta_mx) )
	sy := Mul(psi_k4, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek5() []float64 {

	sx := Mul(psi_k5, Dot(&M, delta_mx) )
	sy := Mul(psi_k5, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek6() []float64 {

	sx := Mul(psi_k6, Dot(&M, delta_mx) )
	sy := Mul(psi_k6, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek7() []float64 {

	sx := Mul(psi_k7, Dot(&M, delta_mx) )
	sy := Mul(psi_k7, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek8() []float64 {

	sx := Mul(psi_k8, Dot(&M, delta_mx) )
	sy := Mul(psi_k8, Dot(&M, delta_my) )
		
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

func GetModeAmplitudek9() []float64 {

	sx := Mul(psi_k9, Dot(&M, delta_mx) )
	sy := Mul(psi_k9, Dot(&M, delta_my) )
		
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
