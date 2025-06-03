-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES (?, ?, ?);

-- name: CreateCourse :exec
INSERT INTO courses (id, name, category_id, price)
VALUES (?,?,?,?);
