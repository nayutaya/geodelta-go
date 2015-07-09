package main

import "testing"
import "math"

func TestX(t *testing.T) {
  if math.Abs(x() - 0.0) > 1e-10 {
    t.Error()
  }
}
