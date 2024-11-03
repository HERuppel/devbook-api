INSERT INTO users (name, nick, email, password)
VALUES
('Usuário 1', 'usuario1', 'usuario1@gmail.com', '$2a$10$hh561NcllHjDyF21bXSVS.G1XrBcqHjG2WNroG95sZlqTcH84Bhe2'),
('Usuário 2', 'usuario2', 'usuario2@gmail.com', '$2a$10$hh561NcllHjDyF21bXSVS.G1XrBcqHjG2WNroG95sZlqTcH84Bhe2'),
('Usuário 3', 'usuario3', 'usuario3@gmail.com', '$2a$10$hh561NcllHjDyF21bXSVS.G1XrBcqHjG2WNroG95sZlqTcH84Bhe2'),
('Usuário 4', 'usuario4', 'usuario4@gmail.com', '$2a$10$hh561NcllHjDyF21bXSVS.G1XrBcqHjG2WNroG95sZlqTcH84Bhe2');

INSERT INTO followers(userId, followerId)
VALUES
(1, 2),
(3, 1),
(1, 3);


