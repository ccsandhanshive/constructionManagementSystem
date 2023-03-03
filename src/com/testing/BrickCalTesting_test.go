package main

import (
	"testing"

	"bitsys.sys/construction_management_system/serviceImplement"
)

func TestBrickCal(t *testing.T) {
	m := serviceImplement.CreateMaterialCal()
	output := m.BrickCal(30)
	if output != 90 {
		t.Errorf("output should be 90 got %d", output)
	}

}
