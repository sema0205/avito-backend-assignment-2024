CREATE TABLE IF NOT EXISTS banner
(
    id         SERIAL PRIMARY KEY NOT NULL,
    content    JSON               NOT NULL,
    is_active  BOOL               NOT NULL DEFAULT true,
    updated_at TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tag
(
    id         SERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS feature
(
    id         SERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS banner_feature_item
(
    id         SERIAL PRIMARY KEY NOT NULL,
    banner_id  INT UNIQUE         NOT NULL REFERENCES banner (id) ON UPDATE CASCADE ON DELETE CASCADE,
    feature_id INT                NOT NULL REFERENCES feature (id) ON UPDATE CASCADE ON DELETE CASCADE,
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS banner_tag_item
(
    id         SERIAL PRIMARY KEY NOT NULL,
    banner_id  INT                NOT NULL REFERENCES banner (id) ON UPDATE CASCADE ON DELETE CASCADE,
    tag_id     INT                NOT NULL REFERENCES tag (id) ON UPDATE CASCADE ON DELETE CASCADE,
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    UNIQUE (banner_id, tag_id)
);

INSERT INTO banner(content)
VALUES ('{"title": "some_title", "text": "some_text", "url": "some_url"}');
INSERT INTO tag(id)
VALUES (1);
INSERT INTO feature(id)
VALUES (1);

INSERT INTO banner_feature_item(banner_id, feature_id)
VALUES (1, 1);
INSERT INTO banner_tag_item(banner_id, tag_id)
VALUES (1, 1);
