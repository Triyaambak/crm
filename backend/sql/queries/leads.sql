-- name: GetAllLeads :many
SELECT * FROM leads;

-- name: CreateLead :exec
INSERT INTO leads (id,name,phone,created_at)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetLead :one
SELECT * FROM leads WHERE id = $1;

-- name: DeleteLead :exec
DELETE FROM leads WHERE id = $1;
