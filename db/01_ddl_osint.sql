CREATE DATABASE IF NOT EXISTS osint DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
use osint;

-- CORE ------------------------------------------------
CREATE TABLE results (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  source VARCHAR(255),
  ip_address VARCHAR(255),
  cloud_source VARCHAR(255),
  orgnization VARCHAR(255),
  product VARCHAR(255),
  service VARCHAR(255),
  socket VARCHAR(255),
  port INT,
  banner text,
  raw_data text,
  invisible BOOLEAN NOT NULL,
  scan_at DATETIME,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin AUTO_INCREMENT = 1;