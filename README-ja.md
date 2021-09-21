# zabbix-getter

zabbix-senderの逆でzabbixから値を取得します。

## 実行例

zabbix-getter -e http://192.168.1.100/api_jsonrpc.php -s test -k testitem

## 設定

優先順位は、環境変数＞コマンドラインオプション （後勝ちで上書き）

| コマンドラインオプション | 環境変数 | 設定内容     | サンプル                              |
| ------- | --------- | --------------------- | ------------------------------------ |
| （なし） | USERNAME  | zabbixユーザー名                  | Admin |
| （なし） | PASSWORD  | zabbixパスワード                  | zabbix |
| -e     | ENDPOINT  | zabbix APIエンドポイント           | http://192.168.1.100/api_jsonrpc.php |
| -s     | (なし)     | zabbixに登録されたホスト名（キーの方） | testhost |
| -k     | (なし)     | ホストアイテムのキー                 | system.hostname |
| -v     | LOGLEVEL  | （未実装）ログ出力レベル 冗長 1 TRACE << 5 FATAL 簡潔 | 1 |
| -h     | (なし)     | ヘルプメッセージを出力                |  |
