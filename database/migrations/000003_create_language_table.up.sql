CREATE TABLE languages(
  id SERIAL PRIMARY KEY,
  prog_lang VARCHAR(50) NOT NULL,
  file_path VARCHAR(50) NOT NULL,
  initial_code TEXT NOT NULL,
  test_code TEXT NOT NULL,
  problem_id INTEGER NOT NULL REFERENCES problems(id) ON DELETE
  CASCADE,
  UNIQUE(prog_lang, problem_id)
);