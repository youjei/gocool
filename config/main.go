package config

//fmt.Println("config show")
var Configution map[string]interface{} = map[string]interface{}{
	//网站名称
	"website": "gocool framework v1.0",

	//网站路径
	"basePath": "d:/workspace/gocool",

	//网站插件
	"extension": "/gocool/extensions.go",

	//多语言设置
	"language": "zh-cn",

	//网站主题
	"theme": "default",
}
