use Film_Rec_System;

INSERT INTO Person(FullName)
VALUES
    ('Хью Джекман'),
    ('Леонардо ДиКаприо'),
    ('Роберт Дауни Младший'),
    ('Мэтью Макконахи');

INSERT INTO Film(FullName)
VALUE ('Волк-с-Уолл-Стрит');

INSERT INTO Role(FilmID, PersonID, CharacterName)
VALUES (1, 2, 'Джордан Белфорт'),
       (1, 4, 'Марк Ханна');

