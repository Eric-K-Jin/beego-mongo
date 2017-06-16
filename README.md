goweb
<p>使用golang+mongodb开发</p>
<p>需要下载golang的三方库mgo</p>
<p>go get gopkg.in/mgo.v2 获取mgo如果报错，请直接上GitHub上下载mgo，GitHub地址是https://github.com/go-mgo/mgo/tree/v2</p>
<br>
<ul>
<li>在conf/app.conf中修改自己的mongodb配置</li>
<li>libs中mongodb.go封装了基本CURD操作</li>
<li>libs中可以新建文件写数据库层的操作代码，可借鉴bank.go</li>
</ul>