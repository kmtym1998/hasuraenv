package cli

type GlobalConfig struct {
	HasuraenvPath string
}

func newGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		HasuraenvPath: "~/.hasuraenv",
	}
}
