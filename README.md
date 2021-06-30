# goutils

## bytes
以`[]byte`为基础封装出`Bytes`，支持数据编码转换、hash求值、大小写及ToString方法

- 支持编码包括：Hex、URLBase64、StdBase64
- 支持hash包括：MD5、SHA1及SHA2系列所有方法
- 支持hmac签名

## crypto
- 支持加密方法：AES-128、AES-192、AES-256、DES、3DES
- 支持链接模式：ECB、CFB/CFB8、OFB、CTR、CBC
- 支持填充模式：pkcs7、pkcs5、zero、无填充

## sign
构造接口签名

- 支持数据初始化方式包括：URL参数解析、map赋值
- 支持升序、降序对key排序
- 支持前后缀拼接签名原始数据
- 支持kv分隔符设置、多个参数之间分隔符设置
- 支持设置拼接中是否包含key或者value
- 支持`Bytes`类型所有签名格式