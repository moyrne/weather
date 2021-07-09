package weather

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// cityId 城市ID
// msTimeUnix 时间戳
// http://d1.weather.com.cn/weather_index/${cityId}.html?_=${msTimeUnix}

type WeComCn struct{}

var ErrCityNotFound = errors.New("city name not found")

func (w WeComCn) Get(cityName string) (data Data, err error) {
	cityCode, err := GetCityID(cityName)
	if err != nil {
		return data, err
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

	return w.Parse(all)
}

var weComCnRegex = regexp.MustCompile(`var[^=]*=([^\{]*\{[^;]*);`)
var ErrBodyMatchFailed = errors.New("regexp sub match failed")

func (w WeComCn) Parse(body []byte) (data Data, err error) {
	result := weComCnRegex.FindAllSubmatch(body, 1)
	if len(result) == 0 {
		return data, ErrBodyMatchFailed
	}
	var value map[string]map[string]string
	if err := json.Unmarshal(result[0][1], &value); err != nil {
		return data, err
	}
	info := value["weatherinfo"]
	t, err := time.Parse("200601021504", info["fctime"])
	if err != nil {
		return data, err
	}
	data = Data{
		City:         info["city"],
		Temperature:  info["temp"],
		TemperatureN: info["tempn"],
		Weather:      info["weather"],
		Wd:           info["wd"],
		Ws:           info["ws"],
		Time:         t,
	}
	return data, nil
}
