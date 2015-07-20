package delta_geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWorldDeltaId(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(byte(2), GetWorldDeltaId(-6.0, +12.0))
	assert.Equal(byte(3), GetWorldDeltaId(-6.0, +6.0))
	assert.Equal(byte(3), GetWorldDeltaId(-6.0, 0.0))
	assert.Equal(byte(3), GetWorldDeltaId(-3.0, +6.0))
	assert.Equal(byte(0), GetWorldDeltaId(0.0, +12.0))
	assert.Equal(byte(0), GetWorldDeltaId(0.0, +6.0))
	assert.Equal(byte(0), GetWorldDeltaId(0.0, 0.0))
	assert.Equal(byte(0), GetWorldDeltaId(+3.0, +6.0))
	assert.Equal(byte(0), GetWorldDeltaId(+6.0, +12.0))
	assert.Equal(byte(1), GetWorldDeltaId(+6.0, +6.0))
	assert.Equal(byte(1), GetWorldDeltaId(+6.0, 0.0))
	assert.Equal(byte(1), GetWorldDeltaId(+9.0, +6.0))
	assert.Equal(byte(1), GetWorldDeltaId(+12.0, 0.0))
	assert.Equal(byte(2), GetWorldDeltaId(+12.0, +12.0))
	assert.Equal(byte(2), GetWorldDeltaId(+12.0, +6.0))
	assert.Equal(byte(2), GetWorldDeltaId(+15.0, +6.0))
	assert.Equal(byte(2), GetWorldDeltaId(+18.0, +12.0))
	assert.Equal(byte(3), GetWorldDeltaId(+18.0, +6.0))
	assert.Equal(byte(3), GetWorldDeltaId(+18.0, 0.0))
	assert.Equal(byte(3), GetWorldDeltaId(+21.0, +6.0))
	assert.Equal(byte(0), GetWorldDeltaId(+24.0, +12.0))
	assert.Equal(byte(0), GetWorldDeltaId(+24.0, +6.0))
	assert.Equal(byte(0), GetWorldDeltaId(+24.0, 0.0))

	assert.Equal(byte(6), GetWorldDeltaId(-6.0, -12.0))
	assert.Equal(byte(7), GetWorldDeltaId(-6.0, -6.0))
	assert.Equal(byte(7), GetWorldDeltaId(-3.0, -6.0))
	assert.Equal(byte(4), GetWorldDeltaId(0.0, -6.0))
	assert.Equal(byte(4), GetWorldDeltaId(0.0, -12.0))
	assert.Equal(byte(4), GetWorldDeltaId(+3.0, -6.0))
	assert.Equal(byte(4), GetWorldDeltaId(+6.0, -12.0))
	assert.Equal(byte(5), GetWorldDeltaId(+6.0, -6.0))
	assert.Equal(byte(5), GetWorldDeltaId(+9.0, -6.0))
	assert.Equal(byte(6), GetWorldDeltaId(+12.0, -6.0))
	assert.Equal(byte(6), GetWorldDeltaId(+12.0, -12.0))
	assert.Equal(byte(6), GetWorldDeltaId(+15.0, -6.0))
	assert.Equal(byte(6), GetWorldDeltaId(+18.0, -12.0))
	assert.Equal(byte(7), GetWorldDeltaId(+18.0, -6.0))
	assert.Equal(byte(7), GetWorldDeltaId(+21.0, -6.0))
	assert.Equal(byte(4), GetWorldDeltaId(+24.0, -6.0))
	assert.Equal(byte(4), GetWorldDeltaId(+24.0, -12.0))
}

