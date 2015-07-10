package geodelta

import "math"

type Lat float64 //  -90.0 <= lat <=  +90.0
type Lng float64 // -180.0 <= lng <= +180.0
type My float64  //   -1.0 <= my  <=   +1.0
type Mx float64  //   -1.0 <= mx  <=   +1.0
type Ny float64  //  -12.0 <= ny  <=  +12.0
type Nx float64  //  -12.0 <= nx  <=  +12.0

const DEG2RAD = math.Pi / 180.0         // 度をラジアンに変換するための係数
const RAD2DEG = 180.0 / math.Pi         // ラジアンを度に変換するための係数
const DELTA_HEIGHT = 0.8660254037844386 // 一辺を1.0とする正三角形の高さ math.Sqrt((1*1)-(0.5*0.5))

// 緯度をメルカトルY座標に変換する
func (lat Lat) ToMy() My {
	return My(math.Atanh(math.Sin(float64(lat)*DEG2RAD)) / math.Pi)
}

// 経度をメルカトルX座標に変換する
func (lng Lng) ToMx() Mx {
	return Mx(float64(lng) / 180.0)
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
func (my My) ToLat() Lat {
	return Lat(math.Asin(math.Tanh(NormalizeM(float64(my))*math.Pi)) * RAD2DEG)
}

// メルカトルX座標を経度に変換する
func (mx Mx) ToLng() Lng {
	return Lng(NormalizeM(float64(mx)) * 180.0)
}

// メルカトルY座標から正規化Y座標に変換する
func (my My) ToNy() Ny {
	return Ny(float64(my) / DELTA_HEIGHT * 12.0)
}

// メルカトルX座標から正規化X座標に変換する
func (mx Mx) ToNx() Nx {
	return Nx(float64(mx) * 12.0)
}

// 正規化Y座標からメルカトルY座標に変換する
func (ny Ny) ToMy() My {
	return My(float64(ny) / 12.0 * DELTA_HEIGHT)
}

// 正規化X座標からメルカトルX座標に変換する
func (nx Nx) ToMx() Mx {
	return Mx(float64(nx) / 12.0)
}

func (lat Lat) ToNy() Ny {
	return lat.ToMy().ToNy()
}

func (lng Lng) ToNx() Nx {
	return lng.ToMx().ToNx()
}

func (ny Ny) ToLat() Lat {
	return ny.ToMy().ToLat()
}

func (nx Nx) ToLng() Lng {
	return nx.ToMx().ToLng()
}

func LatLngToNxNy(lat Lat, lng Lng) (Nx, Ny) {
	return lng.ToNx(), lat.ToNy()
}

func NxNyToLatLng(nx Nx, ny Ny) (Lat, Lng) {
	return ny.ToLat(), nx.ToLng()
}
