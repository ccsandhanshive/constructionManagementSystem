package main

import (
	"testing"

	"bitsys.sys/construction_management_system/serviceImplement"
)

func TestCementCal(t *testing.T) {
	m := serviceImplement.CreateMaterialCal()
	output := m.CementCal(30)
	if output != 2700 {
		t.Errorf("output should be 2700 got %d", output)
	}
}
