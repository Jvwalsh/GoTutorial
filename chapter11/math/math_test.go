package math

import "testing"

type testpair struct {
  values []float64
  average float64
}

type minTestPairs struct {
	values []float64
	min float64
}

type maxTestPairs struct {
	values []float64
	max float64
}

var tests = []testpair{
	{ []float64{1,2}, 1.5 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
	{ []float64{}, 0 },
}
  
var minTests = []minTestPairs{
	{ []float64{1,2}, 1 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, -1 },
	{ []float64{0,0,3}, 0 },
}

var maxTests = []maxTestPairs{
	{ []float64{1,2}, 2 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 1 },
	{ []float64{0,0,3}, 3 },
}


func TestAverage(t *testing.T) {
  for _, pair := range tests {
    v := Average(pair.values)
    if v != pair.average {
      t.Error(
        "For", pair.values,
        "expected", pair.average,
        "got", v,
      )
    }
  }
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
	  v := Min(pair.values)
	  if v != pair.min {
		t.Error(
		  "For", pair.values,
		  "expected", pair.min,
		  "got", v,
		)
	  }
	}
  }

  func TestMax(t *testing.T) {
	for _, pair := range maxTests {
	  v := Max(pair.values)
	  if v != pair.max {
		t.Error(
		  "For", pair.values,
		  "expected", pair.max,
		  "got", v,
		)
	  }
	}
  }