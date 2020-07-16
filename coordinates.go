package coordinates

import "errors"

type Coordinates struct {
	longitude, latitude float64
}

//setter methods:
func (c *Coordinates) SetLongitude(lon float64) error {
	if lon < -90 || lon > 90 {
		return errors.New("Invalid longitude")
	}
	c.longitude = lon
	return nil
}
func (c *Coordinates) SetLatitude(lat float64) error {
	if lat < -90 || lat > 90 {
		return errors.New("Invalid latitude")
	}
	c.latitude = lat
	return nil
}
