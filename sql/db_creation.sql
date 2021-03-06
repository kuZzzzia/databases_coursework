CREATE DATABASE IF NOT EXISTS Film_Rec_System;

USE Film_Rec_System;

CREATE TABLE Film_Rec_System.Person (
    `PersonID` int AUTO_INCREMENT PRIMARY KEY,
    `FullName` text NOT NULL,
    `AlternativeName` text NULL,
    `Photo` varchar(100) NOT NULL DEFAULT 'images/unknown.jpg',
    `DateOfBirth` date NULL
);

CREATE TABLE Film_Rec_System.Film (
    `FilmID` int AUTO_INCREMENT PRIMARY KEY,
    `FullName` text NOT NULL,
    `AlternativeName` text NULL,
    `Poster` varchar(100) NOT NULL DEFAULT 'images/poster.png',
    `Description` text NULL,
    `Duration` int NULL,
    `ProductionYear` numeric(4) NULL CHECK (`ProductionYear` >= 1895),
    `PersonID` int,
    FOREIGN KEY (`PersonID`) REFERENCES Film_Rec_System.Person(`PersonID`)
);

CREATE TABLE Film_Rec_System.Role (
    `FilmID` int,
    `PersonID` int,
    `CharacterName` text NULL CHECK ( LENGTH(`CharacterName`) > 0 ),
    PRIMARY KEY (`FilmID`, `PersonID`),
    FOREIGN KEY (`FilmID`) REFERENCES Film_Rec_System.Film(`FilmID`)
        ON DELETE CASCADE,
    FOREIGN KEY (`PersonID`) REFERENCES Film_Rec_System.Person(`PersonID`)
);

CREATE TABLE Film_Rec_System.Genre (
    `GenreID` int AUTO_INCREMENT PRIMARY KEY,
    `GenreName` varchar(30) NOT NULL UNIQUE CHECK ( LENGTH(`GenreName`) > 0)
);

CREATE TABLE Film_Rec_System.Country (
    `CountryID` int AUTO_INCREMENT PRIMARY KEY,
    `CountryName` varchar(50) NOT NULL UNIQUE CHECK ( LENGTH(`CountryName`) > 0)
);

CREATE TABLE Film_Rec_System.Genre_Film_INT (
    `FilmID` int,
    `GenreID` int,
    PRIMARY KEY (`FilmID`, `GenreID`),
    FOREIGN KEY (`FilmID`) REFERENCES Film_Rec_System.Film(`FilmID`)
        ON DELETE CASCADE,
    FOREIGN KEY (`GenreID`) REFERENCES Film_Rec_System.Genre(`GenreID`)
);

CREATE TABLE Film_Rec_System.Country_Film_INT (
    `FilmID` int,
    `CountryID` int,
    PRIMARY KEY (`FilmID`, `CountryID`),
    FOREIGN KEY (`FilmID`) REFERENCES Film_Rec_System.Film(`FilmID`)
        ON DELETE CASCADE,
    FOREIGN KEY (`CountryID`) REFERENCES Film_Rec_System.Country(`CountryID`)
);

CREATE TABLE Film_Rec_System.User (
    `UserID` int AUTO_INCREMENT PRIMARY KEY,
    `Username` varchar(63) NOT NULL UNIQUE,
    `Password` blob NOT NULL,
    `Hash` blob NOT NULL
);

CREATE TABLE Film_Rec_System.Comment (
    `CommentID` int AUTO_INCREMENT PRIMARY KEY,
    `Date` date NOT NULL,
    `Review` text NOT NULL,
    `UserID` int,
    `FilmID` int,
    FOREIGN KEY (`UserID`) REFERENCES Film_Rec_System.User(`UserID`)
        ON DELETE SET NULL,
    FOREIGN KEY (`FilmID`) REFERENCES Film_Rec_System.Film(`FilmID`)
        ON DELETE CASCADE
);

CREATE TABLE Film_Rec_System.FilmRating (
    `UserID` int,
    `FilmID` int,
    `Rating` bool NOT NULL,
    PRIMARY KEY (`UserID`, `FilmID`),
    FOREIGN KEY (`UserID`) REFERENCES Film_Rec_System.User(`UserID`)
        ON DELETE CASCADE,
    FOREIGN KEY (`FilmID`) REFERENCES Film_Rec_System.Film(`FilmID`)
        ON DELETE CASCADE
);

CREATE TABLE Film_Rec_System.Playlist (
    `PlaylistID` int AUTO_INCREMENT PRIMARY KEY,
    `PlaylistTitle` varchar(100) NOT NULL,
    `Description` text NULL,
    `UserID` int,
    FOREIGN KEY (`UserID`) REFERENCES Film_Rec_System.User(`UserID`)
        ON DELETE SET NULL
);

