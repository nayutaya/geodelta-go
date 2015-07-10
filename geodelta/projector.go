package geodelta

import "math"

type Lat float64
type Lng float64
type Ny float64
type Nx float64

const DEG2RAD = math.Pi / 180.0         // 度をラジアンに変換するための係数
const RAD2DEG = 180.0 / math.Pi         // ラジアンを度に変換するための係数
const DELTA_HEIGHT = 0.8660254037844386 // 一辺を1.0とする正三角形の高さ math.Sqrt((1*1)-(0.5*0.5))

// 緯度をメルカトルY座標に変換する
//   -90.0 <= lat <= +90.0
//    -1.0 <= my  <=  +1.0
func LatToMy(lat float64) float64 {
	return math.Atanh(math.Sin(lat*DEG2RAD)) / math.Pi
}

// 経度をメルカトルX座標に変換する
//   -180.0 <= lng <= +180.0
//     -1.0 <= mx  <=   +1.0
func LngToMx(lng float64) float64 {
	return lng / 180.0
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
func MyToLat(my float64) float64 {
	return math.Asin(math.Tanh(NormalizeM(my)*math.Pi)) * RAD2DEG
}

// メルカトルX座標を経度に変換する
//     -1.0 <= mx  <=   +1.0
//   -180.0 <= lng <= +180.0
func MxToLng(mx float64) float64 {
	return NormalizeM(mx) * 180.0
}

// メルカトルY座標から正規化Y座標に変換する
//    -1.0 <= my <=  +1.0
//   -12.0 <= ny <= +12.0
func MyToNy(my float64) float64 {
	return my / DELTA_HEIGHT * 12.0
}

// メルカトルX座標から正規化X座標に変換する
//    -1.0 <= my <=  +1.0
//   -12.0 <= ny <= +12.0
func MxToNx(mx float64) float64 {
	return mx * 12.0
}

// 正規化Y座標からメルカトルY座標に変換する
//   -12.0 <= ny <= +12.0
//    -1.0 <= my <=  +1.0
func NyToMy(ny float64) float64 {
	return ny / 12.0 * DELTA_HEIGHT
}

// 正規化X座標からメルカトルX座標に変換する
//   -12.0 <= nx <= +12.0
//    -1.0 <= mx <=  +1.0
func NxToMx(nx float64) float64 {
	return nx / 12.0
}

func LatToNy(lat Lat) Ny {
	return Ny(MyToNy(LatToMy(float64(lat))))
}

func LngToNx(lng Lng) Nx {
	return Nx(MxToNx(LngToMx(float64(lng))))
}

func NyToLat(ny Ny) Lat {
	return Lat(MyToLat(NyToMy(float64(ny))))
}

func NxToLng(nx Nx) Lng {
	return Lng(MxToLng(NxToMx(float64(nx))))
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
