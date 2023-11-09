INSERT INTO `users` (`user_id`, `name`) VALUES
(10, 'user1'), (20, 'user2'), (30, 'user3'), (40, 'user4'), (50, 'user5'), (60, 'user6'), (70, 'user7');

INSERT INTO `friend_link` (`user1_id`, `user2_id`) VALUES
(1, 2), (1, 3), (2, 3), (2, 4), (2, 5), (2, 6), (2, 7);

INSERT INTO `block_list` (`user1_id`, `user2_id`)  VALUES
(1, 5);
