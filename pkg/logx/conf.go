package logx

type (
	Conf struct {
		// console,file,both
		Mode string `default:"console" options:"console,file,both"`
		// 当Mode是`file`时有效
		Path string `default:"runtime/logs"`
		// debug,info,warn,error,fatal,panic
		Level string `default:"debug" options:"debug,info,warn,error,fatal,panic"`
		// plain,json
		Encoding string `default:"plain" options:"plain,json"`
		// 保存天数
		KeepDays int `range:"0:365"`
		// daily,level,size 分割方式
		Rotation string `default:"daily" options:"daily,level,level-daily,size"`
		// 轮转时间 单位h
		RotationTime int
		// 文件大小进行分割阈值,单位`MB` 只有Rotation包含size时有效
		MaxSize int
		// 文件前缀
		FilePrefix string
		// 文件后缀
		FileSuffix string `default:".log"`
	}
)
