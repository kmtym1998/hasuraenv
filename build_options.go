package cli

type BuildOptions struct {
	Version        string
	ConfigPathBase string
}

func NewBuildOptions(version, configPathBase string) BuildOptions {
	return BuildOptions{
		Version:        version,
		ConfigPathBase: configPathBase,
	}
}
