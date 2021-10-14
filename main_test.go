package main

import (
	"iris-demo/unit_test"
	"testing"
)

func TestStart(t *testing.T)  {
	newApp()
	t.Run("BootStart", unit_test.BootStart)
}
