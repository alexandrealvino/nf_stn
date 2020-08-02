create database nf_stn;

use nf_stn;

create table IF NOT EXISTS invoices
(
    id             int auto_increment
        primary key,
    document       varchar(14) unique                     not null,
    description    varchar(256)                           not null,
    amount         bigint                           not null,
    referenceMonth int                                    not null,
    referenceYear  int                                    not null,
    isActive       bool  default 1                     not null,
    createdAt      datetime default CURRENT_TIMESTAMP     not null,
    deactivatedAt  datetime default '2020-01-01 00:00:00-00:01' ON UPDATE CURRENT_TIMESTAMP
);

create table IF NOT EXISTS users
(
    id        int auto_increment
        primary key,
    username  varchar(20)  not null,
    hash  char(60)      not null
);

# botar indix aonde eu busco e index composto, username unico

CREATE INDEX IDX_INVOICES_DOCUMENT_ISACTIVE ON invoices (document);
CREATE INDEX IDX_INVOICES_MONTH_ISACTIVE ON invoices (referenceMonth);
CREATE INDEX IDX_INVOICES_YEAR_ISACTIVE ON invoices (referenceYear);
CREATE INDEX IDX_INVOICES_DOCUMENT_MONTH ON invoices (document,referenceMonth);
CREATE INDEX IDX_INVOICES_DOCUMENT_YEAR ON invoices (document,referenceYear);
CREATE INDEX IDX_INVOICES_MONTH_YEAR ON invoices (referenceMonth,referenceYear);

INSERT INTO nf_stn.users (username, hash) VALUES ("username", "$2a$04$/GvrVH49FLVOVqbtXd99oul2Ma8Nw84dHbYqapq93R042Q98OpEAW");
