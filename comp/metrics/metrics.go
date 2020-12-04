package metrics

import "time"

func GetMaxDuration(durations []time.Duration) time.Duration {
	max := durations[0]
	for _, duration := range durations {
		if duration > max {
			max = duration
		}
	}
	return max
}

func GetMinDuration(durations []time.Duration) time.Duration {
	min := durations[0]
	for _, duration := range durations {
		if duration < min {
			min = duration
		}
	}
	return min
}

func GetAverageDuration(durations []time.Duration) time.Duration {
	avg := int64(durations[0] / time.Millisecond)
	for i := 1; i < len(durations); i++ {
		avg = avg + int64(durations[i]/time.Millisecond)
	}
	return time.Duration(avg/int64(len(durations))) * time.Millisecond
}
