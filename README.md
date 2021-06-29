# goutils

## xtype
以`[]byte`为基础封装出`XBS`，支持数据编码转换、hash求值、大小写及ToString方法

支持编码包括：Hex、URLBase64、StdBase64
支持hash包括：MD5、SHA1及SHA2系列所有方法

## xcrypto
支持加密方法：AES-128、AES-192、AES-256、DES、3DES
支持链接模式：ECB、CFB、OFB、CTR、CBC

## sign
构造接口签名，支持md5、sha1、sha2、hmac算法，支持根据url请求拆分参数，排序拼接等操作