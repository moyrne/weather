package weather

import (
	"reflect"
	"testing"
	"time"
)

func TestWeComCn_Get(t *testing.T) {
	type args struct {
		cityName string
	}
	ti, err := time.Parse("200601021504", "202107060800")
	if err != nil {
		t.Error(err)
		return
	}
	tests := []struct {
		name     string
		args     args
		wantData Data
		wantErr  bool
	}{
		{
			name: "1",
			args: args{
				cityName: "深圳",
			},
			wantData: Data{
				City:         "深圳",
				Temperature:  "33",
				TemperatureN: "27",
				Weather:      "阴转阵雨",
				Wd:           "东风转东南风",
				Ws:           "<3级转3-4级",
				Time:         ti,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WeComCn{}
			gotData, err := w.Get(tt.args.cityName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Get() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestWeComCn_Parse(t *testing.T) {
	type args struct {
		body []byte
	}
	ti, err := time.Parse("200601021504", "202107060800")
	if err != nil {
		t.Error(err)
		return
	}
	tests := []struct {
		name     string
		args     args
		wantData Data
		wantErr  bool
	}{
		{
			name: "1",
			args: args{
				body: []byte(`var cityDZ ={"weatherinfo":{"city":"深圳","cityname":"shenzhen","temp":"32","tempn":"27","weather":"阵雨","wd":"微风转东南风","ws":"<3级转3-4级","weathercode":"d3","weathercoden":"n3","fctime":"202107060800"}};var alarmDZ ={"w":[{"w1":"广东省","w2":"深圳市","w3":"","w4":"09","w5":"雷电","w6":"02","w7":"黄色","w8":"2021-07-06 15:27","w9":"深圳市气象台于07月06日15时26分发布深圳市雷电黄色预警信号，请注意防御。（预警信息来源：国家预警信息发布中心）","w10":"202107061527594932雷电黄色","w11":"1012806-20210706152700-0902.html","w12":"2021-07-06 15:29","w13":"广东省深圳市发布雷电黄色预警"}]};var dataSK ={"nameen":"shenzhen","cityname":"深圳","city":"101280601","temp":"31","tempf":"87","WD":"西北风","wde":"NW","WS":"1级","wse":"3km\/h","SD":"68%","sd":"68%","qy":"998","njd":"30km","time":"16:00","rain":"0","rain24h":"0","aqi":"14","aqi_pm25":"14","weather":"阴","weathere":"Overcast","weathercode":"d02","limitnumber":"","date":"07月06日(星期二)"};var dataZS ={"zs":{"date":"2021070611","ct_name":"穿衣指数","ct_hint":"炎热","ct_des_s":"建议穿短衫、短裤等清凉夏季服装。","lk_name":"路况指数","lk_hint":"潮湿","lk_des_s":"有降水，路面潮湿，请小心驾驶。","dy_name":"钓鱼指数","dy_hint":"不宜","dy_des_s":"天气不好，有风，不适合垂钓。","cl_name":"晨练指数","cl_hint":"不宜","cl_des_s":"有降水，请尽量避免户外晨练。","nl_name":"夜生活指数","nl_hint":"较适宜","nl_des_s":"只要您稍作准备就可以放心外出。","uv_name":"紫外线强度指数","uv_hint":"最弱","uv_des_s":"辐射弱，涂擦SPF8-12防晒护肤品。","gm_name":"感冒指数","gm_hint":"少发","gm_des_s":"感冒机率较低，避免长期处于空调屋中。","gj_name":"逛街指数","gj_hint":"较不宜","gj_des_s":"有降水，较不适宜逛街","pl_name":"空气污染扩散条件指数","pl_hint":"良","pl_des_s":"气象条件有利于空气污染物扩散。","tr_name":"旅游指数","tr_hint":"适宜","tr_des_s":"有降水，稍热，出游请注意携带雨具。","co_name":"舒适度指数","co_hint":"较不舒适","co_des_s":"白天有雨，气温较高，闷热。","pj_name":"啤酒指数","pj_hint":"适宜","pj_des_s":"天气炎热，可适量饮用啤酒，不要过量。","hc_name":"划船指数","hc_hint":"不适宜","hc_des_s":"天气不好，建议选择别的娱乐方式。","gl_name":"太阳镜指数","gl_hint":"不需要","gl_des_s":"白天能见度差不需要佩戴太阳镜","wc_name":"风寒指数","wc_hint":"无","wc_des_s":"温度未达到风寒所需的低温，稍作防寒准备即可。","pk_name":"放风筝指数","pk_hint":"不宜","pk_des_s":"天气不好，不适宜放风筝。","ac_name":"空调开启指数","ac_hint":"部分时间开启","ac_des_s":"午后天气炎热可适时开启制冷空调。","ls_name":"晾晒指数","ls_hint":"不太适宜","ls_des_s":"降水可能会淋湿衣物，不适宜晾晒。","xc_name":"洗车指数","xc_hint":"不宜","xc_des_s":"有雨，雨水和泥水会弄脏爱车。","xq_name":"心情指数","xq_hint":"较差","xq_des_s":"雨水带来一丝清凉，让烦躁的心绪降温。","zs_name":"中暑指数","zs_hint":"极易发","zs_des_s":"酷热难耐的夏天，开空调不要太贪凉，注意用电安全。","jt_name":"交通指数","jt_hint":"一般","jt_des_s":"有降水且路面湿滑，注意保持车距。","yh_name":"约会指数","yh_hint":"较不适宜","yh_des_s":"建议尽量不要去室外约会。","yd_name":"运动指数","yd_hint":"较不宜","yd_des_s":"有降水，推荐您在室内进行休闲运动。","ag_name":"过敏指数","ag_hint":"不易发","ag_des_s":"除特殊体质，无需担心过敏问题。","mf_name":"美发指数","mf_hint":"一般","mf_des_s":"天热，头皮皮脂分泌多，注意清洁。","ys_name":"雨伞指数","ys_hint":"带伞","ys_des_s":"有降水，短时间出行不必带伞。","fs_name":"防晒指数","fs_hint":"弱","fs_des_s":"涂抹8-12SPF防晒护肤品。","pp_name":"化妆指数","pp_hint":"防脱水","pp_des_s":"请选用防脱水化妆品。","gz_name":"干燥指数","gz_hint":"干燥","gz_des_s":"空气湿度低，易引起皮肤干燥，建议及时使用补水型护肤品，多喝水保持身体滋润。"},"cn":"深圳"};var fc ={"f":[{"fa":"03","fb":"03","fc":"32","fd":"27","fe":"无持续风向","ff":"东南风","fg":"<3级","fh":"3-4级","fk":"0","fl":"3","fm":"99","fn":"60.4","fi":"7\/6","fj":"今天"},{"fa":"03","fb":"03","fc":"32","fd":"28","fe":"东南风","ff":"东南风","fg":"3-4级","fh":"3-4级","fk":"3","fl":"3","fm":"90.1","fn":"71.5","fi":"7\/7","fj":"星期三"},{"fa":"03","fb":"01","fc":"33","fd":"28","fe":"无持续风向","ff":"无持续风向","fg":"<3级","fh":"<3级","fk":"0","fl":"0","fm":"95.4","fn":"66.9","fi":"7\/8","fj":"星期四"},{"fa":"00","fb":"00","fc":"33","fd":"28","fe":"无持续风向","ff":"无持续风向","fg":"<3级","fh":"<3级","fk":"0","fl":"0","fm":"95.7","fn":"68.7","fi":"7\/9","fj":"星期五"},{"fa":"00","fb":"00","fc":"33","fd":"28","fe":"无持续风向","ff":"无持续风向","fg":"<3级","fh":"<3级","fk":"0","fl":"0","fm":"88.9","fn":"65.9","fi":"7\/10","fj":"星期六"}]}`),
			},
			wantData: Data{
				City:         "深圳",
				Temperature:  "32",
				TemperatureN: "27",
				Weather:      "阵雨",
				Wd:           "微风转东南风",
				Ws:           "<3级转3-4级",
				Time:         ti,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WeComCn{}
			gotData, err := w.Parse(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Parse() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
