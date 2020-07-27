create database api_stn;

use api_stn;

create table cargo
(
    id    int auto_increment
        primary key,
    cargo varchar(25) null,
    constraint cargo_cargo_uindex
        unique (cargo)
);

create table funcionarios
(
    id       int auto_increment
        primary key,
    nome     varchar(50) not null,
    idade    int(3)      not null,
    cpf      char(11)    not null,
    id_cargo int         not null,
    constraint funcionarios_cpf_uindex
        unique (cpf),
    constraint funcionarios_cargo_id_fk
        foreign key (id_cargo) references cargo (id)
);

create index funcionarios_nome_IX
    on funcionarios (nome);


insert into api_stn.cargo (cargo) values ('Engenheiro de Software');
insert into api_stn.cargo (cargo) values ('Médico');
insert into api_stn.cargo (cargo) values ('Advogado');
insert into api_stn.cargo (cargo) values ('Analista');
insert into api_stn.cargo (cargo) values ('Acessor de Investimentos');

# api_go
create database api_go;

use api_go;

create table tickers
(
    id       int auto_increment
        primary key,
    symbol   varchar(10)            not null,
    value    float(64,2)            not null,
    quota    float(64,2)            null,
    avgPrice float(64,2)            null,
    previousClose float(64,2)       null,
    lastChangePercent float(64,2)   null,
    changeFromAvgPrice float(64,2)
);

insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('ABEV3', 75.28, 30.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('B3SA3', 90.22, 25.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('BIDI4', 89.51, 10.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('BRKM5', 238.13, 30.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('CSNA3', 130.55, 25.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('ELET3', 80.92, 10.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('ELET6', 40.69, 30.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('EGIE3', 53.68, 25.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('FLRY3', 62.58, 10.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('ITUB4', 69.28, 30.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('ITSA4', 67.35, 25.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('KLBN11', 64.86, 10.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('PRIO3', 48.75, 30.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('STBP3', 152.02, 25.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('SHUL4', 257.17, 10.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('TAEE4', 63.78, 30.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('TAEE11', 31.12, 25.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('TRIS3', 82.70, 10.00, 0,0,0,0);
insert into tickers (symbol, value, quota, avgPrice, previousClose, lastChangePercent, changeFromAvgPrice) values ('WEGE3', 412.78, 10.00, 0,0,0,0);

create table buys
(
    id       int auto_increment
        primary key,
    symbol   varchar(10) not null,
    price    float(64,2) not null,
    quantity int         not null,
    date     char(10)        not null
);
create index buys_symbol_IX
    on buys (symbol);
create index buys_date_IX
    on buys (date);

insert into buys (symbol, price, quantity, date) values ('BIDI4', 72.08, 4, '15/01/2020');
insert into buys (symbol, price, quantity, date) values ('BRKM5', 68.90, 2, '15/01/2020');
insert into buys (symbol, price, quantity, date) values ('ITUB4', 69.28, 2, '15/01/2020');
insert into buys (symbol, price, quantity, date) values ('ITSA4', 67.35, 5, '15/01/2020');
insert into buys (symbol, price, quantity, date) values ('EGIE3', 53.68, 1, '15/01/2020');
insert into buys (symbol, price, quantity, date) values ('FLRY3', 62.58, 2, '15/01/2020');

insert into buys (symbol, price, quantity, date) values ('TAEE4', 63.78, 6, '20/01/2020');
insert into buys (symbol, price, quantity, date) values ('STBP3', 72.72, 9, '20/01/2020');
insert into buys (symbol, price, quantity, date) values ('B3SA3', 90.22, 2, '20/01/2020');
insert into buys (symbol, price, quantity, date) values ('ABEV3', 75.28, 4, '20/01/2020');
insert into buys (symbol, price, quantity, date) values ('WEGE3', 74.44, 2, '20/01/2020');
insert into buys (symbol, price, quantity, date) values ('TRIS3', 82.70, 5, '20/01/2020');
insert into buys (symbol, price, quantity, date) values ('CSNA3', 59.80, 4, '20/01/2020');
insert into buys (symbol, price, quantity, date) values ('ELET3', 80.92, 2, '20/01/2020');

insert into buys (symbol, price, quantity, date) values ('SHUL4', 23.52, 2, '24/01/2020');
insert into buys (symbol, price, quantity, date) values ('KLBN11', 64.86, 3, '24/01/2020');

insert into buys (symbol, price, quantity, date) values ('BIDI4', 17.43, 1, '27/01/2020');
insert into buys (symbol, price, quantity, date) values ('ELET6', 40.69, 1, '27/01/2020');
insert into buys (symbol, price, quantity, date) values ('STBP3', 79.30, 10, '27/01/2020');
insert into buys (symbol, price, quantity, date) values ('BRKM5', 71.88, 2, '27/01/2020');
insert into buys (symbol, price, quantity, date) values ('WEGE3', 82.06, 2, '27/01/2020');
insert into buys (symbol, price, quantity, date) values ('CSNA3', 70.75, 5, '27/01/2020');

insert into buys (symbol, price, quantity, date) values ('TAEE11', 31.12, 1, '31/01/2020');
insert into buys (symbol, price, quantity, date) values ('PRIO3', 48.75, 1, '31/01/2020');
insert into buys (symbol, price, quantity, date) values ('WEGE3', 81.00, 2, '31/01/2020');
insert into buys (symbol, price, quantity, date) values ('WEGE3', 41.00, 1, '31/01/2020');
insert into buys (symbol, price, quantity, date) values ('BRKM5', 97.35, 3, '31/01/2020');

insert into buys (symbol, price, quantity, date) values ('SHUL4', 89.52, 6, '17/02/2020');
insert into buys (symbol, price, quantity, date) values ('SHUL4', 14.89, 1, '17/02/2020');

insert into buys (symbol, price, quantity, date) values ('SHUL4', 129.24, 9, '02/03/2020');
insert into buys (symbol, price, quantity, date) values ('WEGE3', 134.28, 3, '02/03/2020');

insert into buys (symbol, price, quantity, date) values ('ITSA4', 54.30, 5, '11/03/2020');

insert into buys (symbol, price, quantity, date) values ('ITSA4', 8.92, 1, '30/04/2020');

insert into buys (symbol, price, quantity, date) values ('SHUL4', 11.6, 200, '10/06/2020');
insert into buys (symbol, price, quantity, date) values ('SHUL4', 11.4, 66, '10/06/2020');
insert into buys (symbol, price, quantity, date) values ('WEGE3', 46.00, 100, '10/06/2020');

## TODO put in schema
create table prices
(
    id       int auto_increment
        primary key,
    symbol   varchar(10) not null,
    lastPrice    float(64,2) ,
    lastClosePrice    float(64,2) ,
    weekResult    float(64,2) ,
    monthResult    float(64,2) ,
    yearResult    float(64,2),
    lastUpdate     varchar(10)
);

insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('ABEV3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('BIDI4', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('WEGE3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('SHUL4', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('BRKM5', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('STBP3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('CSNA3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('B3SA3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('TRIS3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('ELET3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('ITUB4', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('ITSA4', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('KLBN11', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('TAEE4', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('FLRY3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('EGIE3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('PRIO3', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('ELET6', 0,0,0,0,0,'0000000000');
insert into prices (symbol, lastPrice, lastClosePrice, weekResult, monthResult, yearResult, lastUpdate) values ('TAEE11', 0,0,0,0,0,'0000000000');

DELETE FROM buys WHERE id = 87;
DELETE FROM buys WHERE id = 72;
DELETE FROM buys WHERE id = 71;
DELETE FROM buys WHERE id = 86;
DELETE FROM buys WHERE id = 54;
DELETE FROM buys WHERE id = 55;
DELETE FROM buys WHERE id = 56;

DELETE FROM tickers WHERE id = 24;
DELETE FROM tickers WHERE id = 3;
DELETE FROM tickers WHERE id = 4;
DELETE FROM tickers WHERE id = 5;
DELETE FROM tickers WHERE id = 6;

UPDATE tickers  SET value = '100.00', quota = 10.00 WHERE id = 7;

show databases ;


SHOW VARIABLES LIKE 'max_connections';
SET GLOBAL max_connections = 1000;


# buys
#
# 1,BIDI4,72.08,4,15/01/2020
# 2,BRKM5,68.9,2,15/01/2020
# 3,ITUB4,69.28,2,15/01/2020
# 4,ITSA4,67.35,5,15/01/2020
# 5,EGIE3,53.68,1,15/01/2020
# 6,FLRY3,62.58,2,15/01/2020
# 7,TAEE4,63.78,6,20/01/2020
# 8,STBP3,72.72,9,20/01/2020
# 9,B3SA3,90.22,2,20/01/2020
# 10,ABEV3,75.28,4,20/01/2020
# 11,WEGE3,74.44,2,20/01/2020
# 12,TRIS3,82.7,5,20/01/2020
# 13,CSNA3,59.8,4,20/01/2020
# 14,ELET3,80.92,2,20/01/2020
# 15,SHUL4,23.52,2,24/01/2020
# 16,KLBN11,64.86,3,24/01/2020
# 17,BIDI4,17.43,1,27/01/2020
# 18,ELET6,40.69,1,27/01/2020
# 19,STBP3,79.3,10,27/01/2020
# 20,BRKM5,71.88,2,27/01/2020
# 21,WEGE3,82.06,2,27/01/2020
# 22,CSNA3,70.75,5,27/01/2020
# 23,TAEE11,31.12,1,31/01/2020
# 24,PRIO3,48.75,1,31/01/2020
# 25,WEGE3,81,2,31/01/2020
# 26,WEGE3,41,1,31/01/2020
# 27,BRKM5,97.35,3,31/01/2020
# 28,SHUL4,89.52,6,17/02/2020
# 29,SHUL4,14.89,1,17/02/2020
# 30,SHUL4,129.24,9,02/03/2020
# 31,WEGE3,134.28,3,02/03/2020
# 32,ITSA4,54.3,5,11/03/2020
# 33,ITSA4,8.92,1,30/04/2020
# 34,SHUL4,2320,200,10/06/2020
# 35,SHUL4,752.4,66,10/06/2020
# 36,WEGE3,4600,100,10/06/2020
#
# tickers
#
# 1,ABEV3,75.28,0.77,18.82,14.6,-0.54,-22.42
# 2,B3SA3,90.22,0.92,45.11,66.37,0.36,47.13
# 3,BIDI4,89.51,0.91,17.9,18.28,2.93,2.12
# 4,BRKM5,238.13,2.43,34.02,24.08,-2.15,-29.22
# 5,CSNA3,130.55,1.33,14.51,11.8,1.2,-18.68
# 6,ELET3,80.92,0.83,40.46,38.67,-0.23,-4.42
# 7,ELET6,40.69,0.42,40.69,39.27,-0.15,-3.49
# 8,EGIE3,53.68,0.55,53.68,44.2,0.64,-17.66
# 9,FLRY3,62.58,0.64,31.29,24.87,-1.5,-20.52
# 10,ITUB4,69.28,0.71,34.64,27.05,-2.35,-21.91
# 11,ITSA4,67.35,0.69,11.87,10.52,0,-11.37
# 12,KLBN11,64.86,0.66,21.62,21.14,0.71,-2.22
# 13,PRIO3,48.75,0.5,48.75,40.47,-1.65,-16.98
# 14,STBP3,152.02,1.55,8,5.81,-2.52,-27.38
# 15,SHUL4,3329.57,34.03,11.72,12.1,-0.82,3.24
# 16,TAEE4,63.78,0.65,10.63,9.42,-0.11,-11.38
# 17,TAEE11,31.12,0.32,31.12,28.27,0.11,-9.16
# 18,TRIS3,82.7,0.85,16.54,14.31,-0.49,-13.48
# 19,WEGE3,5012.78,51.24,45.57,68.16,13.83,49.57
#
#
