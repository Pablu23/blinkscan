-- name: CreateSession :one
insert into session (
  account_id
) values (
  $1
)
RETURNING *;

-- name: GetUserForSession :one
select * from account as a
join session as s on a.id = s.account_id
where s.id = $1;
