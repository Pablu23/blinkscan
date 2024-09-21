create table provider (
  id uuid PRIMARY KEY not null default gen_random_uuid(),
  url text not null,
  name text not null
);

create table image (
  id uuid PRIMARY KEY not null default gen_random_uuid(),
  path text not null
);

create table account (
  id uuid PRIMARY KEY not null default gen_random_uuid(),
  name text not null unique,
  base64_pwd_hash text not null,
  base64_pwd_salt text not null
);

create table manga (
  id uuid PRIMARY KEY not null default gen_random_uuid(),
  provider_id uuid not null,
  title text not null,
  thumbnail_id uuid,
  latest_chapter int,
  requested_from uuid,
  last_updated timestamp,
  created timestamp not null,
  FOREIGN KEY(provider_id) REFERENCES provider(id),
  FOREIGN KEY(thumbnail_id) REFERENCES image(id),
  FOREIGN KEY(requested_from) REFERENCES account(id)
);

create table chapter (
  id uuid PRIMARY KEY not null default gen_random_uuid(),
  title text not null,
  "number" int not null,
  manga_id uuid not null,
  FOREIGN KEY(manga_id) REFERENCES manga(id)
);

create table chapter_image (
  chapter_id uuid not null default gen_random_uuid(),
  image_id uuid not null,
  alignment int not null,
  PRIMARY KEY(chapter_id, image_id),
  FOREIGN KEY(chapter_id) REFERENCES chapter(id),
  FOREIGN KEY(image_id) REFERENCES image(id)
);

create table account_subscribed_manga (
  id uuid PRIMARY KEY not null default gen_random_uuid(),
  account_id uuid not null,
  manga_id uuid not null,
  FOREIGN KEY(manga_id) REFERENCES manga(id),
  FOREIGN KEY(account_id) REFERENCES account(id)
);

create table account_viewed_chapter (
  id uuid PRIMARY KEY not null default gen_random_uuid(),
  account_id uuid not null,
  chapter_id uuid not null,
  viewed_at timestamp not null,
  FOREIGN KEY(chapter_id) REFERENCES chapter(id),
  FOREIGN KEY(account_id) REFERENCES account(id)
);
