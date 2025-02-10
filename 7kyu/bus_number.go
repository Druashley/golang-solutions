package sevenkyu

func BusNumber(busStops [][2]int) int {
	// passengerCount := 0

	// for i := 0; i < len(busStops); i++ {
	// 	passengerCount += busStops[i][0]
	// 	passengerCount -= busStops[i][1]
	// }

	// if passengerCount >= 0 {
	// 	return passengerCount
	// } else {
	// 	return 0
	// }
	// first attempt, now I want to try the other way

	passengerCount := 0

	for _, stop := range busStops {
		passengerCount += stop[0]
		passengerCount -= stop[1]
	}

	return passengerCount
}
