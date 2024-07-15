package logObject

func Get(moduleName string) LogInterface {
	return CreateConsole(moduleName)
}
