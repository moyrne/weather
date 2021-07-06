package weather

import "time"

type Data struct {
	City         string    `json:"city"`          // 城市
	Temperature  string    `json:"temperature"`   // 最低气温
	TemperatureN string    `json:"temperature_n"` // 最高气温
	Weather      string    `json:"weather"`       // 天气
	Wd           string    `json:"wd"`            // 风向
	Ws           string    `json:"ws"`            // 风速
	Time         time.Time `json:"time"`          // 时间
}

type Weather interface {
	Get(cityName string) (Data, error)
}
