package cli

import (
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ApiConfigs = apiConfigs

// SetApiConfig sets the configs.
func SetApiConfig(c ApiConfigs) {
	configs = c
}

func GetApplianceConfigs() ApiConfigs {
	return configs
}

// SetName sets the name of the APIConfig.
func (a *APIConfig) SetName(name string) {
	a.name = name
}

// GetName returns the name of the APIConfig
func (a *APIConfig) GetName() string {
	return a.name
}

// SetViper sets the restish viper variable
func SetViper(v *viper.Viper) {
	apis = v
}

// SetCurrentConfig sets the currentConfig.
func SetCurrentConfig(apiName string) bool {
	if cfg, ok := configs[apiName]; ok {
		currentConfig = cfg
		return true
	}
	return false
}

// GetCurrentConfig returns the currently set APIConfig
func GetCurrentConfig() *APIConfig {
	return currentConfig
}

// SetTTY sets the tty.
func SetTTY(b bool) {
	tty = b
}

// SetAurora sets the aurora.
func SetAurora(a aurora.Aurora) {
	au = a
}

// ResetRegistries resets the registries used for internal bookkeeping.
func ResetRegistries() {
	authHandlers = map[string]AuthHandler{}
	contentTypes = []contentTypeEntry{}
	encodings = map[string]ContentEncoding{}
	linkParsers = []LinkParser{}
	loaders = []Loader{}
}

// GenericRequest exposes the internal generic function.
func GenericRequest(method string, addr string, args []string) {
	generic(method, addr, args)
}

// EditRequest exposes the internal edit function.
func EditRequest(
	addr string,
	args []string,
	interactive,
	noPrompt bool,
	exitFunc func(int),
	editMarshal func(interface{}) ([]byte, error),
	editUnmarshal func([]byte, interface{}) error,
	ext string,
) {
	edit(addr, args, interactive, noPrompt, exitFunc, editMarshal, editUnmarshal, ext)
}

// InitConfig calls the internal initConfig function.
func InitConfig(appName string) {
	initConfig(appName, "")
}

// InitCache calls the internal initCache function.
func InitCache(appName string) {
	initCache(appName)
}

// MatchTemplate calls the internal matchTemplate function.
func MatchTemplate(url, template string) string {
	return matchTemplate(url, template)
}

// EnableVerbose sets the enableVerbose boolean to true.
func EnableVerbose() {
	enableVerbose = true
}

// IsVerbose returns the state of the internal enableVerbose boolean.
func IsVerbose() bool {
	return enableVerbose
}

// GetInteractiveConfigureClosure exposes the internal configure function
func GetInteractiveConfigureClosure() func(cmd *cobra.Command, args []string) {
	return askInitAPIDefault
}
