CREATE TABLE users(
    Id  SERIAL PRIMARY KEY,
	Username VARCHAR(255) NOT NULL,
	PasswordHash VARCHAR(255) NOT NULL,
	AvatarUrl VARCHAR(255),
	SumExperienceINT DEFAULT 0,
	AmountExperienceToLvl INT,
	Lvl INT DEFAULT 0
);

CREATE TABLE quests(
    Id SERIAL PRIMARY KEY, 
    Title VARCHAR(255) NOT NULL,
    Description TEXT NOT NULL,
    Dificulty VARCHAR(255) NOT NULL,
    Completed BOOLEAN DEFAULT FALSE,
    User_Id INT REFERENCES users(id)
);
