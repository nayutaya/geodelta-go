package geodelta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatToMy(t *testing.T) {
	assert := assert.New(t)
	assert.InDelta(+1.0, float64(LatToMy(+85.0511)), 1.0e-5)
	assert.InDelta(0.0, float64(LatToMy(0.0)), 1.0e-5)
	assert.InDelta(-1.0, float64(LatToMy(-85.0511)), 1.0e-5)
}

func TestLngToMx(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Mx(+1.0), LngToMx(+180.0))
	assert.Equal(Mx(+0.5), LngToMx(+90.0))
	assert.Equal(Mx(0.0), LngToMx(0.0))
	assert.Equal(Mx(-0.5), LngToMx(-90.0))
	assert.Equal(Mx(-1.0), LngToMx(-180.0))
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
	assert.InDelta(+85.0511, float64(MyToLat(+1.0)), 1.0e-4)
	assert.InDelta(0.0, float64(MyToLat(0.0)), 1.0e-4)
	assert.InDelta(-85.0511, float64(MyToLat(-1.0)), 1.0e-4)
}

func TestMxToLng(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Lng(-90.0), MxToLng(+1.5))
	assert.Equal(Lng(+180.0), MxToLng(+1.0))
	assert.Equal(Lng(+90.0), MxToLng(+0.5))
	assert.Equal(Lng(0.0), MxToLng(0.0))
	assert.Equal(Lng(-90.0), MxToLng(-0.5))
	assert.Equal(Lng(-180.0), MxToLng(-1.0))
	assert.Equal(Lng(+90.0), MxToLng(-1.5))
}

func TestMyToNy(t *testing.T) {
	max := DELTA_HEIGHT
	assert := assert.New(t)
	assert.Equal(Ny(+12.0), MyToNy(My(+max)))
	assert.Equal(Ny(0.0), MyToNy(0.0))
	assert.Equal(Ny(-12.0), MyToNy(My(-max)))
}

func TestMxToNx(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Nx(+12.0), MxToNx(+1.0))
	assert.Equal(Nx(0.0), MxToNx(0.0))
	assert.Equal(Nx(-12.0), MxToNx(-1.0))
}

func TestNyToMy(t *testing.T) {
	max := DELTA_HEIGHT
	assert := assert.New(t)
	assert.Equal(My(+max), NyToMy(+12.0))
	assert.Equal(My(0.0), NyToMy(0.0))
	assert.Equal(My(-max), NyToMy(-12.0))
}

func TestNxToMx(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Mx(+1.0), NxToMx(+12.0))
	assert.Equal(Mx(0.0), NxToMx(0.0))
	assert.Equal(Mx(-1.0), NxToMx(-12.0))
}

func TestLatToNy(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Ny(0.0), LatToNy(0.0))
	assert.Equal(Ny(MyToNy(LatToMy(+82.4674))), LatToNy(+82.4674))
	assert.Equal(Ny(MyToNy(LatToMy(-82.4674))), LatToNy(-82.4674))
}

func TestLngToNx(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Nx(0.0), LngToNx(0.0))
	assert.Equal(Nx(MxToNx(LngToMx(+180.0))), LngToNx(+180.0))
	assert.Equal(Nx(MxToNx(LngToMx(-180.0))), LngToNx(-180.0))
}

func TestNyToLat(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Lat(0.0), NyToLat(0.0))
	assert.Equal(Lat(MyToLat(NyToMy(+12.0))), NyToLat(+12.0))
	assert.Equal(Lat(MyToLat(NyToMy(-12.0))), NyToLat(-12.0))
}

func TestNxToLng(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Lng(0.0), NxToLng(Nx(0.0)))
	assert.Equal(Lng(MxToLng(NxToMx(+12.0))), NxToLng(Nx(+12.0)))
	assert.Equal(Lng(MxToLng(NxToMx(-12.0))), NxToLng(Nx(-12.0)))
}

/* Ruby
func TestLatlngToNxy(t *testing.T) {
	assert := assert.New(t)
  assert.Equal([0.0, 0.0], @mod.latlng_to_nxy(0.0, 0.0))
  assert.Equal(
    [LngToNx(+180.0), LatToNy(+82.4674)],
    @mod.latlng_to_nxy(+82.4674, +180.0))
  assert.Equal(
    [LngToNx(-180.0), LatToNy(-82.4674)],
    @mod.latlng_to_nxy(-82.4674, -180.0))
}

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
