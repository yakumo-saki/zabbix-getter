package zabbix

type ZabbixError struct {
	Msg string
	Err error
}

// fmt.Println など、出力するための関数は、内部的に error型だった場合に Error を使う
// 行番号など、エラーメッセージに付け足すための処理を書いておく
func (e *ZabbixError) Error() string {
	return e.Msg
}

// Unwrap を定義しておけば、errors.Unwrap が使える
func (e *ZabbixError) Unwrap() error {
	return e.Err
}
