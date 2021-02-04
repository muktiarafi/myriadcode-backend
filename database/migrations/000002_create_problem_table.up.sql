CREATE TYPE difficulty AS ENUM ('Easy', 'Medium', 'Hard');
CREATE TABLE problems(
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  description TEXT NOT NULL,
  problem_path TEXT NOT NULL,
  difficulty difficulty NOT NULL
);