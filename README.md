Hi，there



## Zing ai

* 名字只不过是一个代号

## 各部分介绍

 1. ai-back  go语言后端 使用gin 、gorm、go-redis
 2. ai-backend  react项目后台管理系统，功能比较简单，目前只有prompt管理、用户管理、版本管理
 3. ai-client flutter移动端
 4. ai-website  react 框架  ai网站


## 功能

* 版本管理
  * 版本兼容：允许用户使用多个版本，每次登录判断版本是否被允许使用，不允许则提示更新，若不更新则自动退出
  * 版本发布：并给用户提示
* 用户权限
  * 将用户拉入黑名单禁用，则app不会对其开放
* 核心功能chat
  * 文字聊天：流式回复，
  * 生成图片：使用百度文心一格

- 提供prompt



## 展示

flutter端具体看个人网站：https://holdme.fun/2024/03/23/4Zing/

web部分暂无



## 说明

1. 通过go后端对接open ai 接口 https://github.com/sashabaranov/go-openai

2. flutter端请在Constant.dart配置需要的信息

3. go端请在config/config.yaml配置需要的信息

4. web端使用 https://github.com/northes/go-moonshot kimi ai 接口实现文字回复，

   支持markdown，并有引用网站提示

5. web端在constants/index配置自己的token




```
ps:个人全栈开发，练手项目，故略有简陋
```













