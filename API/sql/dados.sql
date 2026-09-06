insert into usuarios(nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$UMpaG1TWwq3Qmp/f2KtxlOktCK0LKT.KUEtpYuQxiwNrH/LF35omi"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$UMpaG1TWwq3Qmp/f2KtxlOktCK0LKT.KUEtpYuQxiwNrH/LF35omi"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$UMpaG1TWwq3Qmp/f2KtxlOktCK0LKT.KUEtpYuQxiwNrH/LF35omi");

insert into seguidores(usuario_id, seguidor_id)
values  
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Publicação teste do 1", 1),
("Publicação do Usuário 2", "Publicação teste do 2", 2),
("Publicação do Usuário 3", "Publicação teste do 3", 3);