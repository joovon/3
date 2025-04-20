package engine

var CorePosTop = NewVectorValue("ext_coreposTop", "m", "Vortex core position (x,y) + polarization (z)", corePosTop)
var CorePosBot = NewVectorValue("ext_coreposBot", "m", "Vortex core position (x,y) + polarization (z)", corePosBot)


func corePosTop() []float64 {
	m := M.Buffer()
	m_z := m.Comp(Z).HostCopy().Scalars()
	s := m.Size()
	Nx, Ny, Nz := s[X], s[Y], s[Z]

	max := float32(-1.0)
	var maxX, maxY int

	// Search for core position only in the top layer, z = Nz-1

	// Avoid the boundaries so the neighbor interpolation can't go out of bounds.
	for y := 1; y < Ny-1; y++ {
		for x := 1; x < Nx-1; x++ {
			m := abs(m_z[Nz-1][y][x])
			if m > max {
				maxX, maxY = x, y
				max = m
			}
		}
	}

	pos := make([]float64, 3)
	mz := m_z[Nz-1]

	// sub-cell interpolation in X and Y, but not Z
	pos[X] = float64(maxX) + interpolate_maxpos(
		max, -1, abs(mz[maxY][maxX-1]), 1, abs(mz[maxY][maxX+1])) -
		float64(Nx)/2 + 0.5
	pos[Y] = float64(maxY) + interpolate_maxpos(
		max, -1, abs(mz[maxY-1][maxX]), 1, abs(mz[maxY+1][maxX])) -
		float64(Ny)/2 + 0.5

	c := Mesh().CellSize()
	pos[X] *= c[X]
	pos[Y] *= c[Y]
	pos[Z] = float64(m_z[Nz-1][maxY][maxX]) // 3rd coordinate is core polarization

	pos[X] += GetShiftPos() // add simulation window shift
	return pos
}

func corePosBot() []float64 {
	m := M.Buffer()
	m_z := m.Comp(Z).HostCopy().Scalars()
	s := m.Size()
	Nx, Ny := s[X], s[Y]

	max := float32(-1.0)
	var maxX, maxY int

	// Search for core position only in the bottom layer, z = 0

	// Avoid the boundaries so the neighbor interpolation can't go out of bounds.
	for y := 1; y < Ny-1; y++ {
		for x := 1; x < Nx-1; x++ {
			m := abs(m_z[0][y][x])
			if m > max {
				maxX, maxY = x, y
				max = m
			}
		}
	}

	pos := make([]float64, 3)
	mz := m_z[0]

	// sub-cell interpolation in X and Y, but not Z
	pos[X] = float64(maxX) + interpolate_maxpos(
		max, -1, abs(mz[maxY][maxX-1]), 1, abs(mz[maxY][maxX+1])) -
		float64(Nx)/2 + 0.5
	pos[Y] = float64(maxY) + interpolate_maxpos(
		max, -1, abs(mz[maxY-1][maxX]), 1, abs(mz[maxY+1][maxX])) -
		float64(Ny)/2 + 0.5

	c := Mesh().CellSize()
	pos[X] *= c[X]
	pos[Y] *= c[Y]
	pos[Z] = float64(m_z[0][maxY][maxX]) // 3rd coordinate is core polarization

	pos[X] += GetShiftPos() // add simulation window shift
	return pos
}

// func interpolate_maxpos(f0, d1, f1, d2, f2 float32) float64 {
// 	b := (f2 - f1) / (d2 - d1)
// 	a := ((f2-f0)/d2 - (f0-f1)/(-d1)) / (d2 - d1)
// 	return float64(-b / (2 * a))
// }

// func abs(x float32) float32 {
// 	if x > 0 {
// 		return x
// 	} else {
// 		return -x
// 	}
// }