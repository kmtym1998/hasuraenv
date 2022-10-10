package cli

type GlobalConfig struct {
	HasuraenvPath ConfigPath
}

type ConfigPath struct {
	VersionsDir string
	Current     string
}

func newGlobalConfig(configPathBase string) *GlobalConfig {
	return &GlobalConfig{
		HasuraenvPath: ConfigPath{
			VersionsDir: configPathBase + "/versions",
			Current:     configPathBase + "/current",
		},
	}
}
