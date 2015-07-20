package delta_geometry

import (
	"math"
)

// 指定された座標(x,y)に該当するワールドデルタの番号を返す
// ただし、-12.0 <= x <= +12.0、-12.0 <= y <= +12.0
func GetWorldDeltaId(x float64, y float64) byte {
	mod := func(a float64, b float64) float64 {
		for ; a >= +b; a -= b {
		}
		for ; a < 0.0; a += b {
		}
		return a
	}
	xx := mod(x, 24.0)
	yy := math.Abs(y)
	var base byte
	if y >= 0.0 {
		base = 0
	} else {
		base = 4
	}
	switch {
	case yy >= +2.0*(xx-0.0):
		return base + 0
	case yy <= -2.0*(xx-12.0):
		return base + 1
	case yy >= +2.0*(xx-12.0):
		return base + 2
	case yy <= -2.0*(xx-24.0):
		return base + 3
	default:
		return base + 0
	}
}

func GetSubDeltaId(upper bool, x float64, y float64) byte {
	if upper {
		return GetUpperDeltaId(x, y)
	} else {
		return GetLowerDeltaId(x, y)
	}
}

// 指定された座標(x,y)に該当する上向きのサブデルタの番号を返す
// ただし、0.0 <= x <= +12.0、0.0 <= y <= +12.0
func GetUpperDeltaId(x float64, y float64) byte {
	switch {
	case y < -2.0*(x-6.0):
		return 3
	case y < +2.0*(x-6.0):
		return 2
	case y > 6.0:
		return 1
	default:
		return 0
	}
}

// 指定された座標(x,y)に該当する下向きのサブデルタの番号を返す
// ただし、0.0 <= x <= +12.0、0.0 <= y <= +12.0
func GetLowerDeltaId(x float64, y float64) byte {
	switch {
	case y > -2.0*(x-12.0):
		return 3
	case y > +2.0*x:
		return 2
	case y < 6.0:
		return 1
	default:
		return 0
	}
}

// 指定されたワールドデルタが上向きかどうかを返す
func IsUpperWorldDelta(id byte) bool {
	if id < 4 {
		return id%2 == 1
	} else {
		return id%2 == 0
	}
}

// 指定されたサブデルタが上向きかどうか返す
func IsUpperSubDelta(parent_is_upper bool, id byte) bool {
	if parent_is_upper {
		return id != 0
	} else {
		return id == 0
	}
}

func IsUpperDelta(ids []byte) bool {
	upper := IsUpperWorldDelta(ids[0])
	for i, length := 1, len(ids); i < length; i++ {
		upper = IsUpperSubDelta(upper, ids[i])
	}
	return upper
}

var TRANSFORM_WORLD_DELTA_X []float64 = []float64{+6.0, +0.0, -6.0, -12.0, +6.0, +0.0, -6.0, -12.0}
var TRANSFORM_WORLD_DELTA_Y []float64 = []float64{+0.0, +0.0, +0.0, +0.0, +12.0, +12.0, +12.0, +12.0}

func TransformWorldDelta(id byte, x float64, y float64) (float64, float64) {
	xx := (x + TRANSFORM_WORLD_DELTA_X[id])
	yy := (y + TRANSFORM_WORLD_DELTA_Y[id])
	return xx, yy
	// TODO: mod
	// xx = (x + TRANSFORM_WORLD_DELTA_X[id]) % 12
	// yy = (y + TRANSFORM_WORLD_DELTA_Y[id]) % 12
}

func TransformSubDelta(upper bool, id byte, x float64, y float64) (float64, float64) {
	if upper {
		return TransformUpperDelta(id, x, y)
	} else {
		return TransformLowerDelta(id, x, y)
	}
}

var TRANSFORM_UPPER_DELTA_X []float64 = []float64{-3.0, -3.0, -6.0, -0.0}
var TRANSFORM_UPPER_DELTA_Y []float64 = []float64{-0.0, -6.0, -0.0, -0.0}

func TransformUpperDelta(id byte, x float64, y float64) (float64, float64) {
	xx := (x + TRANSFORM_UPPER_DELTA_X[id]) * 2
	yy := (y + TRANSFORM_UPPER_DELTA_Y[id]) * 2
	return xx, yy
}

