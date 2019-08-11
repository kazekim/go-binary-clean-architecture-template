package jhtest

import (
	"fmt"
	"reflect"
	"testing"
)

type TestModel interface {
	DoTest(test *testing.T) TestResult
}

type TestResult interface {
}

type TestUnit struct {
	model     TestModel
	expected  TestResult
	checkFunc TestCheckFunc
}

type TestFunction struct {
	test            *testing.T
	descriptionText string
	units           []*TestUnit
}

type TestCheckFunc func(result TestResult, expected TestResult) error

func NewTestUnitWithFunction(model TestModel, expected TestResult, checkFunc TestCheckFunc) *TestUnit {
	return &TestUnit{
		model:     model,
		expected:  expected,
		checkFunc: checkFunc,
	}
}

func NewTestUnit(model TestModel, expected TestResult) *TestUnit {

	is := func(result TestResult, expected TestResult) error {

		if !reflect.DeepEqual(result, expected) {
			return fmt.Errorf("Expected %v, found %v.", result, expected)
		}
		return nil
	}

	return &TestUnit{
		model:     model,
		expected:  expected,
		checkFunc: is,
	}
}

func NewTestFunction(test *testing.T, desciptionText string, units []*TestUnit) *TestFunction {
	return &TestFunction{
		test:            test,
		descriptionText: desciptionText,
		units:           units,
	}
}

func (t *TestFunction) DoTest() {

	for _, tc := range t.units {
		t.test.Run(t.descriptionText, func(t *testing.T) {
			result := tc.model.DoTest(t)
			if err := tc.checkFunc(result, tc.expected); err != nil {
				t.Error(err)
			}
		})
	}
}
