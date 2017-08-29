package study

import (
  "fmt"
  "github.com/astaxie/beego/logs"
)

func main() {
  //创建一个BeeLogger
	log := logs.NewLogger(1000)
	//设置BeeLogger日志等级
	log.SetLevel(logs.LevelDebug)
	//console方式的日志记录方式只有一个配置参数level，
	//这里的日志级别是日志的输出等级，而上面的SetLevel设置的等级是日志写入缓冲区的等级
	//也就是说如果SetLevel设置的等级高于以参数方式提供的等级，则将以SetLevel的等级为准
	//否则的话，则最终输出的日志以参数方式提供的等级为准。
	log.SetLogger("console", fmt.Sprintf("{\"level\":%d}", logs.LevelNotice))
	log.Debug("%s", "this is a debug message")
	log.Informational("%s", "this is an informational message")
	log.Notice("%s", "this is a notice message")
	log.Flush()
	//由于beego的日志是由单独的协程输出的，这里阻塞一下程序让协程有机会执行
	fmt.Scanln()
	//关闭日志
	log.Close()
}
