CREATE TABLE `host` (
  `id`       INTEGER PRIMARY KEY AUTOINCREMENT,
  `hostname` TEXT NOT NULL,
  `ipaddr`   TEXT NOT NULL,
  `desc`     TEXT NOT NULL DEFAULT ''
);

CREATE TABLE `sgroup` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` TEXT NOT NULL UNIQUE,
  `desc` TEXT NOT NULL DEFAULT ''
);

CREATE TABLE `host_sgroup` (
  `host_id`   INTEGER NOT NULL REFERENCES `host`   (`id`),
  `sgroup_id` INTEGER NOT NULL REFERENCES `sgroup` (`id`),
  CONSTRAINT host_sgroup1 PRIMARY KEY (`host_id`, `sgroup_id`)
);

CREATE TABLE `policy` (
  `id`       INTEGER PRIMARY KEY AUTOINCREMENT,
  `from`     INTEGER NOT NULL REFERENCES `sgroup` (`id`),
  `to`       INTEGER NOT NULL REFERENCES `sgroup` (`id`),
  `port`     INTEGER NOT NULL,
  `protocol` TEXT NOT NULL DEFAULT 'tcp',
  `desc`     TEXT NOT NULL DEFAULT ''
);
