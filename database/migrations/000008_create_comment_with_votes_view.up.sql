CREATE VIEW comments_with_votes AS
WITH v AS (
	SELECT comment_id, SUM(vote) AS total_vote
	FROM votes
	GROUP BY comment_id
)
SELECT c.id, user_id, content, COALESCE(total_vote, 0) AS vote,
created_at, updated_at, nickname, image_path, problem_id
FROM comments AS c
LEFT JOIN v ON c.id = v.comment_id
JOIN users ON c.user_id = users.id;