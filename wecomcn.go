package weather

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

// cityId 城市ID
// msTimeUnix 时间戳
// http://d1.weather.com.cn/weather_index/${cityId}.html?_=${msTimeUnix}

type WeComCn struct{}

var ErrCityNotFound = errors.New("city name not found")

func (w WeComCn) Get(cityName string) (data Data, err error) {
	cityCode, ok := weComCnCities[cityName]
	if !ok {
		return data, errors.WithMessage(ErrCityNotFound, cityName)
	}

	resp, err := http.DefaultClient.Get("http://d1.weather.com.cn/weather_index/" + cityCode + ".html?_=" + strconv.Itoa(int(time.Now().UnixNano()/1000000)))
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	fmt.Println(string(all))

	return data, nil
}

func (w WeComCn) Parse(body []byte) (data Data, err error) {
	return data, nil
}
