package config

// Setting stores credentials and application settings.
type Setting struct {
	Avi        Avi
	Database   Database
	Infoblox   Infoblox
	Lbm        Lbm
	Nsr        Nsr
	Backup     Backup
	NetAPI     NetAPI
	Prometheus Prometheus
}

// Avi stores avi settings.
type Avi struct {
	Password   string
	SDKVersion string
	Tenant     string
	User       string
}

// Database stores database settings.
type Database struct {
	Database string
	Host     string
	Password string
	Port     int
	SSLMode  string
	User     string
}

// Infoblox stores infoblox settings.
type Infoblox struct {
	Enable   bool
	Host     string
	Password string
	User     string
}

// Nsr stores netscaler settings.
type Nsr struct {
	Password string
	User     string
}

// Lbm stores application settings.
type Lbm struct {
	AdminGroup         string
	CorsAllowedOrigins []string
	Environment        string
	GenericPRD         int
	KeyFile            string
	PemFile            string
	RunTLS             bool
}

// Git backup settings.
type Backup struct {
	User     string
	Password string
	Remote   string
	Branch   string
	RepoName string
	FullName string
	Email    string
	Enable   bool
}

// Network API.
type NetAPI struct {
	Enable bool
	URI    string
}

// Prometheus.
type Prometheus struct {
	Enable bool
	URI    string
}
