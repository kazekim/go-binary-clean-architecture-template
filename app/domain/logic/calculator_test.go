package logic

import (
	"github.com/kazekim/go-binary-clean-architecture-template/pkg/jhtest"
	"testing"
)

type AgeTestModel struct {
	yearOfBirth int
}

type AgeTestResult struct {
	age int
}

func TestCalculatorAge2(t *testing.T) {

	units := []*jhtest.TestUnit{
		NewTestUnit(1986, 33),
		NewTestUnit(1926, 93),
		NewTestUnit(1969, 50),
		NewTestUnit(1999, 20),
	}

	function := jhtest.NewTestFunction(t, "Calculate Age Function", units)
	function.DoTest()
}

func (model AgeTestModel) DoTest(t *testing.T) jhtest.TestResult {

	// if Error while calculate, you can use t.Error(message) here
	//t.Error(err.Error())

	age := CalculateAge(model.yearOfBirth)

	result := AgeTestResult{
		age,
	}
	return result
}

func NewTestUnit(yearOfBirth int, expectedAge int) *jhtest.TestUnit {
	return jhtest.NewTestUnit(
		AgeTestModel{
			yearOfBirth: yearOfBirth,
		},
		AgeTestResult{
			age: expectedAge,
		})
}
