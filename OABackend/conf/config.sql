CREATE TABLE `recipes` (
                           `Id` int NOT NULL AUTO_INCREMENT,
                           `Name` text,
                           `Ingredients` text,
                           `Rating` int DEFAULT NULL,
                           PRIMARY KEY (`Id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

