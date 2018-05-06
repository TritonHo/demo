# demo

It is a tutorial project to let people understand the backend framework step-by-step.

From phase1 to phase11, you will see know a CRUD endpoint is refactored from 250 lines to only 80 lines.
Also, you will learn modern coding skills, such as struct tag(~=Java annotation), reflection, middleware.

這是一個教學用專案，讓你一步一步理解backend framework的基本結構。
從第一版到第十一版，你將會看到一個250行的CRUD endpoint，被改寫成只有80行。
而且，你將會學習到現在編程技巧，例如：struct tag（大約等同Java annotation），reflection和middleware。


## Contents

| Branch | Description |
| -----------|--------|
| phase1 | Demonstrate the codes written by novice developer. No attention to code reusability decreases maintainability and productivity. | 
| phase2 | Introduce of mux library. The routing information is centralized in main.go. | 
| phase3 | Remove hardcoded database username and password in source code. Introduces dependency injection of database object. | 
| phase4 | Introduce ORM. | 
| phase5 | Introduce input binder. Also, demonstrates the power(and pain) of reflection. | 
| phase6 | Introduce input validator. | 
| phase7 | Introduce middleware to handle http output. | 
| phase8 | Introduce database transaction manager. Remove of global database object in handler. | 
| phase9 | Add user model into the system. Introduce jwt authorization in middleware. | 
| phase10 | Wrap up repetitive code in CRUD handler | 
| phase11 | Introduce 3rd party middleware. Make the system become production-ready. | 


| 分支 | 說明 |
| -----------|--------|
| phase1 | 一個菜鳥所寫的程式碼。沒有著重程式重用性，讓程式變得難以維護，也降低人們的生產力。 | 
| phase2 | 引入mux程式庫。Routing資料現在全集中在main.go。 | 
| phase3 | 刪掉在程式碼中hardcoded的資料庫用戶名字和密碼。引入資料庫物件的dependency injection。 | 
| phase4 | 引入ORM. | 
| phase5 | 引入輸入處理器，也說明reflection的強大（和痛苦）。 | 
| phase6 | 引入輸入檢查器。 | 
| phase7 | 引入中間件來負責HTTP輸出。 | 
| phase8 | 引入資料庫交易管理者。刪掉在handler的全域資料庫物件。 | 
| phase9 | 加入用戶物件到系統。在中間件層加入jwt身份驗證 | 
| phase10 | 把重複性程式碼變成獨立程序| 
| phase11 | 引入第三方套件，讓系統質素到達能上線級別。 | 

## Getting Started

### Mercurial and Git

Some libraries require installation of Mercurial and Git. Please install these yourselves.<br />
If you are using ubuntu, you may run the following command:<br />
(Remarks: Never ask me anything about Windows environment)

sudo apt-get install mercurial git

部份程式庫需要使用Mercurial和Git，請自行安裝：<br />
如果你使用ubuntu，你可以執行以下指令：<br />
（請勿問我任何Windows上操作問題）

sudo apt-get install mercurial git

### Go 1.5+

The Go packages in ubuntu is outdated. Thus I suggest you downloading the latest [Go Compiler](https://golang.org/dl/) yourselves.<br />
In following step, the Go compiler will be stored in your home directory.

1. Extract the go1.X-linux-amd64.tar.gz file into ~/go-compiler: <br />
	mkdir ~/go-compiler && tar -xvf go1.6.linux-amd64.tar.gz -C ~/go-compiler --strip-components 1
2. Create the a folder to store all go source code:<br />
	mkdir ~/go
3. Open ~/.bashrc add following lines in the bottom<br />
	export PATH=$PATH:[home_directory]/go-compiler/bin <br />
	export GOROOT=[home_directory]/go-compiler <br />
	export GOPATH=[home_directory]/go <br />
4. restart your terminal
5. Run: go version <br />
	If your compiler installed correctly, you will see something like "go version go1.6 linux/amd64"
6. cd ~/go/src && git clone https://github.com/TritonHo/demo.git

在ubuntu中的Go package是老舊的，我建議你下載最新[Go 編譯器](https://golang.org/dl/)

1. 把下載到的go1.X-linux-amd64.tar.gz解壓縮到 ~/go-compiler：<br />
	mkdir ~/go-compiler && tar -xvf go1.6.linux-amd64.tar.gz -C ~/go-compiler --strip-components 1
2. 建立資料夾，用來存放所有的原始碼：<br />
	mkdir ~/go
3. 打開 ~/.bashrc，把以下內容加到最底：<br />
	export PATH=$PATH:[home_directory]/go-compiler/bin <br />
	export GOROOT=[home_directory]/go-compiler<br />
	export GOPATH=[home_directory]/go<br />
4. 重開你的terminal
5. 執行: go version<br />
	如果你的Go編譯器正確設定，你應該會看到"go version go1.6 linux/amd64"
6. cd ~/go/src && git clone https://github.com/TritonHo/demo.git

### PostgreSQL 9.3+

If you are using ubuntu, I would suggest you simply run: "sudo apt-get install postgresql"<br />
The default installation is enough for development purpose.<br />
The ~/go/src/demo/schema folder contains readme.txt. It will teach you how to create the objects in database.

如果你正在使用ubuntu，我建議你直接執行"sudo apt-get install postgresql"<br />
標準安裝在開發環境下足夠使用了。<br />
資料夾 ~/go/src/demo/schema內的readme.txt，會教你怎樣一步一步地建立資料庫內所需物件

## Need Help?

If you have a question or feature request, [ask me in facebook](https://www.facebook.com/tritonho). GitHub will be used exclusively for bug reports and pull requests.<br />
I also provide backend courses. Please contract me if you want to pay for more knowledge.

如果你有疑問或請求，[從facebook找我](https://www.facebook.com/tritonho). GitHub只用作錯誤回報和pull request.<br />
我有提供後端教學課程。如果你願意付費來換取更多知識，歡迎找我。

## Donate

If you think this tutorial is really helpful, I suggest you donate your 2 hours salary to [Free Software foundation](https://my.fsf.org/donate/) or [Open Culture foundation](http://ocf.tw/donate/)

如果你覺得這份教學真的能幫忙，我建議你把你的２小時工資捐到[自由軟體基金會](https://my.fsf.org/donate/)或[開放文化基金會](http://ocf.tw/donate/)
