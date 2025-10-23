package utils

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/Abhishek191100/Enterprise-API-Gateway/logger"
	"gopkg.in/yaml.v3"
)

type DBOptions struct {
	Encrypt                bool  `yaml:"encrypt"`
	TrustServerCertificate bool  `yaml:"trustServerCertificate"`
	ConnectionTimeout      int32 `yaml:"connectionTimeout"`
}

type MSSQLConnection struct {
	Host     string    `yaml:"host"`
	Port     string    `yaml:"port"`
	User     string    `yaml:"user"`
	Password string    `yaml:"password"`
	Database string    `yaml:"database"`
	Options  DBOptions `yaml:"options"`
}

// Load MSSQL config from YAML
func loadMSSQLConfig(path string) (*MSSQLConnection, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var raw map[string]interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	mssqlRaw, ok := raw["mssql"]
	if !ok {
		return nil, fmt.Errorf("mssql config not found")
	}
	mssqlBytes, err := yaml.Marshal(mssqlRaw)
	if err != nil {
		return nil, err
	}
	var cfg MSSQLConnection
	if err := yaml.Unmarshal(mssqlBytes, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func OpenMSSQLDBConnection() (*sql.DB, error) {
	cfg, err := loadMSSQLConfig("../../config/db.yaml")
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to load MSSQL config: %v", err), "ERROR")
		return nil, err
	}
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=%t;TrustServerCertificate=%t;connection timeout=%d",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.Database,
		cfg.Options.Encrypt, cfg.Options.TrustServerCertificate, cfg.Options.ConnectionTimeout)
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to open MSSQL connection: %v", err), "ERROR")
		return nil, err
	}
	if err := db.Ping(); err != nil {
		logger.Log(fmt.Sprintf("Failed to ping MSSQL: %v", err), "ERROR")
		return nil, err
	}
	return db, nil
}

func GetAllAPIKeyHashes() ([]string, error) {
	
	db, err := OpenMSSQLDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	rows, err := db.Query("SELECT key_value FROM dbo.apikeys")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var hashes []string
	for rows.Next() {
		var hash string
		if err := rows.Scan(&hash); err != nil {
			return nil, err
		}
		hashes = append(hashes, hash)
	}
	return hashes, nil
}