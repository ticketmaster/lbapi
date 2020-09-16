package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

var GlobalConfig *Setting

// Read will open a local copy of config.toml and read the settings. If the file
// does not exist, the function will populate the settings with environmental vars.
func (c *Setting) Read() {
	if _, err := toml.DecodeFile("./etc/config.toml", &c); err != nil {
		////////////////////////////////////////////////////////////////////////
		// Avi
		////////////////////////////////////////////////////////////////////////
		c.Avi.Password = os.Getenv("AVI_PASSWORD")
		c.Avi.User = os.Getenv("AVI_USER")
		c.Avi.Tenant = os.Getenv("AVI_TENANT")
		c.Avi.SDKVersion = os.Getenv("AVI_SDK_VERSION")
		////////////////////////////////////////////////////////////////////////
		// Database
		////////////////////////////////////////////////////////////////////////
		LBDatabasePort, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
		c.Database.Database = os.Getenv("DATABASE")
		c.Database.Host = os.Getenv("DATABASE_HOST")
		c.Database.Password = os.Getenv("DATABASE_PASSWORD")
		c.Database.Port = LBDatabasePort
		c.Database.SSLMode = os.Getenv("DATABASE_SSL_MODE")
		c.Database.User = os.Getenv("DATABASE_USER")
		////////////////////////////////////////////////////////////////////////
		// Nsr
		////////////////////////////////////////////////////////////////////////
		c.Nsr.Password = os.Getenv("NSR_PASSWORD")
		c.Nsr.User = os.Getenv("NSR_USER")
		////////////////////////////////////////////////////////////////////////
		// Infoblox
		////////////////////////////////////////////////////////////////////////
		c.Infoblox.User = os.Getenv("INFOBLOX_USER")
		c.Infoblox.Password = os.Getenv("INFOBLOX_PASSWORD")
		c.Infoblox.Host = os.Getenv("INFOBLOX_HOST")
		enableIb := os.Getenv("INFOBLOX_ENABLE")
		if strings.ToLower(enableIb) == "true" {
			c.Infoblox.Enable = true
		}
		////////////////////////////////////////////////////////////////////////
		// Lbm
		////////////////////////////////////////////////////////////////////////
		c.Lbm.KeyFile = os.Getenv("LBM_KEYFILE")
		c.Lbm.PemFile = os.Getenv("LBM_PEMFILE")
		enableTLS := os.Getenv("LBMRUNTLS")
		if strings.ToLower(enableTLS) == "true" {
			c.Lbm.RunTLS = true
		}
		////////////////////////////////////////////////////////////////////////
		// Backup
		////////////////////////////////////////////////////////////////////////
		c.Backup.User = os.Getenv("BACKUP_USER")
		c.Backup.FullName = os.Getenv("BACKUP_FULLNAME")
		c.Backup.Email = os.Getenv("BACKUP_EMAIL")
		c.Backup.Password = os.Getenv("BACKUP_PASSWORD")
		c.Backup.Remote = os.Getenv("BACKUP_REMOTE")
		c.Backup.Branch = os.Getenv("BACKUP_BRANCH")
		c.Backup.RepoName = os.Getenv("BACKUP_REPONAME")
		enableBackup := os.Getenv("BACKUP_ENABLE")
		if strings.ToLower(enableBackup) == "true" {
			c.Backup.Enable = true
		}
	}
}

// Set will populate the exported Setting struct with values from either env or local toml file.
func Set() (r *Setting) {
	r = new(Setting)
	r.Read()
	return
}

func SetGlobal() {
	GlobalConfig = Set()
	return
}
