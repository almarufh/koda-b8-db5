CREATE TABLE "contacts" (
    "Id" bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "First_name" VARCHAR(40) NOT NULL,
    "Last_name" VARCHAR(40) NOT NULL,
    "Email" VARCHAR(40) NOT NULL,
    "Phone" VARCHAR(13),
    "CreatedAt" TIMESTAMP DEFAULT NOW(),
    "UpdatedAt" TIMESTAMP DEFAULT NOW()
);

INSERT INTO "contacts" ("First_name", "Last_name", "Email", "Phone") VALUES
('Almaruf', 'Hidayat', 'almarufhidayat99@gmail.com', '082393468568');

SELECT * FROM "contacts";
UPDATE "contacts" SET "Email" = 'raflis@mail.com' WHERE "Id" = 2;

DELETE FROM "contacts" WHERE "Id" = 4;

DROP TABLE contacts;
