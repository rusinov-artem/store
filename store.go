package store

import (
	"fmt"
	"strconv"
)

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
	this.clearTotals()
	statement := this.statementHeader() + this.statementDetails()
	totalOwd, points := this.calculateTotal()
	statement += this.footer(totalOwd, points)
	return statement
}

func (this *Customer) clearTotals() {
	this.totalAmount = 0
	this.points = 0
}

func (this *Customer) statementDetails() string {
	details := ""
	for _, rental := range this.RentalList {
		details += this.statementDetail(rental)
	}
	return details
}

func (this *Customer) statementDetail(rental Rental) string {
	return fmt.Sprintf("\t%s\t%s\n",
		rental.Movie.Title,
		strconv.FormatFloat(rental.DetermineAmount(), 'f', 1, 64),
	)
}

func (this *Customer) statementHeader() string {
	return "Rental Record for " + this.Name + "\n"
}

func (this *Customer) GetOwed() float64 {
	return this.totalAmount
}

func (this *Customer) GetPoints() int {
	return this.points
}

func (this *Customer) footer(totalOwed float64, points int) string {
	footer := "You owed " + strconv.FormatFloat(totalOwed, 'f', 1, 64) + "\n"
	footer += "You earned " + strconv.Itoa(points) + " frequent renter points\n"
	return footer
}

func (this *Customer) calculateTotal() (float64, int) {
	for i := range this.RentalList {
		rentalAmount := this.RentalList[i].DetermineAmount()
		earnedPoints := this.RentalList[i].Points()

		this.totalAmount += rentalAmount
		this.points += earnedPoints
	}
	return this.totalAmount, this.points
}