func TestGetUpperDeltaId(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(byte(3), GetUpperDeltaId(0.0, 0.0))
	assert.Equal(byte(3), GetUpperDeltaId(1.5, 3.0))
	assert.Equal(byte(3), GetUpperDeltaId(3.0, 3.0))
	assert.Equal(byte(3), GetUpperDeltaId(3.0, 0.0))
	assert.Equal(byte(2), GetUpperDeltaId(9.0, 3.0))
	assert.Equal(byte(2), GetUpperDeltaId(9.0, 0.0))
	assert.Equal(byte(2), GetUpperDeltaId(10.5, 3.0))
	assert.Equal(byte(2), GetUpperDeltaId(12.0, 0.0))
	assert.Equal(byte(1), GetUpperDeltaId(4.5, 9.0))
	assert.Equal(byte(1), GetUpperDeltaId(6.0, 12.0))
	assert.Equal(byte(1), GetUpperDeltaId(6.0, 9.0))
	assert.Equal(byte(1), GetUpperDeltaId(7.5, 9.0))
	assert.Equal(byte(0), GetUpperDeltaId(3.0, 6.0))
	assert.Equal(byte(0), GetUpperDeltaId(4.5, 3.0))
	assert.Equal(byte(0), GetUpperDeltaId(6.0, 6.0))
	assert.Equal(byte(0), GetUpperDeltaId(6.0, 3.0))
	assert.Equal(byte(0), GetUpperDeltaId(6.0, 0.0))
	assert.Equal(byte(0), GetUpperDeltaId(7.5, 3.0))
	assert.Equal(byte(0), GetUpperDeltaId(9.0, 6.0))
}

func TestGetLowerDeltaId(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(byte(3), GetLowerDeltaId(9.0, 12.0))
	assert.Equal(byte(3), GetLowerDeltaId(9.0, 9.0))
	assert.Equal(byte(3), GetLowerDeltaId(10.5, 9.0))
	assert.Equal(byte(3), GetLowerDeltaId(12.0, 12.0))
	assert.Equal(byte(2), GetLowerDeltaId(0.0, 12.0))
	assert.Equal(byte(2), GetLowerDeltaId(1.5, 9.0))
	assert.Equal(byte(2), GetLowerDeltaId(3.0, 12.0))
	assert.Equal(byte(2), GetLowerDeltaId(3.0, 9.0))
	assert.Equal(byte(1), GetLowerDeltaId(4.5, 3.0))
	assert.Equal(byte(1), GetLowerDeltaId(6.0, 3.0))
	assert.Equal(byte(1), GetLowerDeltaId(6.0, 0.0))
	assert.Equal(byte(1), GetLowerDeltaId(7.5, 3.0))
	assert.Equal(byte(0), GetLowerDeltaId(3.0, 6.0))
	assert.Equal(byte(0), GetLowerDeltaId(4.5, 9.0))
	assert.Equal(byte(0), GetLowerDeltaId(6.0, 12.0))
	assert.Equal(byte(0), GetLowerDeltaId(6.0, 9.0))
	assert.Equal(byte(0), GetLowerDeltaId(6.0, 6.0))
	assert.Equal(byte(0), GetLowerDeltaId(7.5, 9.0))
	assert.Equal(byte(0), GetLowerDeltaId(9.0, 6.0))
}

func TestIsUpperWorldDelta(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(false, IsUpperWorldDelta(0))
	assert.Equal(true, IsUpperWorldDelta(1))
	assert.Equal(false, IsUpperWorldDelta(2))
	assert.Equal(true, IsUpperWorldDelta(3))
	assert.Equal(true, IsUpperWorldDelta(4))
	assert.Equal(false, IsUpperWorldDelta(5))
	assert.Equal(true, IsUpperWorldDelta(6))
	assert.Equal(false, IsUpperWorldDelta(7))
}

func TestIsUpperSubDelta(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(false, IsUpperSubDelta(true, 0))
	assert.Equal(true, IsUpperSubDelta(true, 1))
	assert.Equal(true, IsUpperSubDelta(true, 2))
	assert.Equal(true, IsUpperSubDelta(true, 3))
	assert.Equal(true, IsUpperSubDelta(false, 0))
	assert.Equal(false, IsUpperSubDelta(false, 1))
	assert.Equal(false, IsUpperSubDelta(false, 2))
	assert.Equal(false, IsUpperSubDelta(false, 3))
}

func TestIsUpperDelta(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(false, IsUpperDelta([]byte{0}))
	assert.Equal(true, IsUpperDelta([]byte{1}))
	assert.Equal(true, IsUpperDelta([]byte{4}))
	assert.Equal(false, IsUpperDelta([]byte{5}))

	assert.Equal(true, IsUpperDelta([]byte{0, 0}))
	assert.Equal(false, IsUpperDelta([]byte{0, 1}))
	assert.Equal(false, IsUpperDelta([]byte{0, 2}))
	assert.Equal(false, IsUpperDelta([]byte{0, 3}))

	assert.Equal(false, IsUpperDelta([]byte{4, 0}))
	assert.Equal(true, IsUpperDelta([]byte{4, 1}))
	assert.Equal(true, IsUpperDelta([]byte{4, 2}))
	assert.Equal(true, IsUpperDelta([]byte{4, 3}))

	assert.Equal(false, IsUpperDelta([]byte{0, 0, 0}))
	assert.Equal(true, IsUpperDelta([]byte{0, 0, 0, 0}))
	assert.Equal(false, IsUpperDelta([]byte{0, 0, 0, 0, 0}))
}

