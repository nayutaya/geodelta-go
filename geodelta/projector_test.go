package geodelta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatToMy(t *testing.T) {
	assert := assert.New(t)
	assert.InDelta(+1.0, float64(Lat(+85.0511).ToMy()), 1.0e-5)
	assert.InDelta(0.0, float64(Lat(0.0).ToMy()), 1.0e-5)
	assert.InDelta(-1.0, float64(Lat(-85.0511).ToMy()), 1.0e-5)
}

func TestLngToMx(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Mx(+1.0), Lng(+180.0).ToMx())
	assert.Equal(Mx(+0.5), Lng(+90.0).ToMx())
	assert.Equal(Mx(0.0), Lng(0.0).ToMx())
	assert.Equal(Mx(-0.5), Lng(-90.0).ToMx())
	assert.Equal(Mx(-1.0), Lng(-180.0).ToMx())
}

func TestNormalizeM(t *testing.T) {
	assert := assert.New(t)
	assert.InDelta(-0.5, NormalizeM(+3.5), 1e-15)
	assert.InDelta(+0.5, NormalizeM(+2.5), 1e-15)
	assert.InDelta(-0.5, NormalizeM(+1.5), 1e-15)
	assert.InDelta(+1.0, NormalizeM(+1.0), 1e-15)
	assert.InDelta(+0.5, NormalizeM(+0.5), 1e-15)
	assert.InDelta(0.0, NormalizeM(0.0), 1e-15)
	assert.InDelta(-0.5, NormalizeM(-0.5), 1e-15)
	assert.InDelta(-1.0, NormalizeM(-1.0), 1e-15)
	assert.InDelta(+0.5, NormalizeM(-1.5), 1e-15)
	assert.InDelta(-0.5, NormalizeM(-2.5), 1e-15)
	assert.InDelta(+0.5, NormalizeM(-3.5), 1e-15)
}

func TestMyToLat(t *testing.T) {
	assert := assert.New(t)
	assert.InDelta(+85.0511, float64(My(+1.0).ToLat()), 1.0e-4)
	assert.InDelta(0.0, float64(My(0.0).ToLat()), 1.0e-4)
	assert.InDelta(-85.0511, float64(My(-1.0).ToLat()), 1.0e-4)
}

func TestMxToLng(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Lng(-90.0), Mx(+1.5).ToLng())
	assert.Equal(Lng(+180.0), Mx(+1.0).ToLng())
	assert.Equal(Lng(+90.0), Mx(+0.5).ToLng())
	assert.Equal(Lng(0.0), Mx(0.0).ToLng())
	assert.Equal(Lng(-90.0), Mx(-0.5).ToLng())
	assert.Equal(Lng(-180.0), Mx(-1.0).ToLng())
	assert.Equal(Lng(+90.0), Mx(-1.5).ToLng())
}

func TestMyToNy(t *testing.T) {
	max := DELTA_HEIGHT
	assert := assert.New(t)
	assert.Equal(Ny(+12.0), My(+max).ToNy())
	assert.Equal(Ny(0.0), My(0.0).ToNy())
	assert.Equal(Ny(-12.0), My(-max).ToNy())
}

func TestMxToNx(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Nx(+12.0), Mx(+1.0).ToNx())
	assert.Equal(Nx(0.0), Mx(0.0).ToNx())
	assert.Equal(Nx(-12.0), Mx(-1.0).ToNx())
}

func TestNyToMy(t *testing.T) {
	max := DELTA_HEIGHT
	assert := assert.New(t)
	assert.Equal(My(+max), Ny(+12.0).ToMy())
	assert.Equal(My(0.0), Ny(0.0).ToMy())
	assert.Equal(My(-max), Ny(-12.0).ToMy())
}

func TestNxToMx(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Mx(+1.0), Nx(+12.0).ToMx())
	assert.Equal(Mx(0.0), Nx(0.0).ToMx())
	assert.Equal(Mx(-1.0), Nx(-12.0).ToMx())
}

func TestLatToNy(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Ny(0.0), Lat(0.0).ToNy())
	assert.Equal(Lat(+82.4674).ToMy().ToNy(), Lat(+82.4674).ToNy())
	assert.Equal(Lat(-82.4674).ToMy().ToNy(), Lat(-82.4674).ToNy())
}

func TestLngToNx(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Nx(0.0), Lng(0.0).ToNx())
	assert.Equal(Lng(+180.0).ToMx().ToNx(), Lng(+180.0).ToNx())
	assert.Equal(Lng(-180.0).ToMx().ToNx(), Lng(-180.0).ToNx())
}

func TestNyToLat(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Lat(0.0), Ny(0.0).ToLat())
	assert.Equal(Ny(+12.0).ToMy().ToLat(), Ny(+12.0).ToLat())
	assert.Equal(Ny(-12.0).ToMy().ToLat(), Ny(-12.0).ToLat())
}

func TestNxToLng(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Lng(0.0), Nx(0.0).ToLng())
	assert.Equal(Nx(+12.0).ToMx().ToLng(), Nx(+12.0).ToLng())
	assert.Equal(Nx(-12.0).ToMx().ToLng(), Nx(-12.0).ToLng())
}

func TestLatLngToNxNy(t *testing.T) {
	assert := assert.New(t)
	{
		nx, ny := LatLngToNxNy(Lat(0.0), Lng(0.0))
		assert.Equal(Nx(0.0), nx)
		assert.Equal(Ny(0.0), ny)
	}
	{
		nx, ny := LatLngToNxNy(Lat(+82.4674), Lng(+180.0))
		assert.Equal(Lng(+180.0).ToNx(), nx)
		assert.Equal(Lat(+82.4674).ToNy(), ny)
	}
	{
		nx, ny := LatLngToNxNy(Lat(-82.4674), Lng(-180.0))
		assert.Equal(Lng(-180.0).ToNx(), nx)
		assert.Equal(Lat(-82.4674).ToNy(), ny)
	}
}

/* Ruby
func TestNxyToLatlng(t *testing.T) {
	assert := assert.New(t)
  assert.Equal([0.0, 0.0], @mod.nxy_to_latlng(0.0, 0.0))
  assert.Equal(
    [NyToLat(+12.0), NxToLng(+12.0)],
    @mod.nxy_to_latlng(+12.0, +12.0))
  assert.Equal(
    [NyToLat(-12.0), NxToLng(-12.0)],
    @mod.nxy_to_latlng(-12.0, -12.0))
}

func TestRushLatlngNxy(t *testing.T) {
	assert := assert.New(t)
  1000.times {
    lat1, lng1 = [(rand * 180.0) - 90.0, (rand * 360.0) - 180.0]
    nx, ny     = @mod.latlng_to_nxy(lat1, lng1)
    lat2, lng2 = @mod.nxy_to_latlng(nx, ny)
    assert_in_delta(lat1, lat2, 1.0E-10)
    assert_in_delta(lng1, lng2, 1.0E-13)
  }
}
*/
