package checklist

type appEnv struct {
	sprint string
	tag    string
}

func Run() int {
	app := appEnv{}
	app.github()
	app.confluence()

	return 0
}
