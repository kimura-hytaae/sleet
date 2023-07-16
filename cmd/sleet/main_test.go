package main

import "testing"

func _Example_Main() {
	goMain([]string{"./sleet", "-help"})
	// Output:
	// [OPTIONS] [LOCATIONs...] [DAYs...] [WEATHERs]
	// OPTIONS
	// -v, --version            print the version and exit.
	// -h, --help               print this mesasge and exit.

	// LOCATION
	// specify the location in the following ways. 次の方法などで指定します。
	// - 緯度経度
	// - ポスタルコード
	// - 市の名前
	// DAY
	// specify the location in the following ways. 次の方法などで指定します。
	// - 日付指定
	// - 何日かの指定
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./sleet", "Osaka"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
