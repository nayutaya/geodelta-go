package geodelta

import "math"

type Lat float64
type Lng float64
type My float64
type Mx float64
type Ny float64
type Nx float64

const DEG2RAD = math.Pi / 180.0         // 度をラジアンに変換するための係数
const RAD2DEG = 180.0 / math.Pi         // ラジアンを度に変換するための係数
const DELTA_HEIGHT = 0.8660254037844386 // 一辺を1.0とする正三角形の高さ math.Sqrt((1*1)-(0.5*0.5))

// 緯度をメルカトルY座標に変換する
//   -90.0 <= lat <= +90.0
//    -1.0 <= my  <=  +1.0
func LatToMy(lat Lat) My {
	return My(math.Atanh(math.Sin(float64(lat)*DEG2RAD)) / math.Pi)
}
func (this Lat) ToMy() My {
	return LatToMy(this)
}

// 経度をメルカトルX座標に変換する
//   -180.0 <= lng <= +180.0
//     -1.0 <= mx  <=   +1.0
func LngToMx(lng Lng) Mx {
	return Mx(float64(lng) / 180.0)
}
func (this Lng) ToMx() Mx {
	return LngToMx(this)
}

// メルカトルX座標/Y座標を正規化する
func NormalizeM(m float64) float64 {
	if m > +1.0 {
		return m - math.Ceil((m-1)/+2.0)*2.0
	} else if m < -1.0 {
		return m + math.Ceil((m+1)/-2.0)*2.0
	} else {
		return m
	}
}

// メルカトルY座標を緯度に変換する
//    -1.0 <= my  <=  +1.0
//   -90.0 <= lat <= +90.0
func MyToLat(my My) Lat {
	return Lat(math.Asin(math.Tanh(NormalizeM(float64(my))*math.Pi)) * RAD2DEG)
}
func (this My) ToLat() Lat {
	return MyToLat(this)
}

// メルカトルX座標を経度に変換する
//     -1.0 <= mx  <=   +1.0
//   -180.0 <= lng <= +180.0
func MxToLng(mx Mx) Lng {
	return Lng(NormalizeM(float64(mx)) * 180.0)
}
func (this Mx) ToLng() Lng {
	return MxToLng(this)
}

// メルカトルY座標から正規化Y座標に変換する
//    -1.0 <= my <=  +1.0
//   -12.0 <= ny <= +12.0
func MyToNy(my My) Ny {
	return Ny(float64(my) / DELTA_HEIGHT * 12.0)
}
func (this My) ToNy() Ny {
	return MyToNy(this)
}

// メルカトルX座標から正規化X座標に変換する
//    -1.0 <= my <=  +1.0
//   -12.0 <= ny <= +12.0
func MxToNx(mx Mx) Nx {
	return Nx(float64(mx) * 12.0)
}
func (this Mx) ToNx() Nx {
	return MxToNx(this)
}

// 正規化Y座標からメルカトルY座標に変換する
//   -12.0 <= ny <= +12.0
//    -1.0 <= my <=  +1.0
func NyToMy(ny Ny) My {
	return My(float64(ny) / 12.0 * DELTA_HEIGHT)
}
func (this Ny) ToMy() My {
	return NyToMy(this)
}

// 正規化X座標からメルカトルX座標に変換する
//   -12.0 <= nx <= +12.0
//    -1.0 <= mx <=  +1.0
func NxToMx(nx Nx) Mx {
	return Mx(float64(nx) / 12.0)
}
func (this Nx) ToMx() Mx {
	return NxToMx(this)
}

func LatToNy(lat Lat) Ny {
	return Ny(MyToNy(LatToMy(lat)))
}
func (this Lat) ToNy() Ny {
	return LatToNy(this)
}

func LngToNx(lng Lng) Nx {
	return Nx(MxToNx(LngToMx(lng)))
}
func (this Lng) ToNx() Nx {
	return LngToNx(this)
}

func NyToLat(ny Ny) Lat {
	return Lat(MyToLat(NyToMy(ny)))
}
func (this Ny) ToLat() Lat {
	return NyToLat(this)
}

func NxToLng(nx Nx) Lng {
	return Lng(MxToLng(NxToMx(nx)))
}
func (this Nx) ToLng() Lng {
	return NxToLng(this)
}

/* Ruby
func LatlngToNxy(lat, lng) {
  return [
    self.lng_to_nx(lng),
    self.lat_to_ny(lat),
  ]
}

func NxyToLatlng(nx, ny) {
  return [
    self.ny_to_lat(ny),
    self.nx_to_lng(nx),
  ]
}
*/
