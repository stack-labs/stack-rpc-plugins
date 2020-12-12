# Stackway

`cmd`目录为网关简单用例，包括`yml`配置以及`plugin`的示范

```shell script
$ go run main.go plugin.go --config=stack_config.yml
```

- 如果作为单纯网关，不需要启动`server`，配置中`stack.server.protocol`可以使用`mock`
- `example`插件配置可以自己定义，示范的配置层级为`stackway.example`，对应代码段:
    - `cfg.Get("stackway", "example").Scan(&conf)`
