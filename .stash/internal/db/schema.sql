-- 4Life community site schema

CREATE TABLE IF NOT EXISTS site_settings (
  key   text PRIMARY KEY,
  value text NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS nav_folders (
  id    bigserial PRIMARY KEY,
  label text NOT NULL,
  href  text NOT NULL,
  sort  int  NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS nav_items (
  id       bigserial PRIMARY KEY,
  folder_id bigint REFERENCES nav_folders(id) ON DELETE CASCADE,
  label    text NOT NULL,
  href     text NOT NULL,
  external bool NOT NULL DEFAULT false,
  sort     int  NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS pages (
  slug             text PRIMARY KEY,
  title            text NOT NULL,
  meta_description text NOT NULL DEFAULT '',
  hero_label       text,
  hero_headline    text,
  hero_body        text,
  hero_image       text,
  hero_image_alt   text,
  hero_image_side  text DEFAULT 'right',
  hero_theme       text NOT NULL DEFAULT 'dark',
  published        bool NOT NULL DEFAULT true,
  sort             int  NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS sections (
  id          bigserial PRIMARY KEY,
  page_slug   text NOT NULL REFERENCES pages(slug) ON DELETE CASCADE,
  kind        text NOT NULL,
  theme       text NOT NULL DEFAULT 'light',
  label       text,
  headline    text,
  body        text,
  bullets     jsonb,
  image       text,
  image_alt   text,
  image_side  text DEFAULT 'left',
  gallery     jsonb,
  button_label text,
  button_url  text,
  bg_image    text,
  bg_overlay  numeric(4,2) DEFAULT 0,
  divider     int DEFAULT 0,
  sort        int NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS accordion_items (
  id         bigserial PRIMARY KEY,
  section_id bigint NOT NULL REFERENCES sections(id) ON DELETE CASCADE,
  title      text NOT NULL,
  body       text NOT NULL,
  sort       int  NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS posts (
  slug         text PRIMARY KEY,
  title        text NOT NULL,
  excerpt      text NOT NULL DEFAULT '',
  cover        text NOT NULL DEFAULT '',
  external_url text NOT NULL DEFAULT '',
  body         text NOT NULL DEFAULT '',
  published_at date NOT NULL DEFAULT CURRENT_DATE,
  published    bool NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS submissions (
  id         bigserial PRIMARY KEY,
  kind       text NOT NULL,
  name       text NOT NULL DEFAULT '',
  email      text NOT NULL DEFAULT '',
  message    text NOT NULL DEFAULT '',
  extras     jsonb,
  created_at timestamptz NOT NULL DEFAULT now()
);
