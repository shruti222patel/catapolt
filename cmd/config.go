package cmd

type config struct {
	Name string `mapstructure:"name"`
	Urls []struct {
		URL     string `mapstructure:"url"`
		Browser string `mapstructure:"browser"`
	} `mapstructure:"urls"`
	Apps []struct {
		App  string `mapstructure:"app"`
		Args string `mapstructure:"args"`
	} `mapstructure:"apps"`
	ShellCmds []struct {
		Cmd []string `mapstructure:"cmd"`
	} `mapstructure:"shell_cmds"`
}
