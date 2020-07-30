create database nf_stn;

use nf_stn;

create table invoices
(
    id             int auto_increment
        primary key,
    document       varchar(14)                            not null,
    description    varchar(256)                           not null,
    amount         float(64, 2)                           not null,
    referenceMonth int                                    not null,
    referenceYear  int                                    not null,
    isActive       tinyint  default 1                     not null,
    createdAt      datetime default CURRENT_TIMESTAMP     not null,
    deactivatedAt  datetime default '2020-01-01 00:00:00-00:01' ON UPDATE CURRENT_TIMESTAMP
);

create table users
(
    id        int auto_increment
        primary key,
    username  varchar(20)  not null,
    password  char(8)      not null
);

# Invoice
#     ReferenceMonth : INTEGER
#     ReferenceYear : INTEGER
#     Document : VARCHAR(14)
#     Description : VARCHAR(256)
#     Amount : DECIMAL(16, 2)
#     IsActive : TINYINT
#     CreatedAt  : DATETIME
#     DeactivatedAt : DATETIME