# zh-wechat
最近chatGPT异常火爆，想到将其接入到个人微信是件比较有趣的事，所以有了这个项目。项目基于[openwechat](https://github.com/eatmoreapple/openwechat)
开发
###目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复
 
# 注册openai
chatGPT注册可以参考[这里](http://www.zhanghaobk.com/archives/chatgpt%E6%B3%A8%E5%86%8C%E5%AF%B9%E6%8E%A5%E5%BE%AE%E4%BF%A1)

# 安装使用
````
# 获取项目
git clone https://github.com/zhang19523zhao/zh-wechat.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

启动前需替换config中的api_key