CREATE TABLE Film_Rec_System.Playlist_Film_INT (
    `FilmID` int,
    `PlaylistID` int,
    PRIMARY KEY (`FilmID`, `PlaylistID`),
    FOREIGN KEY (`FilmID`) REFERENCES Film_Rec_System.Film(`FilmID`)
        ON DELETE CASCADE,
    FOREIGN KEY (`PlaylistID`) REFERENCES Film_Rec_System.Playlist(`PlaylistID`)
        ON DELETE CASCADE
);

CREATE TABLE Film_Rec_System.PlaylistRating (
    `UserID` int,
    `PlaylistID` int,
    `Rating` bool NOT NULL,
    PRIMARY KEY (`UserID`, `PlaylistID`),
    FOREIGN KEY (`UserID`) REFERENCES Film_Rec_System.User(`UserID`)
        ON DELETE CASCADE,
    FOREIGN KEY (`PlaylistID`) REFERENCES Film_Rec_System.Playlist(`PlaylistID`)
        ON DELETE CASCADE
);

CREATE VIEW Film_With_Director AS
    SELECT f.FilmID, f.FullName, f.AlternativeName, f.Poster, f.`Description`, f.Duration, f.ProductionYear, f.PersonID, p.FullName AS PersonName FROM Film AS f LEFT JOIN Person AS p on f.PersonID = p.PersonID;

CREATE VIEW Film_Cast AS
    SELECT f.FilmID, f.FullName AS FilmName, f.ProductionYear, r.CharacterName, r.PersonID, p.FullName FROM Role AS r LEFT JOIN Person AS p ON r.PersonID = p.PersonID LEFT JOIN Film AS f on r.FilmID = f.FilmID;

CREATE VIEW Film_Comments_With_Users AS
    SELECT c.FilmID, c.CommentID, c.Review, c.Date, c.UserID, u.Username FROM Comment AS c LEFT JOIN User AS u ON u.UserID = c.UserID;

CREATE VIEW Film_Genres AS
    SELECT inter.FilmID, g.GenreName FROM Genre_Film_INT AS inter LEFT JOIN Genre AS g on inter.GenreID = g.GenreID;

CREATE VIEW Film_Countries AS
    SELECT inter.FilmID, c.CountryName FROM Country_Film_INT AS inter LEFT JOIN Country AS c on inter.CountryID = c.CountryID;

CREATE VIEW Playlists_For_Film AS
    SELECT inter.FilmID, p.PlaylistID, p.PlaylistTitle FROM Playlist_Film_INT as inter LEFT JOIN Playlist AS p on inter.PlaylistID = p.PlaylistID;

CREATE VIEW Playlist_With_Username AS
    SELECT  p.PlaylistID, p.PlaylistTitle, p.`Description`, u.Username FROM Playlist AS p LEFT JOIN User AS u ON p.UserID = u.UserID;

SET GLOBAL log_bin_trust_function_creators = 1;

CREATE FUNCTION getFilmRating(id int)
    RETURNS int
    BEGIN
        DECLARE likeAmount float;
        DECLARE total float;
        DECLARE ret int DEFAULT -1;
        CREATE TEMPORARY TABLE TempTableForFilmRating(Rate BOOL);
        INSERT INTO TempTableForFilmRating SELECT Rating FROM FilmRating WHERE FilmID = id;
        SELECT Count(*) INTO likeAmount FROM TempTableForFilmRating WHERE Rate = TRUE;
        SELECT Count(*) INTO total FROM TempTableForFilmRating;
        DROP TEMPORARY TABLE TempTableForFilmRating;
        IF total > 0 THEN
            SET ret = likeAmount / total * 100;
        END IF;
        RETURN ret;
    END;

CREATE FUNCTION getPlaylistRating(id int)
    RETURNS int
BEGIN
    DECLARE likeAmount float;
    DECLARE total float;
    DECLARE ret int DEFAULT -1;
    CREATE TEMPORARY TABLE TempTableForPlaylistRating(Rate BOOL);
    INSERT INTO TempTableForPlaylistRating SELECT Rating FROM PlaylistRating WHERE PlaylistID = id;
    SELECT Count(*) INTO likeAmount FROM TempTableForPlaylistRating WHERE Rate = TRUE;
    SELECT Count(*) INTO total FROM TempTableForPlaylistRating;
    DROP TEMPORARY TABLE TempTableForPlaylistRating;
    IF total > 0 THEN
        SET ret = likeAmount / total * 100;
    END IF;
    RETURN ret;
END;

