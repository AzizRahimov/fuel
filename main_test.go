package main

import "testing"


func Test_distanceCalculation(t *testing.T) {

	tests := []struct {
		name string
		volumeOfTank int
		literForKm float32
		want float32
	}{
		// TODO: Add test cases.
{"distanceForPerKm",150, 0.8, 187.5 },
{"distanceForBmw", 200, 0.5, 400},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := distanceCalculation(test.volumeOfTank, test.literForKm)

			if  got != test.want {
				t.Error("ErrorForFuel", got, test.want)
			}
		})
	}
}