# zabbix-getter

zabbix-senderの逆でzabbixから値を取得します。

## 実行例

$ zabbix-getter -e http://192.168.1.100/api_jsonrpc.php -s test -k testitem
123

$ zabbix-getter -json -e http://192.168.1.100/api_jsonrpc.php -s test -k testitem
{}

## 設定

優先順位は、環境変数＞コマンドラインオプション （後勝ちで上書き）

| オプション   | 環境変数   | デフォルト | 設定内容     | サンプル                              |
| ---------- | --------- | ---------| ------------ | ------------------------------------ |
| （なし）     | USERNAME  | ""      | zabbixユーザー名                  | Admin |
| （なし）     | PASSWORD  | ""      | zabbixパスワード                  | zabbix |
| -e         | ENDPOINT  | ""       | zabbix APIエンドポイント           | http://192.168.1.100/api_jsonrpc.php |
| -s         | (なし)     | ""       | zabbixに登録されたホスト名（キーの方） | testhost |
| -k         | (なし)     | ""       | ホストアイテムのキー                 | system.hostname |
| -loglevel  | LOG_LEVEL | WARN     | ログ出力レベル TRACE<DEBUG<INFO<WARN<ERROR<FATAL | (CLI) -loglevel TRACE |
| -h         | (なし)     | -        | ヘルプメッセージを出力                |  |
| -output     | OUTPUT  | VALUE     | 出力を [VALUE | JSON] にする。VALUEは値のみ出力 | (CLI) -output JSON |
