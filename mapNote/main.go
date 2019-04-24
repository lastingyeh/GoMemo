package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type myMap struct {
	mMap map[string]map[string]int
}

func (m myMap) forEach() {
	for k, v := range m.mMap {
		fmt.Printf("L1: [%q]\n", k)
		for k2, v2 := range v {
			fmt.Printf("L2: [%q] : %v\n", k2, v2)

		}
		fmt.Println("------------------------")
	}
}

func (m myMap) length(keyL1 string) (int, error) {
	if keyL1 != "" {
		if _, err := m.find(keyL1, ""); err != nil {
			return -1, err
		}
		return len(m.mMap[keyL1]), nil
	}
	return len(m.mMap), nil
}

func (m myMap) update(keyL1, keyL2 string, updateVal int) error {
	if _, err := m.find(keyL1, keyL2); err != nil {
		return err
	}
	m.mMap[keyL1][keyL2] = updateVal
	return nil
}

func (m myMap) remove(keyL1, keyL2 string) error {
	if _, err := m.find(keyL1, keyL2); err != nil {
		return err
	}
	// as L2 existed, just remove L2
	if keyL2 != "" {
		delete(m.mMap[keyL1], keyL2)
		return nil
	}
	// remove L1
	delete(m.mMap, keyL1)
	return nil
}

func (m myMap) find(keyL1, keyL2 string) (interface{}, error) {
	switch {
	case keyL1 == "":
		return nil, errors.New(fmt.Sprintf("mMap[%q] not allowed empty", keyL1))
	case keyL2 == "":
		val, ok := m.mMap[keyL1]
		if !ok {
			return nil, errors.New(fmt.Sprintf("mMap[%q][%q] not found", keyL1, keyL2))
		}
		return val, nil
	case keyL2 != "":
		val, ok := m.mMap[keyL1][keyL2]
		if !ok {
			return nil, errors.New(fmt.Sprintf("mMap[%q][%q] not found", keyL1, keyL2))
		}
		return val, nil
	}
	return nil, errors.New("unexpected error")
}

func main() {
	m := myMap{
		mMap: map[string]map[string]int{
			"A": {"APPLE": 1, "Ant": 2},
			"B": {"Banana": 5, "Bag": 10},
			"C": {"Cat": 10, "Category": 8},
		},
	}

	keyL1, keyL2 := "A", "APPLE"

	fmt.Println(" ########## foreach all data from myMap ########## ")
	m.forEach()

	fmt.Println(" ########## find key length ########## ")
	if lng, err := m.length(keyL1); err != nil {
		fmt.Printf("length(%q) err = %v\n", keyL1, err)
	} else {
		fmt.Printf("len(%q) = %d\n", keyL1, lng)
	}

	fmt.Println(" ########## update key ########## ")
	if err := m.update(keyL1, keyL2, 3); err != nil {
		fmt.Println("update err = ", err)
	} else {
		fmt.Printf("update [%q][%q] successfully\n", keyL1, keyL2)
	}

	fmt.Println(" ########## find key value ########## ")
	find, err := m.find(keyL1, keyL2)
	if err != nil {
		fmt.Println("update err = ", err)
	}
	fmt.Printf("[%q][%q] = %d\n", keyL1, keyL2, find)

	fmt.Println(" ########## delete key ########## ")
	if err := m.remove(keyL1, keyL2); err != nil {
		fmt.Println("delete err = ", err)
	} else {
		fmt.Printf("remove [%q][%q] successfully\n", keyL1, keyL2)
	}

	fmt.Println(" ########## output ########## ")
	m.forEach()

}
