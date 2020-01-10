package validator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrTooLong         = errors.New("Routing number is too long")
	ErrNotMatch        = errors.New("Routing number must have 9 numeric digits")
	ErrInactive        = errors.New("Routing number is inactive")
	ErrInvalidFirstTwo = errors.New("Invalid first two digits")
	ErrInvalidChecksum = errors.New("Invalid checksum")

	nineDigit = regexp.MustCompile("^\\d{9}$")
)

func ABARoutingNumberIsValid(routingNumberToTest string) error {

	if len(routingNumberToTest) > 9 {
		return ErrTooLong
	}

	routing := routingNumberToTest
	routing = strings.Repeat("0", 9-len(routingNumberToTest)) + routing //I refuse to import left-pad for this

	//gotta be 9  digits
	if !nineDigit.MatchString(routing) {
		return ErrNotMatch
	}

	//all 0's is technically a valid routing number, but it's inactive
	if routing == "000000000" {
		return ErrInactive
	}

	//The first two digits of the nine digit RTN must be in the ranges 00 through 12, 21 through 32, 61 through 72, or 80.
	//https://en.wikipedia.org/wiki/Routing_transit_number
	firstTwo := parseInt(routing[0:2])
	switch {
	case 0 <= firstTwo && firstTwo <= 12:
	case 21 <= firstTwo && firstTwo <= 32:
	case 61 <= firstTwo && firstTwo <= 72:
	case firstTwo == 80:
	default:
		return ErrInvalidFirstTwo
	}

	//this is the checksum
	//http://www.siccolo.com/Articles/SQLScripts/how-to-create-sql-to-calculate-routing-check-digit.html
	weights := []int{3, 7, 1}
	var sum int
	for i := 0; i < 8; i++ {
		sum += parseInt(routing[i:i+1]) * weights[i%3]
	}

	if (10-(sum%10))%10 != parseInt(routing[8:9]) {
		return ErrInvalidChecksum
	}

	return nil
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
