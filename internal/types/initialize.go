package types

// OSInfo holds information about the detected OS, package managers, and tools.
type OSInfo struct {
	OS             string         // Operating system detected
	PackageManager PackageManager // Package managers available
	Tools          ToolList       // Common tools
}

type UserInfo struct {
	Username  string
	FirstName string
	LastName  string
	FullName  string
	GroupName string
	Home      string
	Shell     string
}

type Flags struct {
	Debug                bool
	LogLevel             string
	Interactive          bool
	ForceBootstrap       bool
	GHAPIToken           string
	SSHKey               string
	SkipVersionCheck     bool
	InitTemplatesEnabled bool
}

type Variables struct {
	Flags       Flags
	User        UserInfo
	UserDefined map[string]interface{}
}

// InitConfig represents the configuration for the initialization processor.
type InitConfig struct {
	Init            Init                 `mapstructure:"blueprints" yaml:"blueprints" json:"blueprints" toml:"blueprints"`
	PackageManagers []PackageManagerInfo `mapstructure:"packageManagers,omitempty" yaml:"packageManagers,omitempty" json:"packageManagers,omitempty" toml:"packageManagers,omitempty"`
	Repositories    []Repository         `mapstructure:"repositories,omitempty" yaml:"repositories,omitempty" json:"repositories,omitempty" toml:"repositories,omitempty"`
	Packages        []Package            `mapstructure:"packages,omitempty" yaml:"packages,omitempty" json:"packages,omitempty" toml:"packages,omitempty"`
	Services        []Service            `mapstructure:"services,omitempty" yaml:"services,omitempty" json:"services,omitempty" toml:"services,omitempty"`
	Files           []File               `mapstructure:"files,omitempty" yaml:"files,omitempty" json:"files,omitempty" toml:"files,omitempty"`
	Directories     []Directory          `mapstructure:"directories,omitempty" yaml:"directories,omitempty" json:"directories,omitempty" toml:"directories,omitempty"`
	Templates       []Template           `mapstructure:"templates,omitempty" yaml:"templates,omitempty" json:"templates,omitempty" toml:"templates,omitempty"`
	Configuration   []Configuration      `mapstructure:"configuration,omitempty" yaml:"configuration,omitempty" json:"configuration,omitempty" toml:"configuration,omitempty"`
	Variables       Variables            `mapstructure:",squash"`
}

func (u UserInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"username":  u.Username,
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"fullName":  u.FullName,
		"groupName": u.GroupName,
		"home":      u.Home,
		"shell":     u.Shell,
	}
}

func (f Flags) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"debug":                f.Debug,
		"logLevel":             f.LogLevel,
		"interactive":          f.Interactive,
		"forceBootstrap":       f.ForceBootstrap,
		"ghAPIToken":           f.GHAPIToken,
		"sshKey":               f.SSHKey,
		"skipVersionCheck":     f.SkipVersionCheck,
		"initTemplatesEnabled": f.InitTemplatesEnabled,
	}
}
