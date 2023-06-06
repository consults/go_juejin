# go掘金签到小脚本
## 设置cookie
去掘金网站找到![img.png](imgs/img2.png)这个接口，复制下面的cookie到配置文件的cookie里
![img.png](imgs/img.png)

## 设置消息发送
在配置文件里配置msgUrl，签到完毕通知地址

## 运行
直接运行main.go 程序会自动读取配置的信息，也可以编译好之后，直接定时任务执行