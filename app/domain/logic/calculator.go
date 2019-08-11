package logic

import (
	"bitbucket.org/inceptionasia/baymax/pkg/wwdate"
)

func CalculateAge(yearOfBirth int) int {

	currentYear := wwdate.GetCurrentYear()

	age := currentYear - yearOfBirth
	return age
}
