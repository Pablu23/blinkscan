-- name: GetAccounts :many
select * from account;

-- name: GetAccount :one
select * from account
where id = $1;

-- name: GetAccountByName :one
select * from account
where name = $1;

-- name: CreateAccount :one
insert into account (
  name, base64_pwd_hash, base64_pwd_salt
) values (
  $1, $2, $3
)
returning *;

-- name: GetSubscribedForAccount :many
select * 
from manga as m
join account_subscribed_manga as asm on asm.manga_id = m.id
where asm.account_id = $1;

-- name: GetViewedForAccountAndManga :many
select *
from chapter as c
join account_viewed_chapter as avc on avc.chapter_id = c.id
where account_id = $1
and c.manga_id = $2;
