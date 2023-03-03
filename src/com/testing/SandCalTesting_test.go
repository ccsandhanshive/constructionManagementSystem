package main

import (
	"testing"

	"bitsys.sys/construction_management_system/serviceImplement"
)

func TestSandCal(t *testing.T) {
	m := serviceImplement.CreateMaterialCal()
	output := m.SandCal(1)
	if output != 350 {
		t.Errorf("Output shoudl be 350 got %d", output)
	}
}
