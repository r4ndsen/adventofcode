package main

type Race struct {
	Time, Distance int
}

func (r Race) WinningConfigurations() int {
	result := 0

	for timeSpentCharging := 1; timeSpentCharging < r.Time-1; timeSpentCharging++ {

		//speed := timeSpentCharging
		//timeToRace := r.Time - timeSpentCharging
		//distanceCovered := speed * timeToRace

		distanceCovered := timeSpentCharging * (r.Time - timeSpentCharging)

		if distanceCovered >= r.Distance {
			result++
		}
	}

	return result
}
