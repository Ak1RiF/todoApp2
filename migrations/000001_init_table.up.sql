CREATE TABLE users(
    Id  SERIAL PRIMARY KEY,
	Username VARCHAR(255) NOT NULL,
	PasswordHash VARCHAR(255) NOT NULL,
	AvatarUrl VARCHAR(255),
	SumExperience INT DEFAULT 0,
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

CREATE TABLE pets(
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Rarity VARCHAR(255) NOT NULL
);


CREATE TABLE users_pets(
    Id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    pet_id INT REFERENCES pets(id)
);
