INSERT INTO users (name, username, email, password)
values
("Daniel San", "dsan", "dsan@mail.com", "$2a$10$405iDcLOxcrvGFbh67RtV.Zl4XYgsZGidJKbjFXlC7GsxPglYriQK"),
("Han Solo", "hsolo", "hsolo@mail.com", "$2a$10$405iDcLOxcrvGFbh67RtV.Zl4XYgsZGidJKbjFXlC7GsxPglYriQK"),
("Arthur Dent", "adent", "adent@mail.com", "$2a$10$523phhLvXG/WpYmKHCkI4.AJkU3u0jXZAmUjMp0Tu2XhZ.6/pm3YW"),
("Bruce Lee", "blee", "blee@mail.com", "$2a$10$523phhLvXG/WpYmKHCkI4.AJkU3u0jXZAmUjMp0Tu2XhZ.6/pm3YW"),
("Bob Burnquist", "bburnquist", "bburnquist@mail.com", "$2a$10$523phhLvXG/WpYmKHCkI4.AJkU3u0jXZAmUjMp0Tu2XhZ.6/pm3YW"),
("Darth Vader", "dvader", "dvader@mail.com", "$2a$10$405iDcLOxcrvGFbh67RtV.Zl4XYgsZGidJKbjFXlC7GsxPglYriQK");

INSERT INTO followers(user_id, follower_id)
values
(1, 2), (1, 3), (1, 4), (1, 5), (1, 6),
(2, 1), (2, 3), (2, 4), (2, 5), 
(3, 1), (3, 2), (3, 4), (3, 5), (3, 6),
(4, 1), (4, 2), (4, 3), (4, 5), (4, 6),
(5, 1), (5, 2), (5, 3), (5, 4),
(6, 3), (6, 4);

INSERT INTO posts(title, content, author_id)
values
("Cobra Kai is dead!", "A post for testing by Daniel San", 1),
("I miss Chewie...", "A post for testing by Han Solo", 2),
("The real question is...", "Actually I just know that the answer is 42. Testing by Arthur Dent.", 3),
("That time I kicked Chuck Norris ass...", "A post for testing by Bruce Lee", 4),
("Skate in Brazil and its references", "A post for testing by Bob Burnquist", 5),
("Come to the dark side!", "A post for testing by Darth Vader", 6);