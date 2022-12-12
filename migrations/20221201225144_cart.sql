
-- +goose Up
-- +goose StatementBegin

create table cart
(
    uuid uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    manga uuid,
    user_uuid uuid,
    foreign key (manga) references manga(uuid) ON DELETE CASCADE,
    foreign key (user_uuid) references users(uuid) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart;
-- +goose StatementEnd