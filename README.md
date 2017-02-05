# golang-create-kintone-form

## これはなに？ / What's this?

WordPress向けの [cybozu-wp-to-kintoneプラグイン](https://ja.wordpress.org/plugins/cybozu-wp-to-kintone/)
が生成するものと同様のHTMLフォームを生成します。

また、[Serverless Framework](https://serverless.com/)を用いて
kintoneへPOSTするAPIの雛型を作成します。

これらを用いることにより、Webサイトから簡単にkintoneへデータを入れることができるようになります。

## 必要なもの / Requirements

*  AWS CLI
* [Serverless Framework](https://serverless.com/)
*  kintoneのアカウント情報とアプリ、レコード追加権限を付与したAPIトークン

## 使い方 / Usage

* ドメイン、対象の kintoneアプリID、認証情報を指定して実行します。
````
golang-create-kintone-form.exe -domain {{ CYBOZU_DOMAIN }} -appId {{ 対象のアプリID }} -authToken {{ 「ログイン名:パスワード」をBASE64エンコードしたもの }}
````

* 以下のファイルが出力されます。

| 出力されるファイル | 用途 |
|:-----------------|:------------|
| output.html      | HTMLフォーム |
| serverless.yml   | Serverless Framework用の設定ファイル |
| hander.js        | Lambda Function |


* 空のフォルダを用意して、そのフォルダへ serverless.ymlとhander.jsをコピーします。


* Lambda Function内でAPIトークンを指定しているので、対象の kintoneアプリ用に生成した APIトークンに置き換えます。
````
            'X-Cybozu-API-Token': '【 apiToken 】'
````
````
            'X-Cybozu-API-Token': 'xXXxxX0XXX8x9Xx3X1XX985XxxXxXxxX17XxxxX5'
````

* Serverless Frameworkを用いて APIを deployします。
````
sls deploy
````

* deployが正常に完了するとAPIのエンドポイントが返されるのでURL文字列をコピーします。
````
region: ap-northeast-1
api keys:
  None
endpoints:
  POST - https://xxxxxxxxxx.execute-api.ap-northeast-1.amazonaws.com/dev/handler
functions:
  my-kintone-form-dev-handler: arn:aws:lambda:ap-northeast-1:xxxxxxxxxxxx:function:my-kintone-form-dev-handler
````

* HTMLフォーム内の APIエンドポイント指定箇所を書き換えます。
````
          url: "【 AWS API Gateway Endpoint url 】",
````
````
          url: "https://xxxxxxxxxx.execute-api.ap-northeast-1.amazonaws.com/dev/handler",
````

* （オプション）項目の並びを調整し、CSSで整えます。（exampleフォルダにCSSの例があります。）


## 注意点

* APIは保護されていません。Access-Control-Allow-Origin: '*' になっています。HTMLフォームを実際に運用するのであれば、適正なリクエストのみを受け付けるようにAPIを設定してください。そうしなければ、APIがDoS攻撃を受ける恐れがあります。
* Lambda Function内に kintoneのアプリID、APIトークン、CYBOZU_DOMAINが平文で書かれています。KMSを用いるように変更してください。

## License

[Apache v2 License](http://www.apache.org/licenses/LICENSE-2.0.html)
