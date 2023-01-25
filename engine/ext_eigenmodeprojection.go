package engine

// **************************************** 
// Author(s): Joo-Von Kim, C2N, CNRS/Univ. Paris-Saclay
//
// This module projects the magnetization onto user-supplied transverse directions,
// delta_mx, delta_my (obtained, e.g., from the relaxed magnetization state), with
// a spatial convolution with the user-supplied mask psi_k. It returns the amplitudes
// 
// 	a_kx = int_dV {psi_k (m . delta_mx)}
//	a_ky = int_dV {psi_k (m . delta_my)}
//
// The user-supplied masks/vector fields can be added in the source .mx3 file with
//	psi_k.Add( LoadFile(("psi_file.ovf"),1) )	
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
	psi_k		= NewScalarExcitation("psi_k", "", "Eigenmode spatial profile")
	delta_mx	= NewExcitation("delta_mx", "", "Transverse magnetization 1")
	delta_my	= NewExcitation("delta_my", "", "Transverse magnetization 2")
	a_k		= NewVectorValue("a_k", "", "delta_mx,y projection onto psi_k", GetModeAmplitude)
)
	
	
func GetModeAmplitude() []float64 {

	sx := Mul(psi_k, Dot(&M, delta_mx) )
	sy := Mul(psi_k, Dot(&M, delta_my) )

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
