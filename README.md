# GolangProject
Проект представляет собой веб-приложение, которое позволяет пользователям создавать, просматривать, редактировать и удалять статьи.

- `/api/articles` - Получить список всех статей.
- `/api/articles/{id}` - Получить информацию о статье по её id.
- `/api/articles/add` - Добавить новую статью.
- `/api/articles/edit/{id}` - Редактировать информацию о статье по её id.
- `/api/articles/delete/{id}` - Удалить статью по её id.
# DB Structure

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
