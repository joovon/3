#include <stdint.h>
#include <stdio.h>
#include "exchange.h"
#include "float3.h"
#include "stencil.h"

// Simple routine to count the number of spins with mz <= 0
// Used to obtain, e.g., a rough estimate of a skyrmion core size under
// finite temperatures
// See reversedspins.go.
extern "C" __global__ void
setreversedspins(float* __restrict__ s,
                     float* __restrict__ mx, float* __restrict__ my, float* __restrict__ mz,
                     int Nx, int Ny, int Nz, uint8_t PBC) {

    int ix = blockIdx.x * blockDim.x + threadIdx.x;
    int iy = blockIdx.y * blockDim.y + threadIdx.y;
    int iz = blockIdx.z * blockDim.z + threadIdx.z;

    if (ix >= Nx || iy >= Ny || iz >= Nz)
    {
        return;
    }

    int I = idx(ix, iy, iz);                      // central cell index

    float3 m0 = make_float3(mx[I], my[I], mz[I]); // +0
    float mz0 = mz[I];
    int i_;                                       // neighbor index

    if(is0(m0))
    {
        s[I] = 0.0f;
        return;
    }

    if(mz0>0)
    {
        s[I] = 0.0f;
    }
    else
    {
        s[I] = -1.0f;
    }

}
