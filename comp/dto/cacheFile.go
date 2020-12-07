package dto

import (
	"os"
	"strconv"
	"time"
)

type CacheFile struct {
	Created string
	Version string
}

const (
	defaultRenewalInterval         = int64(518_400_000) // six days
	customRenewalIntervalEnvVarKey = "CACHE_RENEWAL_INTERVAL"
)

func (c CacheFile) IsOutdated() bool {
	created, err := strconv.ParseInt(c.Created, 10, 0)
	if err != nil {
		return true
	}
	customRenewalInterval := os.Getenv(customRenewalIntervalEnvVarKey)
	renewalInterval := defaultRenewalInterval
	if len(customRenewalInterval) > 0 {
		if renewal, err := strconv.ParseInt(customRenewalInterval, 10, 0); err == nil && renewal > 0 {
			renewalInterval = renewal * 1_000
		}
	}
	diff := time.Now().Unix() - created
	return diff*1_000 > renewalInterval
}
