CREATE SCHEMA `govie`
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_0900_ai_ci;

USE govie;

SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------------------------------------------------------
-- Core Tables
-- ----------------------------------------------------------------------------

DROP TABLE IF EXISTS `core_entity`;
CREATE TABLE `core_entity` (
    `entityId`    	BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `systemName`	VARCHAR(100) NOT NULL COLLATE utf8mb4_0900_as_cs,
    `title`	        VARCHAR(100) NOT NULL,
    `group`         ENUM('category', 'type', 'list', 'status', 'relation'),
    `priority`      SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    `isActive`      BIT NOT NULL DEFAULT 1,
	UNIQUE INDEX uix_systemName	(`systemName`),
	INDEX ix_group_active (`group`, `isActive`)
);

DROP TABLE IF EXISTS `core_status`;
CREATE TABLE `core_status` (
    `entityId`	    BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `statusGroup`   ENUM('media'),
    `isViewable`    BIT NOT NULL DEFAULT 1,
    `isEditable`    BIT NOT NULL DEFAULT 1,

	INDEX ix_statusGroup (`statusGroup`)
);

DROP TABLE IF EXISTS `core_relationship`;
CREATE TABLE `core_relationship` (
    `entityId`	    BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `object`        VARCHAR(100),
    `data`          JSON
);

DROP TABLE IF EXISTS `core_relation`;
CREATE TABLE `core_relation` (
    `leftId`        BIGINT UNSIGNED NOT NULL,
    `leftEntity`    ENUM('media'),
    `rightId`	    BIGINT UNSIGNED NOT NULL,
    `rightEntity`   ENUM('media'),
    `priority`      SMALLINT UNSIGNED NOT NULL DEFAULT 0,

	INDEX ix_left (`leftEntity`, `leftId`, `rightEntity`, `rightId`),
	INDEX ix_right (`rightEntity`, `rightId`, `leftEntity`, `leftId`)
);


DROP TABLE IF EXISTS core_tag;
CREATE TABLE core_tag (
	tagId          BIGINT UNSIGNED PRIMARY KEY,
	typeId     	   TINYINT UNSIGNED NOT NULL,
	systemName     VARCHAR(100) NOT NULL COLLATE utf8mb4_0900_as_cs,
	title          VARCHAR(100) NOT NULL,
	priority       SMALLINT UNSIGNED NOT NULL DEFAULT 0,

	UNIQUE INDEX uix_systemName	(`systemName`)
);

DROP TABLE IF EXISTS core_tag_type;
CREATE TABLE core_tag_type (
	typeId		 BIGINT UNSIGNED NOT NULL,
	systemName     VARCHAR(100) NOT NULL COLLATE utf8mb4_0900_as_cs,
	title          VARCHAR(100) NOT NULL,

	UNIQUE INDEX uix_systemName	(`systemName`)
);


DROP TABLE IF EXISTS core_service;
CREATE TABLE core_service (
	serviceId		BIGINT UNSIGNED NOT NULL PRIMARY KEY,
	systemName      VARCHAR(100) NOT NULL COLLATE utf8mb4_0900_as_cs,
	title           VARCHAR(100) NOT NULL,
	websiteUrl      VARCHAR(300) NOT NULL,
	websitePort     SMALLINT UNSIGNED NOT NULL DEFAULT 8080,
	datasource      VARCHAR(200) NOT NULL,
	logPath         VARCHAR(100) NOT NULL,
	isActive        BIT NOT NULL DEFAULT true,

    UNIQUE INDEX uix_systemName	(`systemName`)
);

DROP TABLE IF EXISTS core_setting;
CREATE TABLE core_setting (
	settingId		BIGINT UNSIGNED NOT NULL PRIMARY KEY,
	serviceId       BIGINT UNSIGNED,
	title           VARCHAR(200) NOT NULL,
	value           VARCHAR(400) NOT NULL,
	description     VARCHAR(750) NOT NULL,
	isActive        BIT NOT NULL DEFAULT true,

    UNIQUE INDEX ix_serviceId_title (`serviceId`, `title`)
);

-- ----------------------------------------------------------------------------
-- User Tables
-- ----------------------------------------------------------------------------

DROP TABLE IF EXISTS `user_status`;
CREATE TABLE `user_status` (
    `statusId`      BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `title`         VARCHAR(100) NOT NULL,
    `isLocked`      BIT NOT NULL DEFAULT 0,
    `description`   VARCHAR(100) NOT NULL
);

DROP TABLE IF EXISTS `user_user`;
CREATE TABLE `user_user` (
    `userId`        BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `statusId`      BIGINT UNSIGNED NOT NULL,
    `firstName`     VARCHAR(50) NOT NULL,
    `lastName`      VARCHAR(50) NOT NULL,
    `middleName`    VARCHAR(50) NULL,
    `preferredName` VARCHAR(50) NULL,
    `email`         VARCHAR(254) NULL
);

