CREATE TABLE `host` (
  `id`       INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `hostname` TEXT NOT NULL UNIQUE,
  `ipaddr`   TEXT NOT NULL,
  `desc`     TEXT NOT NULL DEFAULT ''
);

CREATE TABLE `sgroup` (
  `id`   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` TEXT NOT NULL UNIQUE,
  `desc` TEXT NOT NULL DEFAULT ''
);

CREATE TABLE `host_sgroup` (
  `host_id`   INTEGER NOT NULL,
  `sgroup_id` INTEGER NOT NULL,
  CONSTRAINT host_sgroup1 PRIMARY KEY (`host_id`, `sgroup_id`)
);

CREATE TABLE `policy` (
  `id`       INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `from`     INTEGER NOT NULL,
  `to`       INTEGER NOT NULL,
  `port`     INTEGER NOT NULL,
  `protocol` TEXT NOT NULL DEFAULT 'tcp',
  `desc`     TEXT NOT NULL DEFAULT ''
);
