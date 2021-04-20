# Install MySQL on Ubuntu 20.04

sudo apt update
sudo apt install mysql-server
sudo systemctl status mysql
sudo mysql_secure_installation

sudo mysql
CREATE USER 'kh'@'localhost' IDENTIFIED BY '1234';
# ALTER USER 'kh'@'localhost' IDENTIFIED BY '1234';
# FLUSH PRIVILEGES;
# REVOKE ALL ON *.* FROM 'kh'@'localhost';
GRANT CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, REFERENCES, RELOAD on *.* TO 'kh'@'localhost';
FLUSH PRIVILEGES;
# SHOW GRANTS FOR 'kh'@'localhost';
EXIT

mysql -u kh -p

# ref:
# https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-20-04


# Allow Remote Access to MySQL
# ref:
# https://www.digitalocean.com/community/tutorials/how-to-allow-remote-access-to-mysql