package geocoder

import geo "github.com/Billups/golang-geo"

// This interface describes a Mapper, which should be a data storage mechanism that can execute interesting queries.
// Currently, mappers should be able to find points within a radius of an origin point.
type Mapper interface {
	PointsWithinRadius(p *geo.Point, radius int) bool
}
