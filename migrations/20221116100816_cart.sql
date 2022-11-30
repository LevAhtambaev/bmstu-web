
-- +goose Up
-- +goose StatementBegin

create table cart
(
    uuid uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
    manga uuid UNIQUE,
    foreign key (manga) references manga(uuid) ON DELETE CASCADE
--     constraint fk_customer
--         FOREIGN KEY (car)
--             REFERENCES cars(uuid)
--                 ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart;
-- +goose StatementEnd