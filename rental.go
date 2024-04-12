package store

type Rental struct {
	Movie    Movie
	DaysRent int
}

func (this *Rental) Amount() float64 {
	owed := 0.0
	switch this.Movie.PriceCode {
	case PC_REGULAR:
		{
			owed += 2
			if this.DaysRent > 2 {
				owed += (float64(this.DaysRent) - 2) * 1.5
			}
		}
	case PC_NEW_RELEASE:
		{
			owed += float64(this.DaysRent) * 3
		}
	case PC_CHILDREN:
		{
			owed += 1.5
			if this.DaysRent > 3 {
				owed += (float64(this.DaysRent) - 3) * 1.5
			}
		}
	}

	return owed
}

func (this *Rental) Points() int {
	earnedPoints := 1

	if this.Movie.PriceCode == PC_NEW_RELEASE && this.DaysRent > 1 {
		earnedPoints += 1
	}

	return earnedPoints
}
