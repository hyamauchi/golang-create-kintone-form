package main

import (
	"./elements"
	"./output"
	"encoding/json"
	"flag"
	"github.com/koron/go-dproxy"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Configure struct {
	domain    string
	appId     string
	authToken string
}

var config Configure

func main() {
	config.domain = os.Getenv("KINTONE_DOMAIN")
	config.appId = os.Getenv("KINTONE_APP_ID")
	config.authToken = os.Getenv("KINTONE_AUTH_TOKEN")

	flag.StringVar(&config.domain, "domain", config.domain, "Domain name")
	flag.StringVar(&config.appId, "appId", config.appId, "appId")
	flag.StringVar(&config.authToken, "authToken", config.authToken, "Auth token")

	flag.Parse()

	url := "https://" + config.domain + ".cybozu.com/k/v1/app/form/fields.json?app=" + config.appId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Cybozu-Authorization", config.authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Responseの内容を使用して後続処理を行う
	execute(resp)
}

func execute(resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(b))

	var v interface{}
	json.Unmarshal([]byte(string(b)), &v)

	p := dproxy.New(v)
	//s, _ := p.M("revision").String()
	//fmt.Println(s)
	html := `<form class="kintone" role="form" method="post" enctype="multipart/form-data">`

	var d dproxy.Drain
	props := d.Map(p.M("properties"))
	for _, prop := range props {
		//fmt.Println(k, prop)

		item := dproxy.New(prop)

		s, _ := item.M("type").String()
		required, _ := item.M("required").Bool()

		input := make(map[string]string)
		input["code"], _ = item.M("code").String()
		input["label"], _ = item.M("label").String()
		input["defaultValue"], _ = item.M("defaultValue").String()
		input["required"] = ""
		if required {
			input["required"] = "required"
		}

		var slice []string
		switch s {
		case "CHECK_BOX", "DROP_DOWN", "RADIO_BUTTON":
			options, _ := item.M("options").Map()
			slice = make([]string, len(options))
			for _, option := range options {
				item := dproxy.New(option)

				index, _ := item.M("index").String()
				pos, _ := strconv.Atoi(index)

				slice[pos], _ = item.M("label").String()
			}
		}

		switch s {
		case "CHECK_BOX":
			html += "\n" + elements.CreateElementCheckbox(input, slice)
		case "DATE":
			html += "\n" + elements.CreateElementText(input, s)
		case "DATETIME":
			html += "\n" + elements.CreateElementText(input, s)
		case "DROP_DOWN":
			html += "\n" + elements.CreateElementMultiSelect(input, s, slice)
		case "MULTI_LINE_TEXT":
			html += "\n" + elements.CreateElementTextarea(input)
		case "MULTI_SELECT":
			html += "\n" + elements.CreateElementMultiSelect(input, s, slice)
		case "NUMBER":
			html += "\n" + elements.CreateElementText(input, s)
		case "RADIO_BUTTON":
			html += "\n" + elements.CreateElementRadio(input, slice)
		case "SINGLE_LINE_TEXT":
			html += "\n" + elements.CreateElementText(input, s)
		default:
		}

	}

	// HTMLファイル出力
	params := make(map[string]string)
	params["body"] = html + `
<p><button type="submit" class="btn btn-primary btn-large">送信</button></p>
</form>`
	output.OutputToFile(params, output.HtmlTemplate(), "./output.html")

	// Lambda Function出力
	params["domain"] = config.domain
	params["appId"] = config.appId
	params["apiToken"] = "【 apiToken 】"
	output.OutputToFile(params, output.LambdaTemplate(), "./handler.js")

	// serverless.yml出力
	params["servicename"] = "my-kintone-form"
	params["region"] = "ap-northeast-1"
	params["stage"] = "dev"
	output.OutputToFile(params, output.ServerlessTemplate(), "./serverless.yml")
}
