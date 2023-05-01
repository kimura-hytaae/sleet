# sleet
  sleetはターミナル上で天気を見ることができるCLIです。

[![build](https://github.com/kimura-hytaae/sleet/actions/workflows/build.yaml/badge.svg)](https://github.com/kimura-hytaae/sleet/actions/workflows/build.yaml)
  [![Coverage Status](https://coveralls.io/repos/github/kimura-hytaae/sleet/badge.svg?branch=main)](https://coveralls.io/github/kimura-hytaae/sleet?branch=main)
  [![codebeat badge](https://codebeat.co/badges/de48853f-d89b-4ac2-bce7-3697e76e052e)](https://codebeat.co/projects/github-com-kimura-hytaae-sleet-main)
  [![Go Report Card](https://goreportcard.com/badge/github.com/kimura-hytaae/sleet)](https://goreportcard.com/report/github.com/kimura-hytaae/sleet)

![Version](https://img.shields.io/badge/Version-0.1.0-informational)
 
# 概要（Overview）
ターミナル上で実行することで天気を瞬時に見ることができます。
 
# 使い方（Usage）
CLIで天気予報を取得します。 各地域の天気と洗濯物の乾きやすいかどうか示す乾燥率も表示する予定です。 

```bash
sleet [OPTIONS] [LOCATIONs...] [DAYs...] [WEATHERs]
OPTIONS
  -v, --version                バージョンを表示します。
  -h, --help                   ヘルプを表示します。
  -
LOCATION
  specify the location in the following ways. 次の方法などで指定します。
  - 緯度経度
  - ポスタルコード
  - 市の名前
DAY
  specify the location in the following ways. 次の方法などで指定します。
  - 日付指定
  - 何日かの指定
```
 
# インストール方法（Installation）
 
 まだ書けていない  
Requirementで列挙したライブラリなどのインストール方法を説明する
 
```bash
pip install sleet_package
```

 
# プロジェクトについて
 
作成情報を列挙する
 
* 開発者：　　　　木村颯
* 所属：　　　　　京都産業大学 先端情報学研究科
* メールアドレス：i2386058@cc.kyoto-su.ac.jp
* 名前の由来：　　sleetは霙(みぞれ)の意味で天気を意識したものです。
* バージョン：　　2023/4/18 ver.1
* アイコン：　　　
 
 ![sleet_icon](https://user-images.githubusercontent.com/92291361/232662040-3a87611a-bf53-41bb-8a86-b6954be1eb1b.svg)

 
# License 
![](https://img.shields.io/github/license/kimura-hytaae/sleet)

"sleet"はMITのライセンスを使用しています。
"sleet" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
 
 "sleet"はオープンソースです。
"sleet" is Open sourse.
