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
	this.customer = Customer{Name: "Costumer Name"}
}

func (this *VideoStoreSuite) TestSingleNewReleaseStatement() {
	this.customer.AddRental(Rental{Movie: NewNewReleaseMovie("The Cell"), DaysRent: 3})

	this.AssertOwedAndPoints(9.0, 2)
}

func (this *VideoStoreSuite) TestDualNewReleaseStatement() {
	this.customer.AddRental(Rental{Movie: NewNewReleaseMovie("The Cell"), DaysRent: 3})
	this.customer.AddRental(Rental{Movie: NewNewReleaseMovie("The Tigger Movie"), DaysRent: 3})

	this.AssertOwedAndPoints(18.0, 4)
}

func (this *VideoStoreSuite) TestSingleChildrenStatement() {
	this.customer.AddRental(Rental{Movie: NewChildrenMovie("The Tigger Movie"), DaysRent: 3})

	this.AssertOwedAndPoints(1.5, 1)
}

func (this *VideoStoreSuite) TestMultipleRegularStatement() {
	this.customer.AddRental(Rental{Movie: NewRegularMovie("Plan 9 from Outer Space"), DaysRent: 1})
	this.customer.AddRental(Rental{Movie: NewRegularMovie("8 1/2"), DaysRent: 2})
	this.customer.AddRental(Rental{Movie: NewRegularMovie("Eraserhead"), DaysRent: 3})

	this.AssertOwedAndPoints(7.5, 3)
}

func (this *VideoStoreSuite) TestOutputFormat() {
	this.customer.AddRental(Rental{Movie: NewRegularMovie("Plan 9 from Outer Space"), DaysRent: 1})
	this.customer.AddRental(Rental{Movie: NewRegularMovie("8 1/2"), DaysRent: 2})
	this.customer.AddRental(Rental{Movie: NewRegularMovie("Eraserhead"), DaysRent: 3})

	this.Equal(
		"Rental Record for Costumer Name\n"+
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

func (this *VideoStoreSuite) AssertOwedAndPoints(owed float64, points int) {
	this.T().Helper()
	this.customer.Statement()
	this.InDelta(owed, this.customer.GetOwed(), 0.01)
	this.Equal(points, this.customer.GetPoints())
}
