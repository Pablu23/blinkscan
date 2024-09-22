// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: chapter_queries.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createChapter = `-- name: CreateChapter :one
insert into chapter (
  title, "number", manga_id
) values (
  $1, $2, $3
)
returning id, title, number, manga_id
`

type CreateChapterParams struct {
	Title   string
	Number  int32
	MangaID uuid.UUID
}

func (q *Queries) CreateChapter(ctx context.Context, arg CreateChapterParams) (Chapter, error) {
	row := q.db.QueryRow(ctx, createChapter, arg.Title, arg.Number, arg.MangaID)
	var i Chapter
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Number,
		&i.MangaID,
	)
	return i, err
}

const getChapter = `-- name: GetChapter :one
select id, title, number, manga_id from chapter
where id = $1
`

func (q *Queries) GetChapter(ctx context.Context, id uuid.UUID) (Chapter, error) {
	row := q.db.QueryRow(ctx, getChapter, id)
	var i Chapter
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Number,
		&i.MangaID,
	)
	return i, err
}

const getChapterImages = `-- name: GetChapterImages :many
select chapter_id, image_id, alignment from chapter_image
where chapter_id = $1
`

func (q *Queries) GetChapterImages(ctx context.Context, chapterID uuid.UUID) ([]ChapterImage, error) {
	rows, err := q.db.Query(ctx, getChapterImages, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChapterImage
	for rows.Next() {
		var i ChapterImage
		if err := rows.Scan(&i.ChapterID, &i.ImageID, &i.Alignment); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChapters = `-- name: GetChapters :many
select id, title, number, manga_id from chapter
`

func (q *Queries) GetChapters(ctx context.Context) ([]Chapter, error) {
	rows, err := q.db.Query(ctx, getChapters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chapter
	for rows.Next() {
		var i Chapter
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Number,
			&i.MangaID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
