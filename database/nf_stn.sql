-- # SELECT host FROM mysql.user WHERE User = 'root';
-- #
-- # CREATE USER 'root'@'192.168.0.54' IDENTIFIED BY 'admin';
-- # GRANT ALL PRIVILEGES ON *.* TO 'root'@'192.168.0.54';

show databases ;

use nf_stn;

create table nf_stn.invoices
(
    id             int auto_increment
        primary key,
    document       varchar(14)                            not null,
    description    varchar(256)                           not null,
    amount         float(64, 2)                           not null,
    referenceMonth int                                    not null,
    referenceYear  int                                    not null,
    isActive       tinyint  default 1                     not null,
    createdAt      datetime                               not null,
    deactivatedAt  datetime default '0000-00-00 00:00:00' null
);

create table nf_stn.hashes
(
    id          int auto_increment
        primary key,
    profile    varchar(20)  not null,
    hash varchar(256) not null
);

-- 1,admin,$2a$04$mi4DMjIrtypqG9udFxikDusT3KK7tEDSbZ3TNSSr3g6kgC51ccJaS
--
--
1,00000000000011,patched,12,2,2022,0,2020-07-20 10:07:37,2020-07-20 10:07:37

