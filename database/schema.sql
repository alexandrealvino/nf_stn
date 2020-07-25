create database nf_stn;

use nf_stn;

create table invoices
(
    id       int auto_increment
        primary key,
    document   varchar(14)          not null,
    description   varchar(256)      not null,
    amount    float(64,2)           not null,
    referenceMonth  int             not null,
    referenceYear   int             not null,
    isActive    tinyint  default 1  not null,
    createdAt   datetime            not null,
    deactivatedAt datetime          not null
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