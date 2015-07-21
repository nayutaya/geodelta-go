package geodelta

import (
	"github.com/nayutaya/geodelta-go/geodelta/delta_geometry"
	"github.com/nayutaya/geodelta-go/geodelta/encoder"
	"github.com/nayutaya/geodelta-go/geodelta/projector"
)

func GetDeltaIds(lat float64, lng float64, level byte) []byte {
	nx, ny := projector.LatLngToNxNy(projector.Lat(lat), projector.Lng(lng))
	return delta_geometry.GetDeltaIds(float64(nx), float64(ny), level)
}

func GetDeltaCode(lat float64, lng float64, level byte) string {
	ids := GetDeltaIds(lat, lng, level)
	return string(encoder.DeltaIds(ids).Encode())
}

/*
module GeoDelta
  def self.get_center_from_delta_ids(ids)
    nx, ny = GeoDelta::DeltaGeometry.get_center(ids)
    return GeoDelta::Projector.nxy_to_latlng(nx, ny)
  end

  def self.get_center_from_delta_code(code)
    ids = GeoDelta::Encoder.decode(code)
    return self.get_center_from_delta_ids(ids)
  end

  def self.get_coordinates_from_ids(ids)
    return GeoDelta::DeltaGeometry.get_coordinates(ids).
      map { |nx, ny| GeoDelta::Projector.nxy_to_latlng(nx, ny) }
  end

  def self.get_coordinates_from_code(code)
    ids = GeoDelta::Encoder.decode(code)
    return self.get_coordinates_from_ids(ids)
  end
end
*/
