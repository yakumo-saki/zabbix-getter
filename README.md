# zabbix-getter

zabbix-senderの逆でzabbixから値を取得します。

## インストール

Windows: exeファイルを適当な場所においてください。
Arch Linux： AURからインストール可能です。 `yay -S zabbix-getter`
それ以外： 実行ファイルを展開して `chmod +x zabbix-getter` してください

## 初期設定

以下の内容の設定ファイルを以下の場所に作成します。
設定値は適切に変更してください。
複数の場所に設定ファイルを作成した場合、ルールに従って上書きされますが、
あまり意味がないのでオススメはしません。

* ~/.config/zabbix-getter.conf (Linux/macOS)
* %APPDATA%\zabbix-getter.conf (Windows)
* $HOME/Library/Preferences/zabbix-getter.conf (macOS)
* 実行ファイルと同ディレクトリのzabbix-getter.conf

About UserConfigDir
https://github.com/golang/go/issues/29960

### configration file example

```
USERNAME=Admin
PASSWORD=zabbix
ENDPOINT=http://192.168.1.123/api_jsonrpc.php
OUTPUT=JSON
LOGLEVEL=WARN
```

※ 詳細は設定についてをご覧ください。

## 実行例

実行結果と致命的なメッセージのみ標準出力(stdout)に出力されます。
それ以外は標準エラー出力(stderr)に出力されます。

`$ zabbix-getter -s test -k testitem`
```
{
  "itemId": "47991",
  "hostId": "10322",
  "key_": "testitem",
  "name": "test item name #1",
  "value": "1234",
  "lastClock": "0",
  "units": "test_units"
}
```

`$ zabbix-getter -s test -k testitem -o VALUE`
1234

## 修了コード

* 0 正常
* 2 ヘルプ (-h) が指定された
* 10以上 異常修了

基本的には 0 以外はすべて異常終了です。

## 設定

1. {CONFIG_DIR}/zabbix-getter.conf
2. {EXEC_DIR}/zabbix-getter.conf
3. 環境変数
4. コマンドラインオプション

上記の順で後勝ちで設定を行えます。

| オプション        | 環境変数   | デフォルト | 設定内容     | サンプル                              |
| --------------- | --------- | ---------| ------------ | ------------------------------------ |
| -h              | (なし)     | -        | ヘルプメッセージを出力                |  |
| （なし）          | USERNAME  | ""      | zabbixユーザー名。環境変数か設定ファイルで指定してください。 | Admin |
| （なし）          | PASSWORD  | ""      | zabbixパスワード。環境変数か設定ファイルで指定してください。 | zabbix |
| -e , --endpoint | ENDPOINT  | ""       | zabbix APIエンドポイント           | http://192.168.1.100/api_jsonrpc.php |
| -s , --hostname | (なし)     | ""       | zabbixに登録されたホスト名（キーの方） | testhost |
| -k , --key      | (なし)     | ""       | ホストアイテムのキー                 | system.hostname |
| -l , --loglevel | LOGLEVEL  | WARN     | ログ出力レベル TRACE>DEBUG>INFO>WARN>ERROR>FATAL | (CLI) -loglevel TRACE |
| -o , --output   | OUTPUT   | VALUE     | 出力を [JSON | VALUE] にする。VALUEは値のみ出力 | (CLI) -output JSON |
| -z , -Z   | （なし）   | -     | -e http(s)://<値>/api_jsonrpc.php のショートハンド | (CLI) -z 10.20.30.40 |
| --debug   | （なし）   | -     | デバッグ用。値の指定チェックを行わない |  |

### --debug について

* 通常は `-s` `-k` による指定が必須ですが、このチェックを行わないようになります。
* 要するに、zabbix-getter.conf に書いてある HOSTNAME と KEY が使用されます。
* 自分の環境の HOSTNAMEやKEYをコード上(launch.json)に露出させたくないために作成された機能です。
