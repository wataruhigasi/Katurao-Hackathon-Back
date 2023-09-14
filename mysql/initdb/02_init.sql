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

INSERT INTO `keijiban_rakugaki` (`body`, `position`) VALUES
  ('svg 1', '{"x": 5, "y": 10}'),
  ('svg 2', '{"x": 10, "y": 15}'),
  ('svg 3', '{"x": 15, "y": 20}'),
  ('svg 4', '{"x": 20, "y": 25}'),
  ('svg 5', '{"x": 25, "y": 30}');

INSERT INTO `thread_rakugaki` (`body`, `thread_id`, `position`) VALUES
  ('svg 1', '1', '{"x": 5, "y": 10}'),
  ('svg 2', '1', '{"x": 10, "y": 15}'),
  ('svg 3', '3', '{"x": 15, "y": 20}'),
  ('svg 4', '4', '{"x": 20, "y": 25}'),
  ('svg 5', '5', '{"x": 25, "y": 30}');
