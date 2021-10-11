USE mysql;
UPDATE user SET authentication_string=PASSWORD('12345') WHERE User='root';
FLUSH PRIVILEGES;