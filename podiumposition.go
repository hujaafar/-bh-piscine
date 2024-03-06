package piscine

func PodiumPosition(podium [][]string) [][]string {
	positionMap := map[string]int{
		"1st Place": 0,
		"2nd Place": 1,
		"3rd Place": 2,
		"4th Place": 3,
	}

	correctPodium := make([][]string, len(podium))

	for _, position := range podium {
		if len(position) > 0 {
			posStr := position[0]

			index := positionMap[posStr]

			correctPodium[index] = position
		}
	}

	return correctPodium
}
