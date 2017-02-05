package output

import (
	"../common"
	"os"
)

func OutputToFile(params map[string]string, t string, filename string) {
	out := common.EvalTemplate(t, params)

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(([]byte)(out))
}

func LambdaTemplate() string {
	return `var https = require('https');

exports.handler = function (event, context) {
    console.log(event);
    var body = event.body;

    var param = { "app": {{.appId}}, "record": {} };
    Object.keys(body).forEach(function (key) {
        if (/^_kintone_control_/.test(key)) {
            param.record[key.replace('_kintone_control_', '')] = { "value": body[key] };
        }
    });

    var json = JSON.stringify(param);
    console.log(json);

    var options = {
        hostname: '{{.domain}}.cybozu.com',
        port: 443,
        path: '/k/v1/record.json',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'X-Cybozu-API-Token': '{{.apiToken}}'
        }
    };
    var req = https.request(options, function (res) {
        console.log('STATUS: ' + res.statusCode);
        console.log('HEADERS: ' + JSON.stringify(res.headers));
        res.setEncoding('utf8');
        res.on('data', function (chunk) {
            console.log('BODY: ' + chunk);
            if (res.statusCode === 200) {
                context.succeed(chunk);
            }
        });

    });
    req.on('error', function (e) {
        console.log('problem with request: ' + e.message);
        context.fail(e.message);
    });
    req.write(json);
    req.end();
};
`
}

func ServerlessTemplate() string {
	return `service: {{.servicename}}
provider:
  name: aws
  runtime: nodejs4.3
  stage: {{.stage}}
  region: {{.region}}
functions:
  handler:
    handler: handler.handler
    events:
      - http:
          path: handler
          method: post
          cors: true
`
}

func HtmlTemplate() string {
	return `<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="UTF-8"/>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>kintoneへ登録</title>
  <link rel="stylesheet" type="text/css" href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"/>
  <link rel="stylesheet" type="text/css" href="//maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"/>
  <script src="//ajax.googleapis.com/ajax/libs/jquery/2.2.4/jquery.min.js"></script>
  <link rel="stylesheet" href="static/jquery.datetimepicker.css">
  <link rel="stylesheet" href="static/wp-to-kintone-form.css">
  <link rel="stylesheet" href="static/style.css">
  <script src="static/jquery.datetimepicker.js"></script>
  <script src="static/wp-to-kintone-form.js"></script>
</head>
<body>
<div>
  <div class="container">
{{.body}}
  </div>
</div>
<p id="page-top"><a href="#top"><span class="glyphicon glyphicon-chevron-up"></span></a></p>
</body>
<script>
  jQuery.noConflict();
  (function ($) {
    $(function () {
      $('form.kintone').submit(function () {
        var param = {};
        $($("form.kintone").serializeArray()).each(function (i, v) {
          param[v.name] = v.value;
        });
        $.ajax({
          url: "【 AWS API Gateway Endpoint url 】",
          type: "POST",
          timeout: 20000,
          data: JSON.stringify(param),
          contentType: 'application/json',
          dataType: 'json'
        }).done(function (response) {
          alert("送信しました。");
        }).fail(function () {
          alert("送信できません。");
        });
        return false;
      });

      var topBtn = $('#page-top');
      topBtn.hide();
      //スクロールが100に達したらボタン表示
      $(window).scroll(function () {
        if ($(this).scrollTop() > 100) {
          topBtn.fadeIn();
        } else {
          topBtn.fadeOut();
        }
      });
      //スクロールしてトップ
      topBtn.click(function () {
        $('body,html').animate({
          scrollTop: 0
        }, 500);
        return false;
      });
    });
  })(jQuery);
</script>
</html>`
}
