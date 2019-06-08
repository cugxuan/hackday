package models

import (
	"fmt"
	"testing"
)

func TestGetArray(t *testing.T) {
	a := GetArray()
	fmt.Println(a.Array)
}

func TestGetFirst(t *testing.T) {
	a, err := GetFirst()
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println(err)
	}
}

func TestCreateFirst(t *testing.T) {
	CreateFirst()
}
