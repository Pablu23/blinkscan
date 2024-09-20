-- name: GetChapters :many
select * from chapter;

-- name: GetChapter :one
select * from chapter
where id = $1;

-- name: CreateChapter :one
insert into chapter (
  title, "number", manga_id
) values (
  $1, $2, $3
)
returning *;

-- name: GetChapterImages :many
select * from chapter_image
where chapter_id = $1;
