package main

func main() {

}

func distanceCalculation(volumeOfTank int, litersPerKm float32 ) float32 {
	distance := float32(volumeOfTank) / litersPerKm

	inaccuracy := 10
	distance = distance - float32(inaccuracy) * distance/100
	return distance
}


