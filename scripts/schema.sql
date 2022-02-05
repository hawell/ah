-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema floor
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema floor
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `floor` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ;
USE `floor` ;

-- -----------------------------------------------------
-- Table `floor`.`Partner`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `floor`.`Partner` ;

CREATE TABLE IF NOT EXISTS `floor`.`Partner` (
                                                 `Id` INT NOT NULL AUTO_INCREMENT,
                                                 `Name` VARCHAR(45) NOT NULL,
    `Address` GEOMETRY NOT NULL,
    `Radius` DOUBLE NOT NULL,
    `Rating` DOUBLE NOT NULL,
    `Wood` TINYINT NOT NULL,
    `Carpet` TINYINT NOT NULL,
    `Tile` TINYINT NOT NULL,
    PRIMARY KEY (`Id`),
    SPATIAL INDEX `Location` (`Address`) VISIBLE,
    INDEX `Rating` (`Rating` ASC) VISIBLE)
    ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