var TRANSFORM_LOWER_DELTA_X []float64 = []float64{-3.0, -3.0, -0.0, -6.0}
var TRANSFORM_LOWER_DELTA_Y []float64 = []float64{-6.0, -0.0, -6.0, -6.0}

func TransformLowerDelta(id byte, x float64, y float64) (float64, float64) {
	xx := (x + TRANSFORM_LOWER_DELTA_X[id]) * 2
	yy := (y + TRANSFORM_LOWER_DELTA_Y[id]) * 2
	return xx, yy
}

func GetDeltaIds(x float64, y float64, level byte) []byte {
	ids := []byte{GetWorldDeltaId(x, y)}
	xx, yy := TransformWorldDelta(ids[0], x, y)
	upper := IsUpperWorldDelta(ids[0])

	for i := byte(2); i <= level; i++ {
		ids = append(ids, GetSubDeltaId(upper, xx, yy))
		xx, yy = TransformSubDelta(upper, ids[len(ids)-1], xx, yy)
		upper = IsUpperSubDelta(upper, ids[len(ids)-1])
	}

	return ids
}

var WORLD_DELTA_CENTER map[byte][]float64 = map[byte][]float64{
	0: []float64{+0.0, +8.0},
	1: []float64{+6.0, +4.0},
	2: []float64{+12.0, +8.0},
	3: []float64{+18.0, +4.0},
	4: []float64{+0.0, -8.0},
	5: []float64{+6.0, -4.0},
	6: []float64{+12.0, -8.0},
	7: []float64{+18.0, -4.0},
}

func GetWorldDeltaCenter(id byte) (float64, float64) {
	xy := WORLD_DELTA_CENTER[id]
	return xy[0], xy[1]
}

var UPPER_SUB_DELTA_DISTANCE map[byte][]float64 = map[byte][]float64{
	0: []float64{+0.0, +0.0},
	1: []float64{+0.0, +4.0},
	2: []float64{+3.0, -2.0},
	3: []float64{-3.0, -2.0},
}

func GetUpperSubDeltaDistance(id byte) (float64, float64) {
	xy := UPPER_SUB_DELTA_DISTANCE[id]
	return xy[0], xy[1]
}

/*
module GeoDelta
  module DeltaGeometry
    LOWER_SUB_DELTA_DISTANCE = {
      0 => [+0.0, +0.0],
      1 => [+0.0, -4.0],
      2 => [-3.0, +2.0],
      3 => [+3.0, +2.0],
    }.freeze.tap { |h| h.values.map(&:freeze) }

    def self.get_lower_sub_delta_distance(id)
      return LOWER_SUB_DELTA_DISTANCE[id]
    end

    def self.get_sub_delta_distance(parent_is_upper, id)
      if parent_is_upper
        return self.get_upper_sub_delta_distance(id)
      else
        return self.get_lower_sub_delta_distance(id)
      end
    end

    def self.get_center(ids)
      w_id, *s_ids = ids

      x, y  = self.get_world_delta_center(w_id)
      upper = self.upper_world_delta?(w_id)
      xs = [x]
      ys = [y]

      s_ids.each.with_index { |id, index|
        x, y  = self.get_sub_delta_distance(upper, id)
        upper = self.upper_sub_delta?(upper, id)
        xs << (x / (2 ** index))
        ys << (y / (2 ** index))
      }

      x = xs.sort.inject(0.0, &:+)
      y = ys.sort.inject(0.0, &:+)

      x -= 24.0 if x > 12.0

      return [x, y]
    end

    def self.get_coordinates(ids)
      cx, cy = self.get_center(ids)
      level  = ids.size
      sign   = (self.upper_delta?(ids) ? +1 : -1)
      scale  = 1.0 / (2 ** (level - 1)) * sign

      dx1 = 0.0
      dy1 = 8.0 * scale
      dx2 = 6.0 * scale
      dy2 = 4.0 * scale

      return [
        [cx,       cy      ],
        [cx + dx1, cy + dy1],
        [cx + dx2, cy - dy2],
        [cx - dx2, cy - dy2],
      ]
    end
  end
end
*/
