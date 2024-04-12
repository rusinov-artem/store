package store

type Rental struct {
	Movie    Movie
	DaysRent int
}

func (this *Rental) DetermineAmount() float64 {
	return this.Movie.DetermineAmount(this.DaysRent)
}

func (this *Rental) Points() int {
	return this.Movie.DeterminePoints(this.DaysRent)
}
