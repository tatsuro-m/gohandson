// STEP03: パッケージを分けてみよう

// mainパッケージの定義
package main

import "github.com/tenntenn/gohandson/greeting"

// main関数から実行される
func main() {
	greeting.Do()
}
