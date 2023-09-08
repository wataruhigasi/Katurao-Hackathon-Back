-- articles テーブルにデータを挿入
INSERT INTO `articles` (`title`, `body`, `position`) VALUES
  ('Article 1', 'This is the body of Article 1.', '{"x": 10, "y": 20}'),
  ('Article 2', 'This is the body of Article 2.', '{"x": 15, "y": 25}'),
  ('Article 3', 'This is the body of Article 3.', '{"x": 20, "y": 30}'),
  ('Article 4', 'This is the body of Article 4.', '{"x": 25, "y": 35}'),
  ('Article 5', 'This is the body of Article 5.', '{"x": 30, "y": 40}');

-- threads テーブルにデータを挿入
INSERT INTO `threads` (`title`, `position`) VALUES
  ('Thread 1', '{"x": 5, "y": 10}'),
  ('Thread 2', '{"x": 10, "y": 15}'),
  ('Thread 3', '{"x": 15, "y": 20}'),
  ('Thread 4', '{"x": 20, "y": 25}'),
  ('Thread 5', '{"x": 25, "y": 30}');

-- comments テーブルにデータを挿入
INSERT INTO `comments` (`article_id`, `body`, `author`) VALUES
  (1, 'Comment 1 for Article 1', 'Author 1'),
  (1, 'Comment 2 for Article 1', 'Author 2'),
  (2, 'Comment 1 for Article 2', 'Author 3'),
  (3, 'Comment 1 for Article 3', 'Author 4'),
  (4, 'Comment 1 for Article 4', 'Author 5');