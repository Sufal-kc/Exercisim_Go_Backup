package lasagna

// OvenTime is the expected oven time in minutes.
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the actual minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the number of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the total elapsed time (preparation + oven time).
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}