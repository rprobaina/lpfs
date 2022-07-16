package lpfs

import (
	"fmt"
	"testing"
)

func TestLoadAverage(t *testing.T) {
	l1, err := GetLoadAverage1()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage1(): %v, err: %v\n", l1, err)

	l5, err := GetLoadAverage5()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage5(): %v, err: %v\n", l5, err)

	l15, err := GetLoadAverage15()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage15(): %v, err: %v\n", l15, err)
}
