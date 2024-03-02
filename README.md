# GolangProject
Проект представляет собой веб-приложение, которое позволяет пользователям создавать, просматривать, редактировать и удалять статьи.
1. У нас будет сайт в котором будут публиковаться статьи
2. У статьи будет тема и сам текст
3. Можно будет публиковать статьи сколько угодно
4. Можно будет искать статьи
5. Будет главная страница, контакты, о нас, добавить статью, смена языка
6. Будут категории статьей 
7. Можно будет сортировать статьи

# Структура API
GET /articles - Получить список всех статей
POST /articles - Создать новую статью
GET /articles/{id} - Получить статью по ID
PUT /articles/{id} - Обновить статью по ID
DELETE /articles/{id} - Удалить статью по ID
# DB Structure

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users(id)
);
