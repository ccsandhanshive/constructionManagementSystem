package main

import (
	"testing"

	"bitsys.sys/construction_management_system/serviceImplement"
)

func TestBrickRateChange(t *testing.T) {
	output := serviceImplement.BrickRateChangewithcommit(200, 0)
	output1 := serviceImplement.BrickRateChangewithcommit(200, 1)

	if output != 0 {
		t.Errorf("Output should be 0 got %d", output)
	}

	if output1 != 1 {
		t.Errorf("Output should be 1 got %d", output1)
	}

}
