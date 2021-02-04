CREATE VIEW helpful_users AS
WITH v AS (
  SELECT comment_id, SUM(vote) AS total_vote
  FROM votes
  GROUP BY comment_id
)
SELECT users.id, nickname, image_path, total_vote
FROM users
JOIN comments ON users.id = comments.user_id
JOIN v ON v.comment_id = comments.id
WHERE total_vote > 0
ORDER BY total_vote DESC
LIMIT 10;