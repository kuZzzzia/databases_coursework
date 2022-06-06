CREATE DATABASE IF NOT EXISTS Film_Rec_System;

USE Film_Rec_System;
# DROP DATABASE Film_Rec_System;


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

CREATE TABLE Film_Rec_System.Discussion (
    `DiscussionID` int AUTO_INCREMENT PRIMARY KEY,
    `Date` date NOT NULL,
    `Review` text NOT NULL,
    `UserID` int,
    `FilmID` int,
    FOREIGN KEY (`UserID`) REFERENCES Film_Rec_System.User(`UserID`)
        ON DELETE SET NULL,
    FOREIGN KEY (`FilmID`) REFERENCES Film_Rec_System.Film(`FilmID`)
        ON DELETE CASCADE
);

CREATE TABLE Film_Rec_System.View (
    `UserID` int,
    `FilmID` int,
    `LastViewDate` date NULL,
    `FilmScore` bool NOT NULL,
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

CREATE TABLE Film_Rec_System.PlaylistScore (
    `UserID` int,
    `PlaylistID` int,
    `Score` bool NOT NULL,
    PRIMARY KEY (`UserID`, `PlaylistID`),
    FOREIGN KEY (`UserID`) REFERENCES Film_Rec_System.User(`UserID`)
        ON DELETE CASCADE,
    FOREIGN KEY (`PlaylistID`) REFERENCES Film_Rec_System.Playlist(`PlaylistID`)
        ON DELETE CASCADE
);

SET GLOBAL log_bin_trust_function_creators = 1;

CREATE FUNCTION getFilmRating(id int)
    RETURNS int
    BEGIN
        DECLARE likeAmount float;
        DECLARE total float;
        DECLARE ret int DEFAULT -1;
        CREATE TEMPORARY TABLE TempTable(Rate BOOL);
        INSERT INTO TempTable SELECT FilmScore FROM View WHERE FilmID = id;
        SELECT Count(*) INTO likeAmount FROM TempTable WHERE Rate = TRUE;
        SELECT Count(*) INTO total FROM TempTable;
        DROP TEMPORARY TABLE TempTable;
        IF total > 0 THEN
            SET ret = likeAmount / total * 100;
        END IF;
        RETURN ret;
    END;

