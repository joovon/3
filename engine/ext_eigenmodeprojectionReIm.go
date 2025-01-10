package engine

// **************************************** 
// Author(s): Joo-Von Kim, C2N, CNRS/Univ. Paris-Saclay
//
// This module projects the magnetization onto user-supplied eigenvectors
// Re(psi_k) and Im(psi_k), which are both 3-component fields. The module
// returns the amplitudes
// 
// 	Re(b_k) = int_dV {(Im(psi_k) x m0).(m-m0)}
//	Im(b_k) = int_dV {(Re(psi_k) x m0).(m-m0)}
//  n_k     = b_k^* b_k = Re(b_k)^2 + Im(b_k)^2
//
// The user-supplied vector fields can be added in the source .mx3 file with
//  M0.Add( LoadFile(("m0_file.ovf"),1) )	
//	psiRe_k.Add( LoadFile(("psi_file.ovf"),1) )	
//	psiIm_k.Add( LoadFile(("psi_file.ovf"),1) )	
//	etc.
//
// Acknowledgements:
// This work was supported by Horizon 2020 Research Framework Programme of the
// European Commission under grant agreement No. 899646 (k-Net). 
//
// **************************************** 


import (
	"github.com/mumax/3/cuda"
)


var (
	M0			= NewExcitation("M0", "", "Equilibrium magnetization configuration")
	psiRe_k		= NewExcitation("psiRe_k", "", "Real part of eigenmode vector")
	psiIm_k		= NewExcitation("psiIm_k", "", "Imaginary part of eigenmode vector")
	b_k			= NewVectorValue("b_k", "", "m projection onto psi(Re,Im)_k", GetModeAmplitudeReIm)
)
	
	
func GetModeAmplitudeReIm() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)
	
	vpRe	:= Dot(Cross(&M, psiIm_k), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}