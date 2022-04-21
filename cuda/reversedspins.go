package cuda

import (
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

// Computes the number of spins with mz <= 0
// See reversedspins.cu
func SetReversedSpins(s *data.Slice, m *data.Slice, mesh *data.Mesh) {
//	cellsize := mesh.CellSize()
	N := s.Size()
	util.Argument(m.Size() == N)
	cfg := make3DConf(N)

	k_setreversedspins_async(s.DevPtr(X),
		m.DevPtr(X), m.DevPtr(Y), m.DevPtr(Z),
		N[X], N[Y], N[Z], mesh.PBC_code(), cfg)
}
