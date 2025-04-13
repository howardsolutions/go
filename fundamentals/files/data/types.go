package data

type distance float64
type distanceKm float64

// Method on distance (miles) type
func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60834 * miles)
}

// Method on distanceKm type
func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60834)
}

func test() {
	d := distance(4.5)
	km := d.ToKm()
	km.ToMiles()

	print(d)
}
