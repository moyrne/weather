package weather

import "time"

type Data struct {
	City         string    // 城市
	Temperature  string    // 最低气温
	TemperatureN string    // 最高气温
	Weather      string    // 天气
	Wd           string    // 风向
	Ws           string    // 风速
	Time         time.Time // 时间
}

type Weather interface {
	Get(cityName string) (Data, error)
}
