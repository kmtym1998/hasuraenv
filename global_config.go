package cli

import "os"

type GlobalConfig struct {
	HasuraenvPath ConfigPath
}

type ConfigPath struct {
	VersionsDir string
	Current     string
}

func newGlobalConfig() *GlobalConfig {
	configPathBase := os.Getenv("HOME") + "/.hasuraenv"

	return &GlobalConfig{
		HasuraenvPath: ConfigPath{
			VersionsDir: configPathBase + "/versions",
			Current:     configPathBase + "/current",
		},
	}
}
