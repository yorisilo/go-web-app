# ネットワークプログラミングの勉強 with go
* [Go Web プログラミング](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/)
* [ネットワークプログラミング](http://ash.jp/net/prog_net.htm)
* [Goでnet/httpな簡単なJSON API](https://muunyblue.github.io/ebeb300882677f350ea818c8f333f5b9.html)
* [GoでJSON APIを書く](http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/)
* [go web examples](https://gowebexamples.com/)

# nc コマンド
[nc コマンド 使い方メモ](https://qiita.com/yasuhiroki/items/d470829ab2e30ee6203f)

# OSI参照モデル

https://www.itbook.info/study/p36.html

| プロトコル階層             | 規格                            |
|:--------------------------:|:-------------------------------:|
| 第7層 アプリケーション層   | HTTP, SMTP, POP3, NFS, DNS, ... |
| 第6層 プレゼンテーション層 | HTTP, SMTP, POP3, NFS, DNS, ... |
| 第5層 セッション層         | Socket                          |
| 第4層 トランスポート層     | TCP/UDP                         |
| 第3層 ネットワーク層       | IP                              |
| 第2層 データリンク層       | イーサネット                    |
| 第1層 物理層               | イーサネット                    |

イーサネット規格は以下にまたがる規格
* 第2層 データリンク層
* 第1層 物理層

# トランスポート層
[TCP/UDP - コネクションとコネクションレス](https://www.itbook.info/study/tcp3.html)

トランスポート層の代表的なプロトコル
* TCP コネクション型
* UDP コネクションレス型

>ネットワーク層は、ノード間の通信を行う機能を提供するために、	IP アドレスを使用しています。
>
>IP アドレスはネットワークアドレスとホストアドレスに分かれていて、	どのネットワークのどのノードにパケットを転送すればよいかが	分かるようになっています。
>
>しかし、どのアプリケーションにデータを渡すのかは IP アドレスでは	知る事が出来ません。
>
>そこでトランスポート層では、「ポート番号」を使用して	どのアプリケーションに渡すべきなのかを判断しています。
>
>例えば 1 台のサーバでWebサーバとメールサーバという	２種類の役割を担っているサーバがあるとしましょう。 このサーバを仮に「SV-1」と呼びます。
>
>ちなみにこの役割のことを「サービス」と呼びまして、	それぞれのサービスをWebサービスやメールサービス	なんて呼びます。

## コネクション型・コネクションレス型
> トランスポート層ではの役割は「アプリケーションレベルでの通信を確立する」
> ことであって、そのためにポート番号を利用して制御をしています。
>
> TCP でも UDP でもそれぞれのヘッダにあて先、送信元のポート番号を
> 付与していて、アプリケーションの制御を実施しています。
>
> コネクション型とコネクションレス型の違いは、
> このアプリケーションの制御方法の違いに関係してきます。

## コネクション型
TCP など

> コネクション型はその名の通り、通信を確立するもの同士で
> 連絡を取り合って制御を行う方式です。
>
> 日常で見られるものに例えるのならば電話がまさにコネクション型。
>
> 電話は互いに受話器を持って話をします。
> 話をした言葉は相手にちゃんと届きますし、もし話している途中で
> うまく聞き取れないときには、相手に聞き返してもう一度同じ内容を
> 聞く事だって出来ます。
>
> このように、通信を確立するためにお互いが連絡を取り合って制御する
> 方式をコネクション型と呼んでいて、信頼性がある方式であるといえます。
>
> コネクション型の代表的なプロトコルには「TCP」があります。

## コネクションレス型
UDPなど

> コネクションレス型は、連絡を取り合って制御はせずに、相手が受け取る
> 準備ができているかにかかわらず、送信してしまう方式です。
>
> 手紙はまさにコネクションレス型です。
> 相手が何をしていようと手紙を送りつけて、基本的にいつ届くのかは
> わかりませんし、もしかしたらどこかで紛失してしまい、
> 相手に届かないこともあるかもしれません。
>
> このようにコネクションレス型は、相手が何をしているかは関係なしに
> 送りつける方式です。
> もちろん途中でパケットが消失してしまったり、
> 到着順も考慮していません。
>
> そのためコネクションレス型通信は信頼性のない通信と言われています。
>
> ただし、メリットもあるわけで、コネクション型はお互いが連絡を取り合う
> 必要があることから、安定した通信を実現できるものの、
> 確認のためのデータのやり取りが発生することから、
> ネットワークのトラフィックが増大してしまうデメリットがありますが、
> 逆にコネクションレス型はネットワークへの負荷が軽いため
> 通信の効率が良いメリットがあります。
>
> コネクションレス型の代表的なプロトコルには「UDP」があります。

# web アクセス概観
> 普段ホームページを閲覧する際、ブラウザを開くと思います。アドレスを入力してエンターキーを押すと、あなたが見たいコンテンツが表示されます。この見た目には簡単なユーザの行動には一体何が隠されているのでしょうか？

1. URLを入力する際まずブラウザはDNSサーバにアクセスします。
1. DNSを通してドメインと対応するIPを取得し、IPアドレスからIPに対応したサーバを探しだした後、
1. TCPコネクションの設立を要求します。ブラウザがHTTP Request（リクエスト）パケットを送信し終わると、サーバはリクエストパケットを受け取ってリクエストパケットを処理しはじめます。
1. サーバは自分のサービスをコールし、HTTP Response（レスポンス）パケットを返します。
1. クライアントがサーバからのレスポンスを受け取ると、このレスポンスパケットのbodyを読み出します。
1. すべての内容を受け取ると、このサーバとのTCP接続を切断します。
https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/03.1.html


> ## 本来HTTPはステートレス
> HTTPプロトコルというのは非常に古いプロトコルである。大昔のWebというのは、サーバにある資源（ファイル、画像、文書）を取ってくるという仕事がメインだった。だから、サーバと、その中にある資源を一意に表すことのできるURLを解釈して、それに応じた資源をクライアントに返すという仕事で事足りていた。
>
> たとえば、http://xxx.com/welcome.png にアクセスしたら必ずwelcome.pngが返ってくる。http://xxx.com/index.html にアクセスしたら必ずindex.htmlが返ってくる。何度アクセスしても、サーバの資源が変更されない限りはその内容は同じものだ。別のクライアントからアクセスしてもそれは同じである。
> ## ショッピングカート
> 過去のWebでは要求されたURLに対応するものを愚直に返却するのがサーバの仕事であったが、現代では同じURLを要求しても違う結果が返ってくることが必要とされる場面が多々ある。代表例が通販サイトのショッピングカートだろう。
>
> たとえば、http://xxxx.com/shopping/cart がショッピングカートのURLだったとする。たとえばユーザーAからこのURLが要求された場合はユーザーAのカートの中身を表示させなければならないし、ユーザーBであったらユーザーBのカートの中身を表示させなければならない。また、ログインしていないユーザーであればログインページを表示するということになるだろう。
>
> こうした必要性が生じたときに昔のNetscape Communicationsという会社が提案した方法はCookieという仕組みであった。これはサーバから送信したデータをクライアント側のブラウザで保存できる仕組みであり、その後ほとんどのブラウザに実装された。Cookieはキーと値のペアからなるシンプルな形式のデータからなり、それを受け取ったクライアントは次回以降、同じサーバにリクエストを発行するときは必ずこのキー／値の組を送信する。
>
> ショッピングカートを実装する仕組みはこうだ。サーバはユーザーが正しいユーザーIDとパスワードを送り、サーバがそれを検証して正しいという判断を下すと、クライアントに対してCookieを発行する。これに含まれるデータはふつう、ある程度の長さをもった乱数である。クライアントはCookieを受け取ったら次回以降のリクエストには必ずこのデータを付与してリクエストを発行するので、この乱数は毎回サーバ側に送信される。

> ## サーバ側処理の隠ぺい
> ショッピングカートの例を再度考えよう。サーバ側がやらなくてはならない処理の一つは、各URLが要求された時にユーザーに応じて別の表示状態を構成しなければならないということであった。先に挙げた例はショッピングカートの表示であったが、ショッピングカートだけで構成されるWebサイトというのは無い。多くの場合、トップページにはその人の過去の購入履歴から判断した嗜好に応じた商品が表示されるし、商品の購入を考えても住所を記入するページ、支払方法を選択するページ、最終確認のページ…とさまざまな遷移がある。それらすべての状態をサーバ側で管理する方法を実装するのはなかなか大変である。
>
> 大変であるが、多くの場面で必要とされる方法には、普通はライブラリやフレームワークのサポートが備わっている。Strutsとその派生のJava由来フレームワーク、またはASP.NETなどではサーバ側にセッション（クライアント）ごとに状態を保持させる方法が取られることが多い。あるユーザーと、ユーザーに対応するデータの取得処理は完全に隠ぺいされ、まるで単一ユーザーによって使用されるクライアントアプリケーションを作成するかのごとき考え方で開発可能だ。
>
> この方法はサーバ側にユーザーの状態を多数保持させるステートフルな方法である。このような方法は直観的でわかり易いが、いっぽうで初学者にはまるでサーバの状態とクライアントでの表示内容が密接に結び付き、お互いがお互いの状態を動的かつ容易に変更し合えるという錯覚を生じさせる。
>
> しかしながら、先に述べたようなHTTPの制限があるためにそれは幻想であり、根底にはステートレスなHTTPプロトコルを用いることによる制限は避けられない。（AjaxやWebSocketを使えばまた話は別だが）
>
> ## RESTful
> サーバー側がステートフルであると、URLにはあまり意味がなくなってくる。極端な事を言えば、http://xxxx.com/shopping というアドレスにすべてのページをマッピングし、クライアントから渡されるセッションIDによってあらゆる画面遷移を実現するということも不可能ではない。実際にそのような方法を取ることは無いものの、方向性は同じような方角を指し示しているのは間違いない。あるページにブックマークをしたが、次にそのページを開いたら同じ内容が表示されなかったとか、ブラウザで戻るとか進むボタンを使ったら最初からやり直せと言われたりだとか、そういう経験は誰にでもおありだろう。サーバ側が手綱を握っているので、クライアント側には状態を復元するための情報は開示されないのだ。
>
> ところが最近ではREST(Representational State Transfer)というキーワードをよく耳にするようになった。これはその名の通り、要求リクエストに結果を表示させるために必要なすべての具体的な情報を包含させるべき、という考え方である。原点回帰ともいえるかもしれない。RESTに従ったアプリケーション（RESTfulなアプリケーション）ではURLは非常に説明的だ。たとえば、
>
> GET http://xxxxx.com/users/withpop
> DELETE http://xxxxx.com/users/withpop
> GET http://xxxxx.com/carlist/page/2
> などといったURLが思いつく。RESTfulなアプリケーションでは全ての資源に一意に対応するURLを持つという原則を持っている。1番目はwithpopというユーザーのページ。2番目はwithpopというユーザーを消去するという処理（HTTPのDELETEメソッドを利用する）。3番目は車のリストの2ページ目を表示する。1番目にブックマークを行えば、どの環境からでも常にwithpopというユーザーのページが表示されることが期待できる。


[HTTPは本来ステートレス](https://anopara.net/2014/11/18/http%E3%81%AF%E6%9C%AC%E6%9D%A5%E3%82%B9%E3%83%86%E3%83%BC%E3%83%88%E3%83%AC%E3%82%B9/)

## http サーバ (web サーバ) の動作原理
WebサーバはHTTPサーバとも呼ばれます。HTTPプロトコルを通じてクライアントと通信を行います。このクライアントは普通はWebブラウザを指します（実はモバイルクライアントでも内部ではブラウザによって実現されています。）

Webサーバの動作原理は簡単に説明できます：

1. クライアントがTCP/IPプロトコルによってサーバまでTCP接続を設立します。
1. クライアントはサーバに対してHTTPプロトコルのリクエストパケットを送信し、サーバのリソースドキュメントを要求します。
1. サーバはクライアントに対してHTTPプロトコルの応答パケットを送信し、もし要求されたリソースに動的な言語によるコンテンツが含まれている場合、サーバが動的言語のインタープリターエンジンに"動的な内容"の処理をコールさせます。処理によって得られたデータをクライアントに返します。
1. クライアントとサーバが切断されます。クライアントはHTMLドキュメントを解釈し、クライアントの画面上に図形として結果を表示します。

簡単なHTTPタスクはこのように実現されます。見た目にはとても複雑ですが、原理はとても簡単です。気をつけなければならないのは、クライアントとサーバの間の通信は常に接続されているわけではありません。サーバが応答を送信した後クライアントと接続が切断され、次のリクエストを待ち受けます

## web の作業方法のいくつかの概念

* Request：ユーザが要求するデータ。ユーザのリクエスト情報を解析します。post、get、cookie、url等の情報を含みます。
* Response：サーバがクライアントにデータをフィードバックする必要があります。
* Conn：ユーザの毎回のリクエストリンクです。
* Handler：リクエストを処理し、返すデータを生成する処理ロジック。
* middleware: コアロジックであるApp に対して、前処理と後処理を提供するものが、ミドルウェア [ミドルウェアって何やねん！？](https://qiita.com/bussorenre/items/0ec8722a8f0ecd977104#%E3%83%9F%E3%83%89%E3%83%AB%E3%82%A6%E3%82%A7%E3%82%A2%E3%81%A3%E3%81%A6%E4%BD%95%E3%82%84%E3%81%AD%E3%82%93)

## Go lang の http パッケージについて
* [Go 言語の http パッケージにある Handle とか Handler とか HandleFunc とか HandlerFunc とかよくわからないままとりあえずイディオムとして使ってたのでちゃんと理解したメモ](https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65)

# 3 way hand shake
http://wa3.i-3-i.info/word15429.html

1. 「送っていい？」な質問パケット SYN パケット
1. 「いいわよー」な返事パケット＋「送っていい？」な質問パケット SYN パケット + ACK パケット
1. 「いいよー」な返事パケット ACK パケット

[ACKパケット](http://wa3.i-3-i.info/word15430.html)
[SYNパケット](http://wa3.i-3-i.info/word15429.html)

# ソケット通信
http://cuto.unirita.co.jp/gostudy/post/socket/
## ソケットとは
インターネットを利用するアプリケーションを作ろうというとき、直接コーディングに影響がありそうなのは、前節で紹介したようなアプリケーション層のプロトコルでしょう。 Go言語にもHTTPを扱う機能が組み込まれているので、そのAPIを使うことで、HTTPやその上のプロトコルを利用したアプリケーションを開発できます。

では、HTTPそのものはどのような仕組みで下位のレイヤーを使っているのでしょうか。 現在、ほとんどのOSでは、アプリケーション層からトランスポート層のプロトコルを利用するときのAPIとしてソケットという仕組みを利用しています。

一般に、他のアプリケーションとの通信のことをプロセス間通信（IPC：Inter Process Communication）と呼びます。 OSには、シグナル、メッセージキュー、パイプ、共有メモリなど、数多くのプロセス間通信機能が用意されています。 ソケットも、そのようなプロセス間通信の一種です。 ソケットが他のプロセス間通信と少し違うのは、アドレスとポート番号が分かればローカルのコンピュータ内だけではなく外部のコンピュータとも通信が行える点です。

アプリケーション間のインターネット通信も、このソケットを通じて行います。 たとえば通常のブラウザを利用したHTTP通信では、サーバのTCPポート80番に対して、ソケットを使ったプロセス間通信を行います。
http://ascii.jp/elem/000/001/276/1276572/

## ソケットの種類

| ソケット     | 概要                                                                                                                                                                            | 特定方法              | 通信箇所         |
| :---         | :---                                                                                                                                                                            | :---                  | :---             |
| INETドメイン | ネットワーク上でマシンを越えてのプロセス間通信                                                                                                                                  | IPアドレス+ポート番号 | 他マシンと通信可 |
| UNIXドメイン | 同じマシン上で動いているプロセスが通信を行うためのソケット。アドレス・名前空間としてファイルシステムを使用している。ファイルシステム内のinodeとしてプロセスから参照される。 | ファイル名で一致      | 自マシンのみ     |

なるほど、最近ソケット通信、ソケット通信と言ってるのはUNIXドメインソケットの事か！
[調べなきゃ寝れない！と調べたら余計に寝れなくなったソケットの話](https://qiita.com/kuni-nakaji/items/d11219e4ad7c74ece748)

> ここでソケット通信は、TCP/IPを利用する通信全般のことであるため、トランスポート層を指します。そして、そのソケットを利用してアプリケーション層のHTTP通信などが行われているというイメージですね。
>
> 簡単にまとめると、
> * TCPは、「データの内容はおいといて、通信デバイス間での通信内容を確実に送受信するためのルール」を定めているプロトコル
> * HTTPは、「TCPにさらにルールを追加して、送受信されるデータの形式や送受信タイミングをWebサイト閲覧に最適化する形に定められたルール」
> と考えると分かりやすいかもしれませんね
[知ったかぶりをしていたソケット通信の基礎を改めて学んでみる](https://qiita.com/megadreams14/items/32a3eed4661e55419e1c)

## ソケット通信のライフサイクル
ソケット通信を行う上で知っておきたいライフサイクルについて箇条書きベースで記載してみます。
下記のような処理を実際のプログラムに組み込むような形となります。

### サーバ側のライフサイクル
1. create ソケットの作成
1. bind ソケットを特定のIPアドレスとポートに紐付け
1. listen 接続の待受を開始
1. accept 接続を受信
1. close 接続を切断

### クライアント側のライフサイクル
1. create ソケットの作成
1. bind ソケットを特定のIPアドレスとポートに紐付け
1. connect リモートソケットに接続
1. close 接続を切断

## ソケット通信の基本構造
どんなソケット通信も、基本となる構成は次のような形態です。

サーバ：ソケットを開いて待ち受ける
クライアント：開いているソケットに接続し、通信を行う
Go言語の場合、サーバが呼ぶのはListen()メソッド、クライアントが呼ぶのはDial()メソッドというAPIの命名ルールが決まっており、ソケット通信でも同様です。

通信の手順はプロトコルによって異なります。 一方的な送信しかできないUDPのようなプロトコルもあれば、接続時にサーバがクライアントを認知（Accept()）して双方向にやり取りができるようになるTCPやUnixドメインソケットなどのプロトコルもあります。

Go言語におけるTCPソケットを使った通信のライフサイクルは次の図のようにまとめられます （この図ではサーバから通信を切断していますが、クライアントから切断することもできます）。

# echo サーバ
## echo(エコー)サーバー とは
> echoサーバーとは、受け取ったリクエスト内容をそのままレスポンスするサーバーのことです。
> 普通のサーバーは受け取ったリクエストに対してレスポンス内容を生成して送りますが、echoサーバーでは内部の処理を行いません。
> そのため、利用する言語のTCPサーバーとしての理論的な最高のパフォーマンスを出すことができます
https://qiita.com/kawasin73/items/3371d35166af733c2ce4

## 実装の流れ
* TCPを使う
* ポートを指定して待ち受ける
* リクエストを受け付けたら、受け取った文字列をそのまま送り返す
[Go言語でネットワークのべんきょう](http://hachibeechan.hateblo.jp/entry/2013/09/24/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF%E3%81%AE%E3%81%B9%E3%82%93%E3%81%8D%E3%82%87%E3%81%86)

もう少し詳しく

* ソケットを作成する
* ソケットをバインドする
* Listen状態になる
* コネクションを取り出す
* クライアントから文字列を受け取る
* クライアントへ文字列を返す
http://blog.bati11.info/entry/2017/03/03/223933

### stream oriented Communication (tcp, unix, unixpacket)
* Client:
  * Dial function
  * Conn interface
  * Conn reads and writes

* Server:
  * Listen function
  * Listener interface
  * Listerner.Accept
  * Conn interface again

client
1. Dial server, get connection
2. Read or write to connection
3. Close connection when done

server
1. Listen to connections
2. Accept connection, when a client dials
3. Read or write to accepted connection
4. Close connections when done

## Go での実装
[GoでたたくTCPソケット（前編）](http://ascii.jp/elem/000/001/276/1276572/)

### プロトコルとレイヤー
| レイヤーの名称     | 代表的なプロトコル  |
| :---               | :---                |
| アプリケーション層 | HTTP                |
| トランスポート層   | TCP/UDP/QUIC        |
| インターネット層   | IP                  |
| リンク層           | Wi-Fi、イーサネット |

> このうち、アプリケーションを作るために気にする必要があるのはトランスポート層よりも上のレイヤーだけです。 実際のインターネット通信では、ケーブルや無線を通してIPパケットの形でデータがやり取りされますが、 アプリケーションで直接IPパケットを作ったりするわけではありません。 HTTPやTCPのレベルで決められているルールに従って通信をすれば、それより下のレイヤーで必要になる詳細を気にすることなく、 ネットワークの向こう側にあるアプリケーションとやり取りができるわけです。

> Go言語では、HTTP、TCP、UDPについて組み込みの機能が提供されています。 実用的なアプリケーションでは、それらの機能を使って、自分のアプリケーションに必要なプロトコルを実装していくことになります。

# http サーバ (web サーバ)
# websocket
[WebSocketについて調べてみた。](https://qiita.com/south37/items/6f92d4268fe676347160)
## そもそもWebSocketとは
> Webにおいて双方向通信を低コストで行う為の仕組み。インタラクティブなWebアプリケーションではサーバから任意のタイミングでクライアントに情報の送信とかしたい事があって、例えばFacebookのチャットアプリみたいに多数のクライアントが一つのページにアクセスしてて誰かがメッセージを投稿するとそれをその他のユーザーに通知したい場合があって、そういった時に双方向通信の必要性が出てくる。
>
> 元々はWebにおいてはHTTPしか通信の選択肢が無くてHTTPのロングポーリング使って無理矢理双方向通信実現したりしてたんだけど、流石に無駄が多すぎるし辛いよねって事でWebSocketというプロトコルが作られた。
>
> WebSocketにおいては、TCP上で低コストで双方向通信が実現出来る様になってる。もちろん新しいプロトコルだからブラウザもサーバーも対応してないと使えないんだけど、最近は対応が進んでるんじゃ無いかと思う。

[WebSocketプロトコル](https://www.slideshare.net/tuvistavie/websocket-29202219)
[WebSocket / WebRTCの技術紹介](https://www.slideshare.net/mawarimichi/websocketwebrtc)
[Go言語 - WebSocketのクライアントでJSONメッセージをやり取りしてみる](http://blog.y-yuki.net/entry/2017/04/24/163000)
[Reconnecting Gorilla Websocket client on failure/lost connection](https://groups.google.com/forum/#!topic/gorilla-web/d2YHA309HY0)
https://doc01.pf.iij-engineering.co.jp/pub/sdkdoc/v1/ja_JP/websocketapi/websockif_pub_receiver.html

## websocket 使い方
大まかに、

1. websocket.DialでWebSocketのクライアントを生成してコネクション確立し、
1. websocket.Connを通じてメッセージの送受信を行うという

流れ

実際に送受信するメッセージはJSON形式として、

* websocket.JSON.Sendでサーバへのメッセージ送信、
* websocket.JSON.Receiveでサーバからのメッセージ受信を行います
