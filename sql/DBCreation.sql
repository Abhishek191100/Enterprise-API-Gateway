IF NOT EXISTS (
    SELECT * FROM INFORMATION_SCHEMA.TABLES 
    WHERE TABLE_NAME = 'apikeys' AND TABLE_SCHEMA = 'dbo'
)
BEGIN
    CREATE TABLE dbo.apikeys (
        key_id INT IDENTITY(1,1) PRIMARY KEY,
        key_value CHAR(44) NOT NULL,
        client_id VARCHAR(50) NOT NULL,
        created_at DATETIME DEFAULT GETDATE(),
        expires_at DATETIME NOT NULL,
        status INT DEFAULT 1,
        last_used_at DATETIME DEFAULT NULL
    );
END