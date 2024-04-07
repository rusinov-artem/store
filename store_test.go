package store

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type VideoStoreSuite struct {
	suite.Suite
	customer Customer
}

func (this *VideoStoreSuite) SetupTest() {
	this.customer = Customer{Name: "Fred"}
}

func (this *VideoStoreSuite) TestSingleNewReleaseStatement() {
	this.customer.AddRental(Rental{Movie: Movie{Title: "The Cell", PriceCode: PC_NEW_RELEASE}, DaysRent: 3})
	this.Equal(
		"Rental Record for Fred\n"+
			"\tThe Cell\t9.0\n"+
			"You owed 9.0\n"+
			"You earned 2 frequent renter points\n",
		this.customer.Statement(),
	)
}

func (this *VideoStoreSuite) TestDualNewReleaseStatement() {
	this.customer.AddRental(Rental{Movie: Movie{Title: "The Cell", PriceCode: PC_NEW_RELEASE}, DaysRent: 3})
	this.customer.AddRental(Rental{Movie: Movie{Title: "The Tigger Movie", PriceCode: PC_NEW_RELEASE}, DaysRent: 3})
	this.Equal(
		"Rental Record for Fred\n"+
			"\tThe Cell\t9.0\n"+
			"\tThe Tigger Movie\t9.0\n"+
			"You owed 18.0\n"+
			"You earned 4 frequent renter points\n",
		this.customer.Statement(),
	)
}

func (this *VideoStoreSuite) TestSingleChildrenStatement() {
	this.customer.AddRental(Rental{Movie: Movie{Title: "The Tigger Movie", PriceCode: PC_CHILDREN}, DaysRent: 3})
	this.Equal(
		"Rental Record for Fred\n"+
			"\tThe Tigger Movie\t1.5\n"+
			"You owed 1.5\n"+
			"You earned 1 frequent renter points\n",
		this.customer.Statement(),
	)
}

func (this *VideoStoreSuite) TestMultipleRegularStatement() {
	this.customer.AddRental(Rental{Movie: Movie{Title: "Plan 9 from Outer Space", PriceCode: PC_REGULAR}, DaysRent: 1})
	this.customer.AddRental(Rental{Movie: Movie{Title: "8 1/2", PriceCode: PC_REGULAR}, DaysRent: 2})
	this.customer.AddRental(Rental{Movie: Movie{Title: "Eraserhead", PriceCode: PC_REGULAR}, DaysRent: 3})

	this.Equal(
		"Rental Record for Fred\n"+
			"\tPlan 9 from Outer Space\t2.0\n"+
			"\t8 1/2\t2.0\n"+
			"\tEraserhead\t3.5\n"+
			"You owed 7.5\n"+
			"You earned 3 frequent renter points\n",
		this.customer.Statement(),
	)
}

func TestVideoStore(t *testing.T) {
	suite.Run(t, &VideoStoreSuite{})
}
