# Twitterクローン

技術課題であるTwitterクローンを作成し提出

## Overview

クリーンアーキテクチャに則って作成  
![クリーンアーキテクチャ](./CleanArchitecture.jpg)

### Setup

ローカルでのセットアップに関しては、このdockerを使用し下記コマンドを実施する
```
// front側とback側のgitをcloneしてくる
git clone git@github.com:Kaminashi-Inc/ENG-1009_bukubukutyagamaguti.git ./dev/back

cp ./dev/back/.env.example ./dev/back/.env

git clone git@github.com:bukubukutyagamaguti/vue_front.git
cp ./dev/front/.env.example ./dev/front/.env

make init
```
これでローカルで開発環境が立ち上がる。  
アクセスは、[http://127.0.0.6:8000](http://127.0.0.6:8000)にて行う  
フロントサイドの開発は、[http://127.0.0.6:8081](http://127.0.0.6:8081)にて描写  
サーバーサイドの開発は、[http://127.0.0.6:8080](http://127.0.0.6:8080)にて描写

### Test

テストコードはinterface層とinfra層での導入ができていないので今後対応予定

## Production

基本機能の流れ
1. Loginする
- EmailとPasswordを使用する
- post通信にて対応  

2. Token発行(JWT)
- Login時に発行する
- Tokenには、UserIDとUserNameが入っている

3. Tokenを使用してuser情報等を取得
- ツイートやフォロー等の処理でIDを使用する

## Notes

その他

## 工夫した点

今回のGoの実装に関して工夫した点を記載しています。 

### 単一責任

データベースへの処理  
トークンへの処理  
などのミドルウェアごとに処理をまとめるように行いできるだけ  
コードの拡張の際に処理を行うディレクトリに迷わないように行いました。  

### ディレクトリ構成

クリーンアーキテクチャに則って作成  
参考記事[こちら](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)

#### Interfaces

このディレクトリでは外部との通信するための処理を実装

##### Controllers

interfacesで受けたリクエストを受け取ります。  
リクエストを受け取ったあとにinteractorになげるための初期化も行う

#### Database 

interfacesで受けたリクエストをもとに処理をinteractorに流す際  
にデータベースへの処理をinfrastructuresつなげる役割  

#### Infrastructures

このディレクトリではアプリケーション外部へ通信するための処理を実装

##### Routes
ルーティングを行います。

##### Database
データベースへのアクセスするための初期化設定

#### UseCases

domainやdatabaseを呼ぶ処理を実装  
所謂サービスと似たような役割にて考慮してあえて命名によりにファイルを分けています。  
すべてのUsecaseを呼ぶのではなく、  
必要なUsecaseのみ呼べるようにしたいからこのように作成しています。  

#### Domain

ドメインロジックを実装

## 苦労した点

今回は、基本初めてのことばかりでした。  
クリーンアーキテクチャを使用したディレクトリ構成  
Golangやフレームワークのechoの使用  
JWTを使用したトークン発行
vueを使用したSPA  
などなどの初めてのことが多い実装でした。  

ただ改めてこのような実装をしていくことによって自分の知識が増えていく楽しさや実装していく楽しさを感じることができました。  
正直まだ実装したいことややりたいことも多いので今後もこの開発を行っていき自分の知識のアップデートをしていこうと思います。

今回は間に合わないと思う実装は今後やる予定です。  
全文検索エンジンの開発  
認証認可周りの実装  
JWT以外のToken発行のスキーム  
controller、database、token周りのテストコード

## 未実装の点

軽め
投稿の削除
ユーザのプロフィール
ユーザの設定
ユーザ一覧

重め
投稿またはユーザの検索
認証認可周りの実装
API用のToken発行システム
interface、infra層でのテストコード作成
