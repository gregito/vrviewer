package metrics

import (
	"github.com/gregito/vrviewer/comp/log"
	"strconv"
	"time"
)

func ShowMeasurementsIfHaveAny(singleFetchDurations []time.Duration, totalFetchTime time.Duration) {
	if singleFetchDurations == nil {
		log.Println("-- No measured execution has happened --")
		return
	}
	log.Println("--------- API call measurements ---------")
	log.Println("Total measured call amount: " + strconv.Itoa(len(singleFetchDurations)))
	log.Println("Fetching all results took: " + totalFetchTime.String())
	log.Println("Longest fetching took: " + getMaxDuration(singleFetchDurations).String())
	log.Println("Shortest fetching took: " + getMinDuration(singleFetchDurations).String())
	log.Println("Average fetching took: " + getAverageDuration(singleFetchDurations).String())
	log.Println("-----------------------------------------")
}

func getMaxDuration(durations []time.Duration) time.Duration {
	max := durations[0]
	for _, duration := range durations {
		if duration > max {
			max = duration
		}
	}
	return max
}

func getMinDuration(durations []time.Duration) time.Duration {
	min := durations[0]
	for _, duration := range durations {
		if duration < min {
			min = duration
		}
	}
	return min
}

func getAverageDuration(durations []time.Duration) time.Duration {
	avg := int64(durations[0] / time.Millisecond)
	for i := 1; i < len(durations); i++ {
		avg = avg + int64(durations[i]/time.Millisecond)
	}
	return time.Duration(avg/int64(len(durations))) * time.Millisecond
}
