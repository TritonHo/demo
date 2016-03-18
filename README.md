# demo

It is a tutorial project to let people understand the backend framework step-by-step.

這是一個教學用專案，讓你一步一步理解backend framework的基本結構。

## Contents

[To be done]

## Getting Started

### Mercurial and Git

Some libraries require installation of Mercurial and Git. Please install these yourselves.
If you are using ubuntu, you may run the following command:
(Remarks: Never ask me anything about Windows environment)

sudo apt-get install mercurial git

部份程式庫需要使用Mercurial和Git，請自行安裝：
如果你使用ubuntu，你可以執行以下指令：
（請勿問我任何Windows上操作問題）

sudo apt-get install mercurial git

### Go 1.5+

The Go packages in ubuntu is outdated. Thus I suggest you downloading the latest [Go Compiler](https://golang.org/dl/) yourselves.
In following step, the Go compiler will be stored in your home directory.

1. Extract the go1.X-linux-amd64.tar.gz file into ~/go-compiler: 
	mkdir ~/go-compiler && tar -xvf go1.6.linux-amd64.tar.gz -C ~/go-compiler --strip-components 1
2. Create the a folder to store all go source code:
	mkdir ~/go
3. Open ~/.bashrc add following lines in the bottom
	export PATH=$PATH:<home_directory>/go-compiler/bin 
	export GOROOT=<home_directory>/go-compiler
	export GOPATH=/home/tritonho/go
4. restart your terminal
5. Run: go version
	If your compiler installed correctly, you will see something like "go version go1.6 linux/amd64"
6. cd ~/go/src && git clone https://github.com/TritonHo/demo.git

在ubuntu中的Go package是老舊的，我建議你下載最新[Go 編譯器](https://golang.org/dl/)

1. 把下載到的go1.X-linux-amd64.tar.gz解壓縮到 ~/go-compiler：
	mkdir ~/go-compiler && tar -xvf go1.6.linux-amd64.tar.gz -C ~/go-compiler --strip-components 1
2. 建立資料夾，用來存放所有的原始碼：
	mkdir ~/go
3. 打開 ~/.bashrc，把以下內容加到最底：
	export PATH=$PATH:<home_directory>/go-compiler/bin 
	export GOROOT=<home_directory>/go-compiler
	export GOPATH=<home_directory>/go
4. 重開你的terminal
5. 執行: go version
	如果你的Go編譯器正確設定，你應該會看到"go version go1.6 linux/amd64"
6. cd ~/go/src && git clone https://github.com/TritonHo/demo.git

### PostgreSQL 9.3+

If you are using ubuntu, I would suggest you simply run: "sudo apt-get install postgresql"
The default installation is enough for development purpose.
The ~/go/src/demo/schema folder contains readme.txt. It will teach you how to create the objects in database.

如果你正在使用ubuntu，我建議你直接執行"sudo apt-get install postgresql"
標準安裝在開發環境下足夠使用了。
資料夾 ~/go/src/demo/schema內的readme.txt，會教你怎樣一步一步地建立資料庫內所需物件

## Need Help?

If you have a question or feature request, [ask me in facebook](https://www.facebook.com/tritonho). GitHub will be used exclusively for bug reports and pull requests.
I also provide backend courses. Please contract me if you want to pay for more knowledge.

如果你有疑問或請求，[從facebook找我](https://www.facebook.com/tritonho). GitHub只用作錯誤回報和pull request.
我有提供後端教學課程。如果你願意付費來換取更多知識，歡迎找我。

##Donate

If you think this tutorial is really helpful, I suggest you donate your 2 hours salary to [Free Software foundation](https://my.fsf.org/donate/) or [Open Culture foundation](http://ocf.tw/donate/)

如果你覺得這份教學真的能幫忙，我建議你把你的２小時工資捐到[自由軟體基金會](https://my.fsf.org/donate/)或[開放文化基金會](http://ocf.tw/donate/)
