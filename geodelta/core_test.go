package geodelta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDeltaIds(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0}, GetDeltaIds(+45.0, +0.0, 1))
	assert.Equal([]byte{1}, GetDeltaIds(+45.0, +90.0, 1))
	assert.Equal([]byte{2}, GetDeltaIds(+45.0, +180.0, 1))
	assert.Equal([]byte{3}, GetDeltaIds(+45.0, -90.0, 1))
	assert.Equal([]byte{2}, GetDeltaIds(+45.0, -180.0, 1))

	assert.Equal([]byte{4}, GetDeltaIds(-45.0, +0.0, 1))
	assert.Equal([]byte{5}, GetDeltaIds(-45.0, +90.0, 1))
	assert.Equal([]byte{6}, GetDeltaIds(-45.0, +180.0, 1))
	assert.Equal([]byte{7}, GetDeltaIds(-45.0, -90.0, 1))
	assert.Equal([]byte{6}, GetDeltaIds(-45.0, -180.0, 1))

	assert.Equal([]byte{0}, GetDeltaIds(+0.0, +0.0, 1))
	assert.Equal([]byte{0, 1}, GetDeltaIds(+0.0, +0.0, 2))
	assert.Equal([]byte{0, 1, 1}, GetDeltaIds(+0.0, +0.0, 3))
	assert.Equal([]byte{0, 1, 1, 1}, GetDeltaIds(+0.0, +0.0, 4))
}

func TestGetDeltaCode(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Z", GetDeltaCode(+45.0, +0.0, 1))
	assert.Equal("Y", GetDeltaCode(+45.0, +90.0, 1))
	assert.Equal("X", GetDeltaCode(+45.0, +180.0, 1))
	assert.Equal("W", GetDeltaCode(+45.0, -90.0, 1))
	assert.Equal("X", GetDeltaCode(+45.0, -180.0, 1))

	assert.Equal("V", GetDeltaCode(-45.0, +0.0, 1))
	assert.Equal("T", GetDeltaCode(-45.0, +90.0, 1))
	assert.Equal("S", GetDeltaCode(-45.0, +180.0, 1))
	assert.Equal("R", GetDeltaCode(-45.0, -90.0, 1))
	assert.Equal("S", GetDeltaCode(-45.0, -180.0, 1))

	assert.Equal("Z", GetDeltaCode(+0.0, +0.0, 1))
	assert.Equal("ZM", GetDeltaCode(+0.0, +0.0, 2))
	assert.Equal("Z7", GetDeltaCode(+0.0, +0.0, 3))
	assert.Equal("Z7M", GetDeltaCode(+0.0, +0.0, 4))
}

