package main

import (
	"gorm.io/gorm"
)

// prepare table for test

const mytableSQL = "CREATE TABLE IF NOT EXISTS `mytables` (" +
	"    `ID` int(11) NOT NULL," +
	"    `username` varchar(16) DEFAULT NULL," +
	"    `age` int(8) NOT NULL," +
	"    `phone` varchar(11) NOT NULL," +
	"    INDEX `idx_username` (`username`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
const user = "CREATE TABLE `users` (" +
	"`uuid` longtext COLLATE utf8_bin," +
	"`name` longtext COLLATE utf8_bin," +
	"`age` bigint(20) DEFAULT NULL," +
	"`version` bigint(20) DEFAULT NULL" +
	") ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin;"

func prepare(db *gorm.DB) {
	db.Exec(mytableSQL)
	db.Exec(user)
}
