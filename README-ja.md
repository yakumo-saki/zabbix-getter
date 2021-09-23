# zabbix-getter

zabbix-senderの逆でzabbixから値を取得します。

## インストール

Windows: exeファイルを適当な場所においてください。
それ以外： 実行ファイルを展開して `chmod +x zabbix-getter` してください

## 初期設定

以下の内容の設定ファイルを以下の **どちらか** の場所に作成します。
設定値は適切に変更してください。

* ~/.config/zabbix-getter.conf (Linux)
* %APPDATA%\zabbix-getter.conf (Windows)
* $HOME/Library/Preferences/zabbix-getter.conf (macOS)
* 実行ファイルと同ディレクトリのzabbix-getter.conf

About UserConfigDir
https://github.com/golang/go/issues/29960

```
USERNAME=Admin
PASSWORD=zabbix
ENDPOINT=http://192.168.1.123/api_jsonrpc.php
OUTPUT=JSON
LOGLEVEL=WARN
```

## 実行例

$ zabbix-getter -e http://192.168.1.100/api_jsonrpc.php -s test -k testitem
123

$ zabbix-getter -json -e http://192.168.1.100/api_jsonrpc.php -s test -k testitem
{}

## 修了コード

* 0 正常
* 2 ヘルプ (-h) が指定された
* 10以上 異常修了

基本的には 0 以外はすべて異常終了です。

## 設定

優先順位は、環境変数＜{CONFIG_DIR}/zabbix-getter.conf＜{EXEC_DIR}/zabbix-getter.conf＜＜＜コマンドラインオプション
コマンドラインオプションはすべてを上書きできます。それ以外は上記の順番で　先勝ち（上書きしない）　です


| オプション   | 環境変数   | デフォルト | 設定内容     | サンプル                              |
| ---------- | --------- | ---------| ------------ | ------------------------------------ |
| （なし）     | USERNAME  | ""      | zabbixユーザー名                  | Admin |
| （なし）     | PASSWORD  | ""      | zabbixパスワード                  | zabbix |
| -e         | ENDPOINT  | ""       | zabbix APIエンドポイント           | http://192.168.1.100/api_jsonrpc.php |
| -s         | (なし)     | ""       | zabbixに登録されたホスト名（キーの方） | testhost |
| -k         | (なし)     | ""       | ホストアイテムのキー                 | system.hostname |
| -loglevel  | LOGLEVEL  | WARN     | ログ出力レベル TRACE<DEBUG<INFO<WARN<ERROR<FATAL | (CLI) -loglevel TRACE |
| -h         | (なし)     | -        | ヘルプメッセージを出力                |  |
| -output     | OUTPUT   | VALUE     | 出力を [JSON | VALUE] にする。VALUEは値のみ出力 | (CLI) -output JSON |

