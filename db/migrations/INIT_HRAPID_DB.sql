-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
    `id` INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
    `name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255),
    `username` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255),
    `role` ENUM('USER', 'MANAGER', 'ADMIN') NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE INDEX `utenti_index_0` ON `users` (`id`, `username`, `role`);

CREATE TABLE IF NOT EXISTS `companies` (
    `id` INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
    `company_name` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255),
    `city` VARCHAR(255),
    `province` VARCHAR(255),
    `country` VARCHAR(255),
    `phone1` VARCHAR(255) NOT NULL,
    `phone2` VARCHAR(255),
    `email1` VARCHAR(255) NOT NULL,
    `email2` VARCHAR(255),
    `sales` DOUBLE,
    `date_added` DATETIME,
    `user_added_id` INTEGER,
    PRIMARY KEY (`id`),
    CONSTRAINT FK_companies_user_added FOREIGN KEY (user_added_id) REFERENCES users (id)
);

CREATE INDEX `clienti_index_0` ON `companies` (`id`, `company_name`);

CREATE TABLE IF NOT EXISTS `employees` (
    `id` INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
    `name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `phone1` VARCHAR(255),
    `phone2` VARCHAR(255),
    `email1` VARCHAR(255),
    `email2` VARCHAR(255),
    `role` VARCHAR(255),
    `date_added` DATETIME,
    `user_added_id` INTEGER,
    PRIMARY KEY (`id`),
    CONSTRAINT FK_employees_user_added FOREIGN KEY (user_added_id) REFERENCES users (id)
);

CREATE INDEX `persone_index_0` ON `employees` (`id`);

CREATE TABLE IF NOT EXISTS `company_employees` (
    `id` INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
    `employee_id` INTEGER NOT NULL,
    `company_id` INTEGER NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT FK_company_employees_employee_id FOREIGN KEY (employee_id) REFERENCES employees (id),
    CONSTRAINT FK_company_employees_company_id FOREIGN KEY (company_id) REFERENCES companies (id)
);

CREATE INDEX `company_employees_index_0` ON `company_employees` (`id`, `company_id`);

CREATE TABLE IF NOT EXISTS `lists` (
    `id` INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
    `description` VARCHAR(255) NOT NULL,
    `date_added` DATETIME,
    `user_added_id` INTEGER,
    PRIMARY KEY (`id`),
    CONSTRAINT FK_lists_user_added FOREIGN KEY (user_added_id) REFERENCES users (id)
);

CREATE INDEX `lists_index_0` ON `lists` (`id`, `description`);

CREATE TABLE IF NOT EXISTS `list_companies` (
    `id` INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
    `company_id` INTEGER NOT NULL,
    `list_id` INTEGER NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT FK_list_companies_company_id FOREIGN KEY (company_id) REFERENCES companies (id),
    CONSTRAINT FK_list_companies_list_id FOREIGN KEY (list_id) REFERENCES lists (id)
);

-- +migrate Down
DROP TABLE IF EXISTS `companies`;

DROP TABLE IF EXISTS `users`;

DROP TABLE IF EXISTS `employees`;

DROP TABLE IF EXISTS `company_employees`;

DROP TABLE IF EXISTS `lists`;

DROP TABLE IF EXISTS `list_companies`;