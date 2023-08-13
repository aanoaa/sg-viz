INSERT INTO `host` (`hostname`, `ipaddr`) VALUES
  ("foo-host01", "192.168.0.1"),
  ("foo-host02", "192.168.0.2"),
  ("bar-vserver", "10.0.0.1");

INSERT INTO `sgroup` (`name`) VALUES
  ("foo"),
  ("bar.example.com");

INSERT INTO `host_sgroup` (`host_id`, `sgroup_id`) VALUES
  (1, 1),
  (2, 1),
  (3, 2);

INSERT INTO `policy` (`from`, `to`, `port`) VALUES
  (1, 2, 8080);
