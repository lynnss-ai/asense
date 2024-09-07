# asense
### 基于go-zero的后端服务

### 1. 项目中关于go-zero相关指令
```
(1) 初始化模版
    $ goctl template init --home ./template
(2) 代码格式化
    $ goctl api format --dir .
(3) 根据api文件生成MD接口文档
    $ goctl api go doc --o=./doc --dir=../
(4) 根据api文件生成代码
    $ goctl api go -api ./apis/systemmanage.api -dir . -style gozero --home ../../../template
```

### 2. 项目中关于go-zero的其他配置
```
(1) 健康检测
    - 对应配置开启
    - 健康检查默认端口为6470
    - 默认Path为 /health
Health:
  Enable: true         # 是否开启健康检查
  Port: 8030           # 健康检查端口
  HealthPath: /health  # 健康检查路径
  
(2) 限流
RestConf struct {
    ...
    MaxConns int    `json:",default=10000"` // 最大并发连接数，默认值为 10000 qps
    ... 
}
```