package enums

type AlertFrequency string

const (
	AlertFrequency_ONE_TIME AlertFrequency = "one_time"
	AlertFrequency_15m                     = "15m"
	AlertFrequency_30m                     = "30m"
	AlertFrequency_1hr                     = "1hr"
	AlertFrequency_6hr                     = "6hr"
	AlertFrequency_12hr                    = "12hr"
	AlertFrequency_24hr                    = "24hr"
)