/*
class GeoDeltaTest < Minitest::Test
  def test_get_center_from_delta_ids
    lat, lng = @mod.get_center_from_delta_ids([0])
    assert_in_delta( +71.480, lat, 1.0E-3)
    assert_in_delta(  +0.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_ids([1])
    assert_in_delta( +46.024, lat, 1.0E-3)
    assert_in_delta( +90.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_ids([2])
    assert_in_delta( +71.480, lat, 1.0E-3)
    assert_in_delta(-180.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_ids([3])
    assert_in_delta( +46.024, lat, 1.0E-3)
    assert_in_delta( -90.000, lng, 1.0E-3)

    lat, lng = @mod.get_center_from_delta_ids([4])
    assert_in_delta( -71.480, lat, 1.0E-3)
    assert_in_delta(  +0.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_ids([5])
    assert_in_delta( -46.024, lat, 1.0E-3)
    assert_in_delta( +90.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_ids([6])
    assert_in_delta( -71.480, lat, 1.0E-3)
    assert_in_delta(-180.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_ids([7])
    assert_in_delta( -46.024, lat, 1.0E-3)
    assert_in_delta( -90.000, lng, 1.0E-3)

    lat, lng = @mod.get_center_from_delta_ids([0, 0])
    assert_in_delta(+71.480, lat, 1.0E-3)
    assert_in_delta( +0.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_ids([0, 0, 0])
    assert_in_delta(+71.480, lat, 1.0E-3)
    assert_in_delta( +0.000, lng, 1.0E-3)
  end

  def test_get_center_from_delta_code
    lat, lng = @mod.get_center_from_delta_code("Z")
    assert_in_delta( +71.480, lat, 1.0E-3)
    assert_in_delta(  +0.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_code("Y")
    assert_in_delta( +46.024, lat, 1.0E-3)
    assert_in_delta( +90.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_code("X")
    assert_in_delta( +71.480, lat, 1.0E-3)
    assert_in_delta(-180.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_code("W")
    assert_in_delta( +46.024, lat, 1.0E-3)
    assert_in_delta( -90.000, lng, 1.0E-3)

    lat, lng = @mod.get_center_from_delta_code("V")
    assert_in_delta( -71.480, lat, 1.0E-3)
    assert_in_delta(  +0.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_code("T")
    assert_in_delta( -46.024, lat, 1.0E-3)
    assert_in_delta( +90.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_code("S")
    assert_in_delta( -71.480, lat, 1.0E-3)
    assert_in_delta(-180.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_code("R")
    assert_in_delta( -46.024, lat, 1.0E-3)
    assert_in_delta( -90.000, lng, 1.0E-3)

    lat, lng = @mod.get_center_from_delta_code("ZK")
    assert_in_delta(+71.480, lat, 1.0E-3)
    assert_in_delta( +0.000, lng, 1.0E-3)
    lat, lng = @mod.get_center_from_delta_code("Z2")
    assert_in_delta(+71.480, lat, 1.0E-3)
    assert_in_delta( +0.000, lng, 1.0E-3)
  end

  def test_get_coordinates_from_ids__1
    delta0 = @mod.get_coordinates_from_ids([0])
    assert.Equal(4, delta0.size)
    assert.Equal(2, delta0[0].size)
    assert.Equal(2, delta0[1].size)
    assert.Equal(2, delta0[2].size)
    assert.Equal(2, delta0[3].size)
    assert_in_delta( +71.480, delta0[0][0], 1.0E-3)
    assert_in_delta(  +0.000, delta0[0][1], 1.0E-3)
    assert_in_delta(  +0.000, delta0[1][0], 1.0E-3)
    assert_in_delta(  +0.000, delta0[1][1], 1.0E-3)
    assert_in_delta( +82.467, delta0[2][0], 1.0E-3)
    assert_in_delta( -90.000, delta0[2][1], 1.0E-3)
    assert_in_delta( +82.467, delta0[3][0], 1.0E-3)
    assert_in_delta( +90.000, delta0[3][1], 1.0E-3)

    delta4 = @mod.get_coordinates_from_ids([4])
    assert_in_delta( -71.480, delta4[0][0], 1.0E-3)
    assert_in_delta(  +0.000, delta4[0][1], 1.0E-3)
    assert_in_delta(  +0.000, delta4[1][0], 1.0E-3)
    assert_in_delta(  +0.000, delta4[1][1], 1.0E-3)
    assert_in_delta( -82.467, delta4[2][0], 1.0E-3)
    assert_in_delta( +90.000, delta4[2][1], 1.0E-3)
    assert_in_delta( -82.467, delta4[3][0], 1.0E-3)
    assert_in_delta( -90.000, delta4[3][1], 1.0E-3)
  end

  def test_get_coordinates_from_ids__2
    delta = (0..7).map { |id| @mod.get_coordinates_from_ids([id]) }
    assert.Equal(delta[0][1], delta[1][3])
    assert.Equal(delta[0][1], delta[3][2])
    assert.Equal(delta[0][1], delta[4][1])
    assert.Equal(delta[0][1], delta[5][2])
    assert.Equal(delta[0][1], delta[7][3])
    assert.Equal(delta[0][2], delta[2][3])
    assert.Equal(delta[0][2], delta[3][1])
    assert.Equal(delta[0][3], delta[1][1])
    assert.Equal(delta[0][3], delta[2][2])
    assert.Equal(delta[1][2], delta[2][1])
    assert.Equal(delta[1][2], delta[3][3])
    assert.Equal(delta[1][2], delta[5][3])
    assert.Equal(delta[1][2], delta[6][1])
    assert.Equal(delta[1][2], delta[7][2])
    assert.Equal(delta[4][3], delta[6][2])
    assert.Equal(delta[4][3], delta[7][1])
    assert.Equal(delta[4][2], delta[5][1])
    assert.Equal(delta[4][2], delta[6][3])
  end

  def test_get_coordinates_from_ids__3
    delta = (0..3).map { |id| @mod.get_coordinates_from_ids([0, id]) }
    assert.Equal(delta[0][1], delta[2][3])
    assert.Equal(delta[0][1], delta[3][2])
    assert.Equal(delta[0][2], delta[1][3])
    assert.Equal(delta[0][2], delta[3][1])
    assert.Equal(delta[0][3], delta[1][2])
    assert.Equal(delta[0][3], delta[2][1])
  end

  def test_get_coordinates_from_code
    assert.Equal(
      @mod.get_coordinates_from_ids([0]),
      @mod.get_coordinates_from_code("Z"))
    assert.Equal(
      @mod.get_coordinates_from_ids([0, 1, 2]),
      @mod.get_coordinates_from_code("Z8"))
  end

  def test_rush
    1000.times {
      lat1  = rand * 180.0 -  90.0
      lng1  = rand * 360.0 - 180.0
      level = rand(30) + 1
      code1 = @mod.get_delta_code(lat1, lng1, level)
      lat2,
      lng2  = @mod.get_center_from_delta_code(code1)
      code2 = @mod.get_delta_code(lat2, lng2, level)
      lat3,
      lng3  = @mod.get_center_from_delta_code(code2)
      assert.Equal(code1, code2)
      assert.Equal([lat2, lng2], [lat3, lng3])
    }
  end
end
*/
