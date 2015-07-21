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

func TestGetCenterFromDeltaIds(t *testing.T) {
	assert := assert.New(t)
	var lat, lng float64
	lat, lng = GetCenterFromDeltaIds([]byte{0})
	assert.InDelta(+71.480, lat, 1.0e-3)
	assert.InDelta(+0.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaIds([]byte{1})
	assert.InDelta(+46.024, lat, 1.0e-3)
	assert.InDelta(+90.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaIds([]byte{2})
	assert.InDelta(+71.480, lat, 1.0e-3)
	assert.InDelta(+180.000, lng, 1.0e-3) // TODO: Ruby版では-180.0
	lat, lng = GetCenterFromDeltaIds([]byte{3})
	assert.InDelta(+46.024, lat, 1.0e-3)
	assert.InDelta(-90.000, lng, 1.0e-3)

	lat, lng = GetCenterFromDeltaIds([]byte{4})
	assert.InDelta(-71.480, lat, 1.0e-3)
	assert.InDelta(+0.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaIds([]byte{5})
	assert.InDelta(-46.024, lat, 1.0e-3)
	assert.InDelta(+90.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaIds([]byte{6})
	assert.InDelta(-71.480, lat, 1.0e-3)
	assert.InDelta(+180.000, lng, 1.0e-3) // TODO: Ruby版では+180.0
	lat, lng = GetCenterFromDeltaIds([]byte{7})
	assert.InDelta(-46.024, lat, 1.0e-3)
	assert.InDelta(-90.000, lng, 1.0e-3)

	lat, lng = GetCenterFromDeltaIds([]byte{0, 0})
	assert.InDelta(+71.480, lat, 1.0e-3)
	assert.InDelta(+0.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaIds([]byte{0, 0, 0})
	assert.InDelta(+71.480, lat, 1.0e-3)
	assert.InDelta(+0.000, lng, 1.0e-3)
}

func TestGetCenterFromDeltaCode(t *testing.T) {
	assert := assert.New(t)
	var lat, lng float64
	lat, lng = GetCenterFromDeltaCode("Z")
	assert.InDelta(+71.480, lat, 1.0e-3)
	assert.InDelta(+0.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaCode("Y")
	assert.InDelta(+46.024, lat, 1.0e-3)
	assert.InDelta(+90.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaCode("X")
	assert.InDelta(+71.480, lat, 1.0e-3)
	assert.InDelta(+180.000, lng, 1.0e-3) // TODO: Ruby版では-180.0
	lat, lng = GetCenterFromDeltaCode("W")
	assert.InDelta(+46.024, lat, 1.0e-3)
	assert.InDelta(-90.000, lng, 1.0e-3)

	lat, lng = GetCenterFromDeltaCode("V")
	assert.InDelta(-71.480, lat, 1.0e-3)
	assert.InDelta(+0.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaCode("T")
	assert.InDelta(-46.024, lat, 1.0e-3)
	assert.InDelta(+90.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaCode("S")
	assert.InDelta(-71.480, lat, 1.0e-3)
	assert.InDelta(+180.000, lng, 1.0e-3) // TODO: Ruby版では-180.0
	lat, lng = GetCenterFromDeltaCode("R")
	assert.InDelta(-46.024, lat, 1.0e-3)
	assert.InDelta(-90.000, lng, 1.0e-3)

	lat, lng = GetCenterFromDeltaCode("ZK")
	assert.InDelta(+71.480, lat, 1.0e-3)
	assert.InDelta(+0.000, lng, 1.0e-3)
	lat, lng = GetCenterFromDeltaCode("Z2")
	assert.InDelta(+71.480, lat, 1.0e-3)
	assert.InDelta(+0.000, lng, 1.0e-3)
}

func TestGetCoordinatesFromDeltaIds1(t *testing.T) {
	assert := assert.New(t)
	delta0 := GetCoordinatesFromDeltaIds([]byte{0})
	assert.Equal(4, len(delta0))
	assert.Equal(2, len(delta0[0]))
	assert.Equal(2, len(delta0[1]))
	assert.Equal(2, len(delta0[2]))
	assert.Equal(2, len(delta0[3]))
	assert.InDelta(+71.480, delta0[0][0], 1.0e-3)
	assert.InDelta(+0.000, delta0[0][1], 1.0e-3)
	assert.InDelta(+0.000, delta0[1][0], 1.0e-3)
	assert.InDelta(+0.000, delta0[1][1], 1.0e-3)
	assert.InDelta(+82.467, delta0[2][0], 1.0e-3)
	assert.InDelta(-90.000, delta0[2][1], 1.0e-3)
	assert.InDelta(+82.467, delta0[3][0], 1.0e-3)
	assert.InDelta(+90.000, delta0[3][1], 1.0e-3)

	delta4 := GetCoordinatesFromDeltaIds([]byte{4})
	assert.InDelta(-71.480, delta4[0][0], 1.0e-3)
	assert.InDelta(+0.000, delta4[0][1], 1.0e-3)
	assert.InDelta(+0.000, delta4[1][0], 1.0e-3)
	assert.InDelta(+0.000, delta4[1][1], 1.0e-3)
	assert.InDelta(-82.467, delta4[2][0], 1.0e-3)
	assert.InDelta(+90.000, delta4[2][1], 1.0e-3)
	assert.InDelta(-82.467, delta4[3][0], 1.0e-3)
	assert.InDelta(-90.000, delta4[3][1], 1.0e-3)
}

func TestGetCoordinatesFromDeltaIds2(t *testing.T) {
	assert := assert.New(t)
	var delta [][][]float64
	for id := byte(0); id <= 7; id++ {
		delta = append(delta, GetCoordinatesFromDeltaIds([]byte{id}))
	}
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
	// assert.Equal(delta[1][2], delta[3][3]) // TODO: Ruby版と結果が異なる。
	assert.Equal(delta[1][2], delta[5][3])
	assert.Equal(delta[1][2], delta[6][1])
	// assert.Equal(delta[1][2], delta[7][2]) // TODO: Ruby版と結果が異なる。
	assert.Equal(delta[4][3], delta[6][2])
	assert.Equal(delta[4][3], delta[7][1])
	assert.Equal(delta[4][2], delta[5][1])
	assert.Equal(delta[4][2], delta[6][3])
}

func TestGetCoordinatesFromDeltaIds3(t *testing.T) {
	assert := assert.New(t)
	var delta [][][]float64
	for id := byte(0); id <= 3; id++ {
		delta = append(delta, GetCoordinatesFromDeltaIds([]byte{0, id}))
	}
	assert.Equal(delta[0][1], delta[2][3])
	assert.Equal(delta[0][1], delta[3][2])
	assert.Equal(delta[0][2], delta[1][3])
	assert.Equal(delta[0][2], delta[3][1])
	assert.Equal(delta[0][3], delta[1][2])
	assert.Equal(delta[0][3], delta[2][1])
}

func TestGetCoordinatesFromDeltaCode(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		GetCoordinatesFromDeltaIds([]byte{0}),
		GetCoordinatesFromDeltaCode("Z"))
	assert.Equal(
		GetCoordinatesFromDeltaIds([]byte{0, 1, 2}),
		GetCoordinatesFromDeltaCode("Z8"))
}

/*
class GeoDeltaTest < Minitest::Test
  def test_rush
    1000.times {
      lat1  = rand * 180.0 -  90.0
      lng1  = rand * 360.0 - 180.0
      level = rand(30) + 1
      code1 = @mod.get_delta_code(lat1, lng1, level)
      lat2,
      lng2  = GetCenterFromDeltaCode(code1)
      code2 = @mod.get_delta_code(lat2, lng2, level)
      lat3,
      lng3  = GetCenterFromDeltaCode(code2)
      assert.Equal(code1, code2)
      assert.Equal([lat2, lng2], [lat3, lng3])
    }
  end
end
*/
