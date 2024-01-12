begin;

--creating tables
create table family (
    role varchar(30) not null,
    age integer not null,
    name varchar(30) not null
);

create table appliances (
    name varchar(30) not null,
    color varchar(30) not null
);

create table furniture (
    name varchar(30) not null,
    material varchar(30) not null,
    color varchar(30) not null
);

create table sizes (
    name varchar(30) not null,
    area real not null
);

--fill tables
insert into family (role, age, name) values
    ('мама', 34, 'Лариса'),
    ('папа', 71, 'Олег'),
    ('мама', 17, 'Саша'),
    ('мама', 3, 'Моцарт'),
    ('мама', 10, 'Череп');

insert into appliances (name, color) values
    ('холодильник', 'белый'),
    ('духовка', 'черный'),
    ('чайник', 'синий'),
    ('фен', 'серый'),
    ('блендер', 'розовый'),
    ('утюг', 'зеленый');

insert into furniture values
    ('кровать', 'дерево', 'коралловый'),
    ('диван', 'кожа', 'фиолетовый'),
    ('шкаф', 'фанера', 'желтый'),
    ('стол', 'стекло', 'белый'),
    ('кресло', 'ткань', 'бежевый');

insert into sizes values
    ('спальня', 25),
    ('кухня', 11.5),
    ('ванная комната', 10),
    ('гостиная', 22);

commit;