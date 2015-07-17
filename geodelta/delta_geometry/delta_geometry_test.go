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

/*
class GeoDeltaDeltaGeometryTest < Minitest::Test
  def test_upper_sub_delta?
    assert.Equal(false, @mod.upper_sub_delta?(true,  0))
    assert.Equal(true,  @mod.upper_sub_delta?(true,  1))
    assert.Equal(true,  @mod.upper_sub_delta?(true,  2))
    assert.Equal(true,  @mod.upper_sub_delta?(true,  3))
    assert.Equal(true,  @mod.upper_sub_delta?(false, 0))
    assert.Equal(false, @mod.upper_sub_delta?(false, 1))
    assert.Equal(false, @mod.upper_sub_delta?(false, 2))
    assert.Equal(false, @mod.upper_sub_delta?(false, 3))
  end

  def test_upper_delta?
    assert.Equal(false, @mod.upper_delta?([0]))
    assert.Equal(true,  @mod.upper_delta?([1]))
    assert.Equal(true,  @mod.upper_delta?([4]))
    assert.Equal(false, @mod.upper_delta?([5]))

    assert.Equal(true,  @mod.upper_delta?([0, 0]))
    assert.Equal(false, @mod.upper_delta?([0, 1]))
    assert.Equal(false, @mod.upper_delta?([0, 2]))
    assert.Equal(false, @mod.upper_delta?([0, 3]))

    assert.Equal(false, @mod.upper_delta?([4, 0]))
    assert.Equal(true,  @mod.upper_delta?([4, 1]))
    assert.Equal(true,  @mod.upper_delta?([4, 2]))
    assert.Equal(true,  @mod.upper_delta?([4, 3]))

    assert.Equal(false, @mod.upper_delta?([0, 0, 0]))
    assert.Equal(true,  @mod.upper_delta?([0, 0, 0, 0]))
    assert.Equal(false, @mod.upper_delta?([0, 0, 0, 0, 0]))
  end

  def test_transform_world_delta
    assert.Equal([+6.0, +4.0], @mod.transform_world_delta(0,  +0.0, +4.0))
    assert.Equal([+6.0, +4.0], @mod.transform_world_delta(1,  +6.0, +4.0))
    assert.Equal([+6.0, +4.0], @mod.transform_world_delta(2, +12.0, +4.0))
    assert.Equal([+6.0, +4.0], @mod.transform_world_delta(3, +18.0, +4.0))
    assert.Equal([+6.0, +4.0], @mod.transform_world_delta(4,  +0.0, -8.0))
    assert.Equal([+6.0, +4.0], @mod.transform_world_delta(5,  +6.0, -8.0))
    assert.Equal([+6.0, +4.0], @mod.transform_world_delta(6, +12.0, -8.0))
    assert.Equal([+6.0, +4.0], @mod.transform_world_delta(7, +18.0, -8.0))
  end

  def test_transform_upper_delta
    assert.Equal([+6.0, +8.0], @mod.transform_upper_delta(0, +6.0, +4.0))
    assert.Equal([+6.0, +4.0], @mod.transform_upper_delta(1, +6.0, +8.0))
    assert.Equal([+6.0, +4.0], @mod.transform_upper_delta(2, +9.0, +2.0))
    assert.Equal([+6.0, +4.0], @mod.transform_upper_delta(3, +3.0, +2.0))
  end

  def test_transform_lower_delta
    assert.Equal([+6.0, +4.0], @mod.transform_lower_delta(0, +6.0,  +8.0))
    assert.Equal([+6.0, +8.0], @mod.transform_lower_delta(1, +6.0,  +4.0))
    assert.Equal([+6.0, +8.0], @mod.transform_lower_delta(2, +3.0, +10.0))
    assert.Equal([+6.0, +8.0], @mod.transform_lower_delta(3, +9.0, +10.0))
  end

  def test_get_delta_ids__level1
    assert.Equal([0], @mod.get_delta_ids( 0.0, +6.0, 1))
    assert.Equal([1], @mod.get_delta_ids( 6.0, +6.0, 1))
    assert.Equal([2], @mod.get_delta_ids(12.0, +6.0, 1))
    assert.Equal([3], @mod.get_delta_ids(18.0, +6.0, 1))
    assert.Equal([4], @mod.get_delta_ids( 0.0, -6.0, 1))
    assert.Equal([5], @mod.get_delta_ids( 6.0, -6.0, 1))
    assert.Equal([6], @mod.get_delta_ids(12.0, -6.0, 1))
    assert.Equal([7], @mod.get_delta_ids(18.0, -6.0, 1))
  end

  def test_get_delta_ids__level2
    assert.Equal([0, 0], @mod.get_delta_ids( +0.0,  +8.0, 2))
    assert.Equal([0, 1], @mod.get_delta_ids( +0.0,  +4.0, 2))
    assert.Equal([0, 2], @mod.get_delta_ids( -3.0, +10.0, 2))
    assert.Equal([0, 3], @mod.get_delta_ids( +3.0, +10.0, 2))
    assert.Equal([1, 0], @mod.get_delta_ids( +6.0,  +4.0, 2))
    assert.Equal([1, 1], @mod.get_delta_ids( +6.0,  +8.0, 2))
    assert.Equal([1, 2], @mod.get_delta_ids( +9.0,  +2.0, 2))
    assert.Equal([1, 3], @mod.get_delta_ids( +3.0,  +2.0, 2))
    assert.Equal([2, 2], @mod.get_delta_ids( +9.0, +10.0, 2))
    assert.Equal([3, 3], @mod.get_delta_ids(+15.0,  +2.0, 2))

    assert.Equal([4, 0], @mod.get_delta_ids( +0.0,  -8.0, 2))
    assert.Equal([4, 1], @mod.get_delta_ids( +0.0,  -4.0, 2))
    assert.Equal([4, 2], @mod.get_delta_ids( +3.0, -10.0, 2))
    assert.Equal([4, 3], @mod.get_delta_ids( -3.0, -10.0, 2))
    assert.Equal([5, 0], @mod.get_delta_ids( +6.0,  -4.0, 2))
    assert.Equal([5, 1], @mod.get_delta_ids( +6.0,  -8.0, 2))
    assert.Equal([5, 2], @mod.get_delta_ids( +3.0,  -2.0, 2))
    assert.Equal([5, 3], @mod.get_delta_ids( +9.0,  -2.0, 2))
    assert.Equal([6, 2], @mod.get_delta_ids(+15.0, -10.0, 2))
    assert.Equal([7, 3], @mod.get_delta_ids(+21.0,  -2.0, 2))
  end

  def test_get_delta_ids__level3
    assert.Equal([0, 0, 0], @mod.get_delta_ids(+0.0, +8.0, 3))
    assert.Equal([1, 0, 0], @mod.get_delta_ids(+6.0, +4.0, 3))
  end

  def test_get_delta_ids__level4
    assert.Equal([0, 0, 0, 0], @mod.get_delta_ids(+0.0, +8.0, 4))
    assert.Equal([1, 0, 0, 0], @mod.get_delta_ids(+6.0, +4.0, 4))
  end

  def test_get_world_delta_center
    assert.Equal([ +0.0, +8.0], @mod.get_world_delta_center(0))
    assert.Equal([ +6.0, +4.0], @mod.get_world_delta_center(1))
    assert.Equal([+12.0, +8.0], @mod.get_world_delta_center(2))
    assert.Equal([+18.0, +4.0], @mod.get_world_delta_center(3))
    assert.Equal([ +0.0, -8.0], @mod.get_world_delta_center(4))
    assert.Equal([ +6.0, -4.0], @mod.get_world_delta_center(5))
    assert.Equal([+12.0, -8.0], @mod.get_world_delta_center(6))
    assert.Equal([+18.0, -4.0], @mod.get_world_delta_center(7))
  end

  def test_get_upper_sub_delta_distance
    assert.Equal([+0.0, +0.0], @mod.get_upper_sub_delta_distance(0))
    assert.Equal([+0.0, +4.0], @mod.get_upper_sub_delta_distance(1))
    assert.Equal([+3.0, -2.0], @mod.get_upper_sub_delta_distance(2))
    assert.Equal([-3.0, -2.0], @mod.get_upper_sub_delta_distance(3))
  end

  def test_get_lower_sub_delta_distance
    assert.Equal([+0.0, +0.0], @mod.get_lower_sub_delta_distance(0))
    assert.Equal([+0.0, -4.0], @mod.get_lower_sub_delta_distance(1))
    assert.Equal([-3.0, +2.0], @mod.get_lower_sub_delta_distance(2))
    assert.Equal([+3.0, +2.0], @mod.get_lower_sub_delta_distance(3))
  end

  def test_get_sub_delta_distance
    assert.Equal([+0.0, +4.0], @mod.get_sub_delta_distance(true,  1))
    assert.Equal([+3.0, -2.0], @mod.get_sub_delta_distance(true,  2))
    assert.Equal([+0.0, -4.0], @mod.get_sub_delta_distance(false, 1))
    assert.Equal([-3.0, +2.0], @mod.get_sub_delta_distance(false, 2))
  end

  def test_get_center__level1
    assert.Equal([+0.0, +8.0], @mod.get_center([0]))
    assert.Equal([+6.0, +4.0], @mod.get_center([1]))
    assert.Equal([+0.0, -8.0], @mod.get_center([4]))
    assert.Equal([+6.0, -4.0], @mod.get_center([5]))
  end

  def test_get_center__level2
    assert.Equal([ +0.0,  +8.0], @mod.get_center([0, 0]))
    assert.Equal([ +0.0,  +4.0], @mod.get_center([0, 1]))
    assert.Equal([ -3.0, +10.0], @mod.get_center([0, 2]))
    assert.Equal([ +3.0, +10.0], @mod.get_center([0, 3]))
    assert.Equal([ +6.0,  +4.0], @mod.get_center([1, 0]))
    assert.Equal([ +6.0,  +8.0], @mod.get_center([1, 1]))
    assert.Equal([ +9.0,  +2.0], @mod.get_center([1, 2]))
    assert.Equal([ +3.0,  +2.0], @mod.get_center([1, 3]))
    assert.Equal([ +9.0, +10.0], @mod.get_center([2, 2]))
    assert.Equal([ -9.0,  +2.0], @mod.get_center([3, 3]))

    assert.Equal([ +0.0,  -8.0], @mod.get_center([4, 0]))
    assert.Equal([ +0.0,  -4.0], @mod.get_center([4, 1]))
    assert.Equal([ +3.0, -10.0], @mod.get_center([4, 2]))
    assert.Equal([ -3.0, -10.0], @mod.get_center([4, 3]))
    assert.Equal([ +6.0,  -4.0], @mod.get_center([5, 0]))
    assert.Equal([ +6.0,  -8.0], @mod.get_center([5, 1]))
    assert.Equal([ +3.0,  -2.0], @mod.get_center([5, 2]))
    assert.Equal([ +9.0,  -2.0], @mod.get_center([5, 3]))
    assert.Equal([ -9.0, -10.0], @mod.get_center([6, 2]))
    assert.Equal([ -3.0,  -2.0], @mod.get_center([7, 3]))
  end

  def test_get_center__level3
    assert.Equal([ +0.0,  +8.0], @mod.get_center([0, 0, 0]))
    assert.Equal([ +0.0, +10.0], @mod.get_center([0, 0, 1]))
    assert.Equal([ -1.5,  +5.0], @mod.get_center([0, 1, 2]))
    assert.Equal([ -1.5, +11.0], @mod.get_center([0, 2, 3]))
    assert.Equal([ +3.0, +10.0], @mod.get_center([0, 3, 0]))
  end

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
