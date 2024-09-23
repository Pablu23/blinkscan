-- name: GetMangas :many
select * from manga;

-- name: GetManga :one
select * from manga
where id = $1;

-- name: CreateManga :one
insert into manga (
  provider_id, title, requested_from, created
) values (
  $1, $2, $3, NOW()
)
returning *;

-- name: GetMangasForUser :many
select m.* from manga as m
join account_subscribed_manga as asm on asm.manga_id = m.id
where asm.account_id = $1;
