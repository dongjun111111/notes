# Mac安装 Go + Sublime Text 开发环境
1. Go安装包下载

	官方下载地址: (https://golang.org/dl/)
	选择符合您的操作系统的安装包, 实例的安装包是 `go1.7.1.darwin-amd64.pkg`
2. 安装

	找到下载的安装包, 点击 `go1.7.1.darwin-amd64.pkg` 默认安装
3. Go环境配置
	a. 如果您安装了 `iTerm2` 和 `zsh` 请执行以下操作
	
	```
	# 打开 ~/.zshrc 文件
	sudo vim ~/.zshrc
	#找到文件中的“# User configuration” 用户配置区域加入下面的配置
	export GOPATH=/Users/qiaohongbo/Sites/gowork
	export GOBIN=/usr/local/go/bin
	export PATH="/usr/local/bin:/usr/local/sbin:$GOBIN:$PATH"
	```	
	b. 如果是使用Mac原生的终端
		
	```
	# 打开 ~/.bash_profile 文件
	sudo vim ~/.bash_profile
	# 加入下面的配置
	export GOPATH=/Users/qiaohongbo/Sites/gowork
	# 注意: 查看 GOBIN 和 PATH是否已配置了,可以参考: a
	```
	以上的任一操作完成后关闭命令终端重启后, 使用 `go env` 显示如下:
	
	```
	GOARCH="amd64"
	GOBIN="/usr/local/go/bin"
	GOEXE=""
	GOHOSTARCH="amd64"
	GOHOSTOS="darwin"
	GOOS="darwin"
	GOPATH="/Users/qiaohongbo/Sites/gowork"
	GORACE=""
	GOROOT="/usr/local/go"
	GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
	CC="clang"
	GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/57/n7b508lx6ndb644mb66v10f80000gn/T/go-build000698920=/tmp/go-build -gno-record-gcc-switches -fno-common"
	CXX="clang++"
	CGO_ENABLED="1"
	```
	如果上面的可能正常显示, 说明您的Go安装成功了 :)
	
4.	Sublime Text 安装Go语言插件 GoSublim
	A. 安装 `Package Control`
		点击 `View > Show Console` 输入以下命令:
		
		```
		# Sublime Text 3
		import urllib.request,os,hashlib; h = '2915d1851351e5ee549c20394736b442' + '8bc59f460fa1548d1514676163dafc88'; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); by = urllib.request.urlopen( 'http://packagecontrol.io/' + pf.replace(' ', '%20')).read(); dh = hashlib.sha256(by).hexdigest(); print('Error validating download (got %s instead of %s), please try manual install' % (dh, h)) if dh != h else open(os.path.join( ipp, pf), 'wb' ).write(by)
		
		# Sublime Text 2
		import urllib2,os,hashlib; h = '2915d1851351e5ee549c20394736b442' + '8bc59f460fa1548d1514676163dafc88'; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); os.makedirs( ipp ) if not os.path.exists(ipp) else None; urllib2.install_opener( urllib2.build_opener( urllib2.ProxyHandler()) ); by = urllib2.urlopen( 'http://packagecontrol.io/' + pf.replace(' ', '%20')).read(); dh = hashlib.sha256(by).hexdigest(); open( os.path.join( ipp, pf), 'wb' ).write(by) if dh == h else None; print('Error validating download (got %s instead of %s), please try manual install' % (dh, h) if dh != h else 'Please restart Sublime Text to finish installation')
		```
`Package Control` 安装完成后, `Ship + command + p` 输入 `install Package` 选择并回车.
搜索 `GoSublim` 选择并回车安装
`GoSublim` 完成安装后, `Preferences -> package settings -> GoSublime -> Settings - Uesrs` 配置 `GOPATH、GOROOT` 如下:

	```
	{
	    "env": {
	        "GOPATH": "/Users/qiaohongbo/Sites/gowork",
	        "GOROOT": "/usr/local/go"
	    }
	}
	```
重启Sublime Text, 在 `GOPATH` 目录下新建 `go` 文件, 编写Go语言就可以看到代码自动补全了 : )
			
	
		


