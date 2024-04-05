DROP TABLE car, owner;

CREATE TABLE IF NOT EXISTS car
(
    regNum VARCHAR(255) PRIMARY KEY,
    mark   VARCHAR(255) NOT NULL,
    model  VARCHAR(255),
    year   SMALLINT     NOT NULL,
    owner  CHAR(36),
    FOREIGN KEY (owner) REFERENCES owner (id)
);

CREATE TABLE IF NOT EXISTS owner
(
    id         CHAR(36) PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    surname    VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255)
)
