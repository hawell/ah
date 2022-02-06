CREATE USER 'flooruser'@'localhost' IDENTIFIED BY 'floorpass';
GRANT ALL PRIVILEGES ON `floor`.* TO 'flooruser'@'localhost';
FLUSH PRIVILEGES;
