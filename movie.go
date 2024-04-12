package store

type Movie struct {
	Title             string
	determineAmountFn func(daysRent int) float64
	determinePointsFn func(daysRent int) int
}

func NewNewReleaseMovie(title string) Movie {
	return Movie{
		Title: title,
		determineAmountFn: func(daysRent int) float64 {
			return float64(daysRent) * 3
		},
		determinePointsFn: func(daysRent int) int {
			points := 1
			if daysRent > 1 {
				points += 1
			}
			return points
		},
	}
}

func NewRegularMovie(title string) Movie {
	return Movie{
		Title: title,

		determineAmountFn: func(daysRent int) float64 {
			amount := 2.0
			if daysRent > 2 {
				amount += (float64(daysRent) - 2) * 1.5
			}
			return amount
		},

		determinePointsFn: func(daysRent int) int {
			return 1
		},
	}
}

func NewChildrenMovie(title string) Movie {
	return Movie{
		Title: title,
		determineAmountFn: func(daysRent int) float64 {
			amount := 1.5
			if daysRent > 3 {
				amount += (float64(daysRent) - 3) * 1.5
			}
			return amount
		},
		determinePointsFn: func(daysRent int) int {
			return 1
		},
	}
}

func (this Movie) DetermineAmount(daysRent int) float64 {
	return this.determineAmountFn(daysRent)
}

func (this Movie) DeterminePoints(daysRent int) int {
	return this.determinePointsFn(daysRent)
}
