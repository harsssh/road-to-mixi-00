INSERT INTO `users` (`user_id`, `name`) VALUES
(1, 'user1'), (2, 'user2'), (3, 'user3'), (4, 'user4'), (5, 'user5'), (6, 'user6'), (7, 'user7');

INSERT INTO `friend_link` (`user1_id`, `user2_id`) VALUES
(1, 2), (1, 3), (2, 3), (2, 4), (2, 5), (2, 6), (2, 7);

INSERT INTO `block_list` (`user1_id`, `user2_id`)  VALUES
(1, 5);
