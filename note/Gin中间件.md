# Gin中间件

## Default

```go
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
```

自带两个中间件，Logger用来打印日志输出，Recovery用来panic处理



gin通过

```go
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes
```

来使用中间件

### 自定义中间件

中间件是一种gin.handfunc类型



# swag

```shell
go get -u github.com/swaggo/swag/cmd/swag
```