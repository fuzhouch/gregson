package gregson

type JWTSetting struct {
}

// Setting specifies configuration when building web app.
type Setting struct {
	AppName    string
	Prometheus PromSetting
	JWT        JWTSetting
}

// NewSetting initializes a setting object with default values.
func NewSetting(appName string) *Setting {
	s := new(Setting)
	s.AppName = appName
	setPromDefault(&s.Prometheus, appName)
	return s
}
