# Twitterクローン

技術課題であるTwitterクローンを作成し提出

## ディレクトリ構成

クリーンアーキテクチャに則って作成
参考記事[こちら](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)

### Adapters

このディレクトリでは外部との通信するための処理を実装

#### Controllers

Infrastructuresで受けたリクエストを受け取ります。
リクエストを受け取ったあとにinteractorになげるための初期化も行う

#### Gateways

外部ツールとつなぐためのinterface
基本的には、interfaceの定義とinterfaceを介して外部のメソッドを実行する処理を定義します。

#### Presenter

外部にアウトプットする際にデータをフォーマットしています

### Infrastructures

このディレクトリではアプリケーション外部へ通信するための処理を実装

#### routes
ルーティングを行います。

#### database
データベースへのアクセスするための初期化設定

### Usecases
EntitiesやGatewaysを呼ぶ処理を実装
所謂サービスと似たような役割にて考慮
あえてEntity毎にディレクトリを分けています。
すべてのUsecaseを呼ぶのではなく、必要なUsecaseのみ呼べるようにしたいからこのように作成しています。

### Entities
ドメインロジックを実装

## 工夫した点

## 苦労した点

## 未実装の点