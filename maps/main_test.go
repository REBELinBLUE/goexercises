package main

import (
	"reflect"
	"testing"
)

func TestPeopleOnFloorFive(t *testing.T) {
	actual := peopleOnFloorFive()
	expected := []string{"micheal", "bob", "josh", "charlie", "alex"}

	testPeopleFuncs(actual, expected, t)
}

func TestCharliesTeam(t *testing.T) {
	actual := charliesTeam()
	expected := []string{"alex", "andrew", "bob", "micheal"}

	testPeopleFuncs(actual, expected, t)
}

func TestPeopleWorkingInProduct(t *testing.T) {
	actual := peopleWorkingInProduct()
	expected := []string{"john"}

	testPeopleFuncs(actual, expected, t)
}

func testPeopleFuncs(actual, expected []string, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %s, expected %s", actual, expected)
	}
}
