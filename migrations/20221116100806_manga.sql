-- +goose Up
-- +goose StatementBegin
create table manga
(
    uuid        uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    name        text,
    rate        float,
    year        int,
    genre       text,
    price       int,
    volumes     int,
    description text,
    image       text
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE manga;
-- +goose StatementEnd