DROP TABLE IF EXISTS `user_login`;
CREATE TABLE `user_login` (
    `loginId`       BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `userId`        BIGINT UNSIGNED NOT NULL,
    `username`      VARCHAR(254) NOT NULL,
    `displayName`   VARCHAR(50) NULL,
    `resetOnLogin`  BIT NOT NULL DEFAULT 0
);

DROP TABLE IF EXISTS `user_login_history`;
CREATE TABLE `user_login_history` (
    `historyId`     BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `loginId`       BIGINT UNSIGNED NOT NULL,
    `passwordId`    BIGINT UNSIGNED,
    `type`          ENUM('success', 'failure'),
    `ip`            VARCHAR(16) NULL,
    `userAgent`     VARCHAR(200) NULL,
    `attempts`      TINYINT DEFAULT 1
);

DROP TABLE IF EXISTS `user_password`;
CREATE TABLE `user_password` (
    `passwordId`    BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `loginId`       BIGINT UNSIGNED NOT NULL,
    `userId`        BIGINT UNSIGNED NOT NULL,
    `password`      VARCHAR(512) CHARACTER SET ascii COLLATE ascii_bin  NOT NULL,
    `version`		TINYINT NOT NULL,
	`salt`     		BINARY(16) NULL,
	`isActive`      BIT NOT NULL DEFAULT 1
);


-- ----------------------------------------------------------------------------
-- Media
-- ----------------------------------------------------------------------------

DROP TABLE IF EXISTS `media_asset`;
CREATE TABLE `media_asset` (
    `assetId`	    BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `title`	        VARCHAR(250) NOT NULL,
    `typeId`	    BIGINT UNSIGNED NOT NULL,
    `statusId`      BIGINT UNSIGNED NOT NULL,
    `categoryId`    BIGINT UNSIGNED NOT NULL,
    `isActive`      BIT NOT NULL DEFAULT 1
);

DROP TABLE IF EXISTS `media_asset_item`;
CREATE TABLE `media_asset_item` (
    `assetItemId`	BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `assetId`       BIGINT UNSIGNED NOT NULL,
    `file`	        VARCHAR(250) NOT NULL,
    `statusId`      BIGINT UNSIGNED NOT NULL,
    `mimetypeId`    TINYINT UNSIGNED NOT NULL,
    `storageId`     TINYINT UNSIGNED NOT NULL,
    `breakpointId`  TINYINT UNSIGNED NOT NULL,
    `properties`    JSON,
    `isActive`      BIT NOT NULL DEFAULT 1
);

DROP TABLE IF EXISTS `media_description`;
CREATE TABLE `media_description` (
    `lookupId`      BIGINT UNSIGNED NOT NULL,
    `lookupType`    ENUM('asset', 'item', 'list', 'breakpoint'),
    `description`	TEXT,

	UNIQUE INDEX uix_lookupId_lookupType (`lookupId`, `lookupType`)
);

DROP TABLE IF EXISTS `media_breakpoint`;
CREATE TABLE `media_breakpoint` (
    `breakpointId`  TINYINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `systemName`	VARCHAR(50) NOT NULL COLLATE utf8mb4_0900_as_cs,
    `title`	        VARCHAR(50) NOT NULL,
    `breakpoint`    VARCHAR(50),
    `isActive`      BIT NOT NULL DEFAULT 1,

	UNIQUE INDEX uix_systemName	(`systemName`)
);

DROP TABLE IF EXISTS `media_mimetype`;
CREATE TABLE `media_mimetype` (
    `mimetypeId`	TINYINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `typeId`        BIGINT UNSIGNED NOT NULL,
    `title`	        VARCHAR(100) NOT NULL,
    `mimetype`	    VARCHAR(75) NOT NULL,
    `extension`	    VARCHAR(4) NOT NULL
);

DROP TABLE IF EXISTS `media_storage`;
CREATE TABLE `media_storage` (
    `storageId`	    TINYINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `title`	        VARCHAR(250) NOT NULL,
    `properties`    JSON NULL,
    `isActive`      BIT NOT NULL DEFAULT 1
);

DROP TABLE IF EXISTS `media_list`;
CREATE TABLE `media_list` (
    `listId`	    BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    `listTypeId`    BIGINT UNSIGNED NOT NULL,
    `title`	        VARCHAR(100) NOT NULL,
    `isActive`      BIT NOT NULL DEFAULT 1
);

DROP TABLE IF EXISTS `media_list_lookup`;
CREATE TABLE `media_list_lookup` (
    `listId`	    BIGINT UNSIGNED NOT NULL,
    `assetId`	    BIGINT UNSIGNED NOT NULL,
    `priority`	    SMALLINT UNSIGNED NOT NULL DEFAULT 0,

	UNIQUE INDEX uix_listId_assetId	(`listId`, `assetId`)
);


-- ----------------------------------------------------------------------------
-- Video Data
-- ----------------------------------------------------------------------------

-- Ratings
-- People
-- Copyright & authors





SET FOREIGN_KEY_CHECKS = 1;
