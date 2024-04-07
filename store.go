package store

import "strconv"

type Customer struct {
	Name        string
	RentalList  []Rental
	totalAmount float64
	points      int
}

func (this *Customer) AddRental(rental Rental) {
	this.RentalList = append(this.RentalList, rental)
}

func (this *Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := "Rental Record for " + this.Name + "\n"

	for _, rental := range this.RentalList {
		thisAmount := 0.0

		switch rental.Movie.PriceCode {
		case PC_REGULAR:
			{
				thisAmount += 2
				if rental.DaysRent > 2 {
					thisAmount += (float64(rental.DaysRent) - 2) * 1.5
				}
			}
		case PC_NEW_RELEASE:
			{
				thisAmount += float64(rental.DaysRent) * 3
			}
		case PC_CHILDREN:
			{
				thisAmount += 1.5
				if rental.DaysRent > 3 {
					thisAmount += (float64(rental.DaysRent) - 3) * 1.5
				}
			}
		}

		frequentRenterPoints++

		if rental.Movie.PriceCode == PC_NEW_RELEASE && rental.DaysRent > 1 {
			frequentRenterPoints++
		}

		result += "\t" + rental.Movie.Title + "\t" + strconv.FormatFloat(thisAmount, 'f', 1, 64) + "\n"
		totalAmount += thisAmount

	}

	this.totalAmount = totalAmount
	this.points = frequentRenterPoints
	result += "You owed " + strconv.FormatFloat(totalAmount, 'f', 1, 64) + "\n"
	result += "You earned " + strconv.Itoa(frequentRenterPoints) + " frequent renter points\n"

	return result
}

func (this *Customer) GetOwed() float64 {
	return this.totalAmount
}

func (this *Customer) GetPoints() int {
	return this.points
}
