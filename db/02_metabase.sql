CREATE DATABASE IF NOT EXISTS metabase DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE USER 'metabase'@'%' IDENTIFIED BY 'metabase';
GRANT ALL ON metabase.* TO 'metabase'@'%'  IDENTIFIED BY 'metabase';