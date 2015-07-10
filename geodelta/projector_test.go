package geodelta

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func AssertEqual(t *testing.T, expected float64, actual float64) {
	if expected != actual {
		t.Errorf("expected %f but %f", expected, actual)
	}
}

func AssertInDelta(t *testing.T, expected float64, actual float64, delta float64) {
	if math.Abs(expected-actual) > delta {
		t.Errorf("expected %f but %f", expected, actual)
	}
}

func TestLatToMy(t *testing.T) {
	assert := assert.New(t)
	assert.InDelta(+1.0, LatToMy(+85.0511), 1.0e-5)
	assert.Equal(0.0, LatToMy(0.0))
	assert.InDelta(-1.0, LatToMy(-85.0511), 1.0e-5)
}

func TestLngToMx(t *testing.T) {
	AssertEqual(t, +1.0, LngToMx(+180.0))
	AssertEqual(t, +0.5, LngToMx(+90.0))
	AssertEqual(t, 0.0, LngToMx(0.0))
	AssertEqual(t, -0.5, LngToMx(-90.0))
	AssertEqual(t, -1.0, LngToMx(-180.0))
}

func TestNormalizeM(t *testing.T) {
	table := [...][3]float64{
		{-0.5, NormalizeM(+3.5), 1e-15},
		{+0.5, NormalizeM(+2.5), 1e-15},
		{-0.5, NormalizeM(+1.5), 1e-15},
		{+1.0, NormalizeM(+1.0), 1e-15},
		{+0.5, NormalizeM(+0.5), 1e-15},
		{0.0, NormalizeM(0.0), 1e-15},
		{-0.5, NormalizeM(-0.5), 1e-15},
		{-1.0, NormalizeM(-1.0), 1e-15},
		{+0.5, NormalizeM(-1.5), 1e-15},
		{-0.5, NormalizeM(-2.5), 1e-15},
		{+0.5, NormalizeM(-3.5), 1e-15},
	}
	for i := 0; i < len(table); i++ {
		expected, actual, tolerance := table[i][0], table[i][1], table[i][2]
		if math.Abs(expected-actual) > tolerance {
			t.Errorf("i=%d, expected=%e, actual=%e, delta=%e", i, expected, actual, (expected - actual))
		}
	}
}

func TestMyToLat(t *testing.T) {
	AssertInDelta(t, +85.0511, MyToLat(+1.0), 1.0e-4)
	AssertEqual(t, 0.0, MyToLat(0.0))
	AssertInDelta(t, -85.0511, MyToLat(-1.0), 1.0e-4)
}

func TestMxToLng(t *testing.T) {
	AssertEqual(t, -90.0, MxToLng(+1.5))
	AssertEqual(t, +180.0, MxToLng(+1.0))
	AssertEqual(t, +90.0, MxToLng(+0.5))
	AssertEqual(t, 0.0, MxToLng(0.0))
	AssertEqual(t, -90.0, MxToLng(-0.5))
	AssertEqual(t, -180.0, MxToLng(-1.0))
	AssertEqual(t, +90.0, MxToLng(-1.5))
}

/* Ruby
class GeoDeltaProjectorTest < Minitest::Test
  def test_my_to_ny
    max = @mod::DELTA_HEIGHT
    assert_equal(+12.0, @mod.my_to_ny(+max))
    assert_equal(  0.0, @mod.my_to_ny( 0.0))
    assert_equal(-12.0, @mod.my_to_ny(-max))
  end

  def test_mx_to_nx
    assert_equal(+12.0, @mod.mx_to_nx(+1.0))
    assert_equal(  0.0, @mod.mx_to_nx( 0.0))
    assert_equal(-12.0, @mod.mx_to_nx(-1.0))
  end

  def test_ny_to_my
    max = @mod::DELTA_HEIGHT
    assert_equal(+max, @mod.ny_to_my(+12.0))
    assert_equal( 0.0, @mod.ny_to_my(  0.0))
    assert_equal(-max, @mod.ny_to_my(-12.0))
  end

  def test_nx_to_mx
    assert_equal(+1.0, @mod.nx_to_mx(+12.0))
    assert_equal( 0.0, @mod.nx_to_mx(  0.0))
    assert_equal(-1.0, @mod.nx_to_mx(-12.0))
  end

  def test_lat_to_ny
    assert_equal(0.0, @mod.lat_to_ny(0.0))
    assert_equal(
      @mod.my_to_ny(@mod.lat_to_my(+82.4674)),
      @mod.lat_to_ny(+82.4674))
    assert_equal(
      @mod.my_to_ny(@mod.lat_to_my(-82.4674)),
      @mod.lat_to_ny(-82.4674))
  end

  def test_lng_to_nx
    assert_equal(0.0, @mod.lng_to_nx(0.0))
    assert_equal(
      @mod.mx_to_nx(@mod.lng_to_mx(+180.0)),
      @mod.lng_to_nx(+180.0))
    assert_equal(
      @mod.mx_to_nx(@mod.lng_to_mx(-180.0)),
      @mod.lng_to_nx(-180.0))
  end

  def test_ny_to_lat
    assert_equal(0.0, @mod.ny_to_lat(0.0))
    assert_equal(
      @mod.my_to_lat(@mod.ny_to_my(+12.0)),
      @mod.ny_to_lat(+12.0))
    assert_equal(
      @mod.my_to_lat(@mod.ny_to_my(-12.0)),
      @mod.ny_to_lat(-12.0))
  end

  def test_nx_to_lng
    assert_equal(0.0, @mod.nx_to_lng(0.0))
    assert_equal(
      @mod.mx_to_lng(@mod.nx_to_mx(+12.0)),
      @mod.nx_to_lng(+12.0))
    assert_equal(
      @mod.mx_to_lng(@mod.nx_to_mx(-12.0)),
      @mod.nx_to_lng(-12.0))
  end

  def test_latlng_to_nxy
    assert_equal([0.0, 0.0], @mod.latlng_to_nxy(0.0, 0.0))
    assert_equal(
      [@mod.lng_to_nx(+180.0), @mod.lat_to_ny(+82.4674)],
      @mod.latlng_to_nxy(+82.4674, +180.0))
    assert_equal(
      [@mod.lng_to_nx(-180.0), @mod.lat_to_ny(-82.4674)],
      @mod.latlng_to_nxy(-82.4674, -180.0))
  end

  def test_nxy_to_latlng
    assert_equal([0.0, 0.0], @mod.nxy_to_latlng(0.0, 0.0))
    assert_equal(
      [@mod.ny_to_lat(+12.0), @mod.nx_to_lng(+12.0)],
      @mod.nxy_to_latlng(+12.0, +12.0))
    assert_equal(
      [@mod.ny_to_lat(-12.0), @mod.nx_to_lng(-12.0)],
      @mod.nxy_to_latlng(-12.0, -12.0))
  end

  def test_rush__latlng_nxy
    1000.times {
      lat1, lng1 = [(rand * 180.0) - 90.0, (rand * 360.0) - 180.0]
      nx, ny     = @mod.latlng_to_nxy(lat1, lng1)
      lat2, lng2 = @mod.nxy_to_latlng(nx, ny)
      assert_in_delta(lat1, lat2, 1.0E-10)
      assert_in_delta(lng1, lng2, 1.0E-13)
    }
  end
end
*/
