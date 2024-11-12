-- +migrate Up

ALTER TABLE companies ADD COLUMN status ENUM('ACTIVE', 'PENDING', 'INACTIVE', 'TERMINATED') NOT NULL DEFAULT 'PENDING' AFTER sales;

-- +migrate Down

ALTER TABLE companies drop column status;