func TestTransformWorldDelta(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = TransformWorldDelta(0, +0.0, +4.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformWorldDelta(1, +6.0, +4.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformWorldDelta(2, +12.0, +4.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformWorldDelta(3, +18.0, +4.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformWorldDelta(4, +0.0, -8.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformWorldDelta(5, +6.0, -8.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformWorldDelta(6, +12.0, -8.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformWorldDelta(7, +18.0, -8.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
}

func TestTransformUpperDelta(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = TransformUpperDelta(0, +6.0, +4.0)
	assert.Equal([]float64{+6.0, +8.0}, []float64{x, y})
	x, y = TransformUpperDelta(1, +6.0, +8.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformUpperDelta(2, +9.0, +2.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformUpperDelta(3, +3.0, +2.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
}

func TestTransformLowerDelta(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = TransformLowerDelta(0, +6.0, +8.0)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = TransformLowerDelta(1, +6.0, +4.0)
	assert.Equal([]float64{+6.0, +8.0}, []float64{x, y})
	x, y = TransformLowerDelta(2, +3.0, +10.0)
	assert.Equal([]float64{+6.0, +8.0}, []float64{x, y})
	x, y = TransformLowerDelta(3, +9.0, +10.0)
	assert.Equal([]float64{+6.0, +8.0}, []float64{x, y})
}

func TestGetDeltaIdsLevel1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0}, GetDeltaIds(0.0, +6.0, 1))
	assert.Equal([]byte{1}, GetDeltaIds(6.0, +6.0, 1))
	assert.Equal([]byte{2}, GetDeltaIds(12.0, +6.0, 1))
	assert.Equal([]byte{3}, GetDeltaIds(18.0, +6.0, 1))
	assert.Equal([]byte{4}, GetDeltaIds(0.0, -6.0, 1))
	assert.Equal([]byte{5}, GetDeltaIds(6.0, -6.0, 1))
	assert.Equal([]byte{6}, GetDeltaIds(12.0, -6.0, 1))
	assert.Equal([]byte{7}, GetDeltaIds(18.0, -6.0, 1))
}

func TestGetDeltaIdsLevel2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0, 0}, GetDeltaIds(+0.0, +8.0, 2))
	assert.Equal([]byte{0, 1}, GetDeltaIds(+0.0, +4.0, 2))
	assert.Equal([]byte{0, 2}, GetDeltaIds(-3.0, +10.0, 2))
	assert.Equal([]byte{0, 3}, GetDeltaIds(+3.0, +10.0, 2))
	assert.Equal([]byte{1, 0}, GetDeltaIds(+6.0, +4.0, 2))
	assert.Equal([]byte{1, 1}, GetDeltaIds(+6.0, +8.0, 2))
	assert.Equal([]byte{1, 2}, GetDeltaIds(+9.0, +2.0, 2))
	assert.Equal([]byte{1, 3}, GetDeltaIds(+3.0, +2.0, 2))
	assert.Equal([]byte{2, 2}, GetDeltaIds(+9.0, +10.0, 2))
	assert.Equal([]byte{3, 3}, GetDeltaIds(+15.0, +2.0, 2))

	assert.Equal([]byte{4, 0}, GetDeltaIds(+0.0, -8.0, 2))
	assert.Equal([]byte{4, 1}, GetDeltaIds(+0.0, -4.0, 2))
	assert.Equal([]byte{4, 2}, GetDeltaIds(+3.0, -10.0, 2))
	assert.Equal([]byte{4, 3}, GetDeltaIds(-3.0, -10.0, 2))
	assert.Equal([]byte{5, 0}, GetDeltaIds(+6.0, -4.0, 2))
	assert.Equal([]byte{5, 1}, GetDeltaIds(+6.0, -8.0, 2))
	assert.Equal([]byte{5, 2}, GetDeltaIds(+3.0, -2.0, 2))
	assert.Equal([]byte{5, 3}, GetDeltaIds(+9.0, -2.0, 2))
	assert.Equal([]byte{6, 2}, GetDeltaIds(+15.0, -10.0, 2))
	assert.Equal([]byte{7, 3}, GetDeltaIds(+21.0, -2.0, 2))
}

func TestGetDeltaIdsLevel3(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0, 0, 0}, GetDeltaIds(+0.0, +8.0, 3))
	assert.Equal([]byte{1, 0, 0}, GetDeltaIds(+6.0, +4.0, 3))
}

func TestGetDeltaIdsLevel4(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0, 0, 0, 0}, GetDeltaIds(+0.0, +8.0, 4))
	assert.Equal([]byte{1, 0, 0, 0}, GetDeltaIds(+6.0, +4.0, 4))
}

func TestGetWorldDeltaCenter(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = GetWorldDeltaCenter(0)
	assert.Equal([]float64{+0.0, +8.0}, []float64{x, y})
	x, y = GetWorldDeltaCenter(1)
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = GetWorldDeltaCenter(2)
	assert.Equal([]float64{+12.0, +8.0}, []float64{x, y})
	x, y = GetWorldDeltaCenter(3)
	assert.Equal([]float64{+18.0, +4.0}, []float64{x, y})
	x, y = GetWorldDeltaCenter(4)
	assert.Equal([]float64{+0.0, -8.0}, []float64{x, y})
	x, y = GetWorldDeltaCenter(5)
	assert.Equal([]float64{+6.0, -4.0}, []float64{x, y})
	x, y = GetWorldDeltaCenter(6)
	assert.Equal([]float64{+12.0, -8.0}, []float64{x, y})
	x, y = GetWorldDeltaCenter(7)
	assert.Equal([]float64{+18.0, -4.0}, []float64{x, y})
}

func TestGetUpperSubDeltaDistance(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = GetUpperSubDeltaDistance(0)
	assert.Equal([]float64{+0.0, +0.0}, []float64{x, y})
	x, y = GetUpperSubDeltaDistance(1)
	assert.Equal([]float64{+0.0, +4.0}, []float64{x, y})
	x, y = GetUpperSubDeltaDistance(2)
	assert.Equal([]float64{+3.0, -2.0}, []float64{x, y})
	x, y = GetUpperSubDeltaDistance(3)
	assert.Equal([]float64{-3.0, -2.0}, []float64{x, y})
}

func TestGetLowerSubDeltaDistance(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = GetLowerSubDeltaDistance(0)
	assert.Equal([]float64{+0.0, +0.0}, []float64{x, y})
	x, y = GetLowerSubDeltaDistance(1)
	assert.Equal([]float64{+0.0, -4.0}, []float64{x, y})
	x, y = GetLowerSubDeltaDistance(2)
	assert.Equal([]float64{-3.0, +2.0}, []float64{x, y})
	x, y = GetLowerSubDeltaDistance(3)
	assert.Equal([]float64{+3.0, +2.0}, []float64{x, y})
}

func TestGetSubDeltaDistance(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = GetSubDeltaDistance(true, 1)
	assert.Equal([]float64{+0.0, +4.0}, []float64{x, y})
	x, y = GetSubDeltaDistance(true, 2)
	assert.Equal([]float64{+3.0, -2.0}, []float64{x, y})
	x, y = GetSubDeltaDistance(false, 1)
	assert.Equal([]float64{+0.0, -4.0}, []float64{x, y})
	x, y = GetSubDeltaDistance(false, 2)
	assert.Equal([]float64{-3.0, +2.0}, []float64{x, y})
}

func TestGetCenterLevel1(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = GetCenter([]byte{0})
	assert.Equal([]float64{+0.0, +8.0}, []float64{x, y})
	x, y = GetCenter([]byte{1})
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = GetCenter([]byte{4})
	assert.Equal([]float64{+0.0, -8.0}, []float64{x, y})
	x, y = GetCenter([]byte{5})
	assert.Equal([]float64{+6.0, -4.0}, []float64{x, y})
}

func TestGetCenterLevel2(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = GetCenter([]byte{0, 0})
	assert.Equal([]float64{+0.0, +8.0}, []float64{x, y})
	x, y = GetCenter([]byte{0, 1})
	assert.Equal([]float64{+0.0, +4.0}, []float64{x, y})
	x, y = GetCenter([]byte{0, 2})
	assert.Equal([]float64{-3.0, +10.0}, []float64{x, y})
	x, y = GetCenter([]byte{0, 3})
	assert.Equal([]float64{+3.0, +10.0}, []float64{x, y})
	x, y = GetCenter([]byte{1, 0})
	assert.Equal([]float64{+6.0, +4.0}, []float64{x, y})
	x, y = GetCenter([]byte{1, 1})
	assert.Equal([]float64{+6.0, +8.0}, []float64{x, y})
	x, y = GetCenter([]byte{1, 2})
	assert.Equal([]float64{+9.0, +2.0}, []float64{x, y})
	x, y = GetCenter([]byte{1, 3})
	assert.Equal([]float64{+3.0, +2.0}, []float64{x, y})
	x, y = GetCenter([]byte{2, 2})
	assert.Equal([]float64{+9.0, +10.0}, []float64{x, y})
	x, y = GetCenter([]byte{3, 3})
	assert.Equal([]float64{-9.0, +2.0}, []float64{x, y})

	x, y = GetCenter([]byte{4, 0})
	assert.Equal([]float64{+0.0, -8.0}, []float64{x, y})
	x, y = GetCenter([]byte{4, 1})
	assert.Equal([]float64{+0.0, -4.0}, []float64{x, y})
	x, y = GetCenter([]byte{4, 2})
	assert.Equal([]float64{+3.0, -10.0}, []float64{x, y})
	x, y = GetCenter([]byte{4, 3})
	assert.Equal([]float64{-3.0, -10.0}, []float64{x, y})
	x, y = GetCenter([]byte{5, 0})
	assert.Equal([]float64{+6.0, -4.0}, []float64{x, y})
	x, y = GetCenter([]byte{5, 1})
	assert.Equal([]float64{+6.0, -8.0}, []float64{x, y})
	x, y = GetCenter([]byte{5, 2})
	assert.Equal([]float64{+3.0, -2.0}, []float64{x, y})
	x, y = GetCenter([]byte{5, 3})
	assert.Equal([]float64{+9.0, -2.0}, []float64{x, y})
	x, y = GetCenter([]byte{6, 2})
	assert.Equal([]float64{-9.0, -10.0}, []float64{x, y})
	x, y = GetCenter([]byte{7, 3})
	assert.Equal([]float64{-3.0, -2.0}, []float64{x, y})
}

func TestGetCenterLevel3(t *testing.T) {
	assert := assert.New(t)
	var x, y float64
	x, y = GetCenter([]byte{0, 0, 0})
	assert.Equal([]float64{+0.0, +8.0}, []float64{x, y})
	x, y = GetCenter([]byte{0, 0, 1})
	assert.Equal([]float64{+0.0, +10.0}, []float64{x, y})
	x, y = GetCenter([]byte{0, 1, 2})
	assert.Equal([]float64{-1.5, +5.0}, []float64{x, y})
	x, y = GetCenter([]byte{0, 2, 3})
	assert.Equal([]float64{-1.5, +11.0}, []float64{x, y})
	x, y = GetCenter([]byte{0, 3, 0})
	assert.Equal([]float64{+3.0, +10.0}, []float64{x, y})
}

/*
class GeoDeltaDeltaGeometryTest < Minitest::Test
  def test_get_coordinates__level1
    expected = [
      [+0.0,  +8.0],
      [+0.0,  +0.0], # +0.0, -8.0
      [-6.0, +12.0], # -6.0, +4.0
      [+6.0, +12.0], # +6.0, +4.0
    ]
    assert.Equal(expected, @mod.get_coordinates([0]))

    expected = [
      [ +6.0,  +4.0],
      [ +6.0, +12.0], # +0.0, +8.0
      [+12.0,  +0.0], # +6.0, -4.0
      [ +0.0,  +0.0], # -6.0, -4.0
    ]
    assert.Equal(expected, @mod.get_coordinates([1]))

    expected = [
      [+0.0,  -8.0],
      [+0.0,  +0.0], # +0.0, +8.0
      [+6.0, -12.0], # +6.0, -4.0
      [-6.0, -12.0], # -6.0, -4.0
    ]
    assert.Equal(expected, @mod.get_coordinates([4]))

    expected = [
      [ +6.0,  -4.0],
      [ +6.0, -12.0], # +0.0, -8.0
      [ +0.0,  +0.0], # -6.0, +4.0
      [+12.0,  +0.0], # +6.0, +4.0
    ]
    assert.Equal(expected, @mod.get_coordinates([5]))
  end

  def test_get_coordinates__level2
    expected = [
      [ +0.0,  +8.0],
      [ +0.0, +12.0], # +0.0, +4.0
      [ +3.0,  +6.0], # +3.0, -2.0
      [ -3.0,  +6.0], # -3.0, -2.0
    ]
    assert.Equal(expected, @mod.get_coordinates([0, 0]))

    expected = [
      [ +0.0,  +4.0],
      [ +0.0,  +0.0], # +0.0, -4.0
      [ -3.0,  +6.0], # -3.0, +2.0
      [ +3.0,  +6.0], # +3.0, +2.0
    ]
    assert.Equal(expected, @mod.get_coordinates([0, 1]))

    expected = [
      [ -3.0, +10.0],
      [ -3.0,  +6.0], # +0.0, -4.0
      [ -6.0, +12.0], # -3.0, +2.0
      [ +0.0, +12.0], # +3.0, +2.0
    ]
    assert.Equal(expected, @mod.get_coordinates([0, 2]))

    expected = [
      [ +3.0, +10.0],
      [ +3.0,  +6.0], # +0.0, -4.0
      [ +0.0, +12.0], # -3.0, +2.0
      [ +6.0, +12.0], # +3.0, +2.0
    ]
    assert.Equal(expected, @mod.get_coordinates([0, 3]))

    expected = [
      [ +0.0,  -8.0],
      [ +0.0, -12.0], # +0.0, -4.0
      [ -3.0,  -6.0], # -3.0, +2.0
      [ +3.0,  -6.0], # +3.0, +2.0
    ]
    assert.Equal(expected, @mod.get_coordinates([4, 0]))

    expected = [
      [ +0.0,  -4.0],
      [ +0.0,  +0.0], # +0.0, +4.0
      [ +3.0,  -6.0], # +3.0, -2.0
      [ -3.0,  -6.0], # -3.0, -2.0
    ]
    assert.Equal(expected, @mod.get_coordinates([4, 1]))

    expected = [
      [ +3.0, -10.0],
      [ +3.0,  -6.0], # +0.0, +4.0
      [ +6.0, -12.0], # +3.0, -2.0
      [ +0.0, -12.0], # -3.0, -2.0
    ]
    assert.Equal(expected, @mod.get_coordinates([4, 2]))

    expected = [
      [ -3.0, -10.0],
      [ -3.0,  -6.0], # +0.0, +4.0
      [ +0.0, -12.0], # +3.0, -2.0
      [ -6.0, -12.0], # -3.0, -2.0
    ]
    assert.Equal(expected, @mod.get_coordinates([4, 3]))
  end

  def test_get_coordinates__level3
    expected = [
      [ +0.0,  +8.0],
      [ +0.0,  +6.0], # +0.0, -2.0
      [ -1.5,  +9.0], # -1.5, +1.0
      [ +1.5,  +9.0], # +1.5, +1.0
    ]
    assert.Equal(expected, @mod.get_coordinates([0, 0, 0]))

    expected = [
      [ -1.5,  +5.0],
      [ -1.5,  +3.0], # +0.0, -2.0
      [ -3.0,  +6.0], # -1.5, +1.0
      [ +0.0,  +6.0], # +1.5, +1.0
    ]
    assert.Equal(expected, @mod.get_coordinates([0, 1, 2]))
  end

  def test_rush__center
    1000.times {
      x1     = (rand * 24) - 12
      y1     = (rand * 24) - 12
      level  = rand(20) + 1
      ids1   = @mod.get_delta_ids(x1, y1, level)
      x2, y2 = @mod.get_center(ids1)
      ids2   = @mod.get_delta_ids(x2, y2, level)
      x3, y3 = @mod.get_center(ids2)
      assert.Equal(ids1, ids2)
      assert.Equal([x2, y2], [x3, y3])
    }
  end
end
*/
