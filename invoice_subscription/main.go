package main

import (
	"fmt"
	"math"
	"time"
)

type Subscription struct {
	Id                    int
	CustomerId            int
	MonthlyPriceInDollars int
}

type User struct {
	Id            int
	Name          string
	ActivatedOn   time.Time
	DeactivatedOn time.Time
	CustomerId    int
}

func BillFor(yearMonth string, activeSubscription *Subscription, users *[]User) float64 {
	var total float64
	layout := "2006-01"
	t, err := time.Parse(layout, yearMonth)
	if err != nil {
		panic(err)
	}

	rate := getDailyRate(activeSubscription.MonthlyPriceInDollars, LastDayOfMonth(t).Day())

	for _, user := range *users {
		if user.CustomerId == activeSubscription.CustomerId {
			start := getFirstDay(t, user.ActivatedOn)
			end := getLastDay(t, user.DeactivatedOn)
			total += math.Floor(float64(end-start+1)*rate*100) / 100
		}
	}

	return total
}

func getLastDay(thisMonth time.Time, lastDay time.Time) int {
	if (lastDay == time.Time{}) {
		return LastDayOfMonth(thisMonth).Day()
	}

	return lastDay.Day()
}

func getFirstDay(thisMonth time.Time, firstDay time.Time) int {
	if thisMonth.Year() > firstDay.Year() {
		return 1
	}

	if firstDay.Month() < thisMonth.Month() {
		return 1
	} else if firstDay.Month() == thisMonth.Month() {
		return firstDay.Day()
	}

	// handle error here
	return 0
}

func getDailyRate(price int, numOfDays int) float64 {
	return float64(float64(price) / float64(numOfDays))
}

func getCustomerUsers(activeSubscription *Subscription, users *[]User) []User {
	var customerUsers []User
	for _, user := range *users {
		if user.CustomerId == activeSubscription.CustomerId {
			customerUsers = append(customerUsers, user)
		}
	}

	return customerUsers
}

/*******************
* Helper functions *
*******************/

/*
Takes a time.Time object and returns a time.Time object
which is the first day of that month.

FirstDayOfMonth(time.Date(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb 7
=> time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC))               // Feb 1
*/
func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

/*
Takes a time.Time object and returns a time.Time object
which is the end of the last day of that month.

LastDayOfMonth(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb  7
=> time.Time(2019, 2, 28, 23, 59, 59, 0, time.UTC)           // Feb 28
*/
func LastDayOfMonth(t time.Time) time.Time {
	return FirstDayOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

/*
Takes a time.Time object and returns a time.Time object
which is the next day.

NextDay(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))   // Feb 7
=> time.Time(2019, 2, 8, 0, 0, 0, 0, time.UTC)         // Feb 8

NextDay(time.Time(2019, 2, 28, 0, 0, 0, 0, time.UTC))  // Feb 28
=> time.Time(2019, 3, 1, 0, 0, 0, 0, time.UTC)         // Mar  1
*/
func NextDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func main() {
	constantUsers := []User{
		{
			Id:            1,
			Name:          "Employee #1",
			ActivatedOn:   time.Date(2018, 11, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            2,
			Name:          "Employee #2",
			ActivatedOn:   time.Date(2018, 12, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
	}
	newPlan := Subscription{
		Id:                    1,
		CustomerId:            1,
		MonthlyPriceInDollars: 4,
	}

	fmt.Println(BillFor("2019-01", &newPlan, &constantUsers))
}
