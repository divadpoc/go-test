package libcommon

var (
	Version = "dev"
	Build   = "now"
)

func GetVersionAndBuild() string {
	return "{\"version\":" + Version + ",\"Build date\":" + Build + "}"
}

func GetVersion() string {
	return Version
}

func GetBuild() string {
	return Build
}
