# github.com/wethedevelop/gateway

## 简介
[WeTheDevelop](https://github.com/wethedevelop/wethedevelop.io)的网关项目

## 网关特性

### 同时支持 http/1.1 方式或者 http/3 调用

需要配置证书，配置方式如下：

1. 申请可用于 https 加密的可信证书或自签名证书
1. 在环境变量中增加如下配置
   ```env
   CERT_FILE=#公钥证书路径
   KEY_FILE=#私钥证书路径
   ```
   
若环境变量中不填写相关配置，则程序会正常启动，不提供 http3 网关支持；若填写，则会根据浏览器支持情况，动态在 http/1.1 或者 http3 中切换

## 技术栈简介