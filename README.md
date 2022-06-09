# Курсовая работа по базам данных

### Технологии:
- `golang + gin framework`
- `react.js`
- `mysql`

### Запуск проекта:

- Установить зависимости: `cd frontend && npm install`
- Создать базу данных с помощью скриптов в папке `sql`
- Запуск frontend-сервера происходит в папке `frontend` командой `npm start`
- Конфиг backend-сервера находится в `backend/src/config/config.go`
- Запуск backend-сервера происходит в папке `backend` командой `go build && ./film-network`

### Структура проекта:
```
.
├── backend
│        ├── go.mod
│        ├── go.sum
│        ├── main.go
│ └── src
│     ├── config
│     │ └── config.go
│     ├── database
│     │ ├── categories.go
│     │ ├── comment.go
│     │ ├── db.go
│     │ ├── film.go
│     │ ├── person.go
│     │ ├── playlist.go
│     │ ├── rating.go
│     │ └── user.go
│     ├── images
│     │ ├── 1.png
│     │ ├── 2.png
│     │ ├── 3.png
│     │ ├── 4.png
│     │ ├── 5.png
│     │ ├── 6.png
│     │ ├── poster.png
│     │ └── unknown.jpg
│     └── server
│         ├── auth.go
│         ├── categories.go
│         ├── comments.go
│         ├── films.go
│         ├── jwt.go
│         ├── people.go
│         ├── playlist.go
│         ├── rating.go
│         ├── server.go
│         └── user.go
├── frontend
│ ├── package.json
│ ├── package-lock.json
│ ├── public
│ │ ├── favicon.ico
│ │ ├── index.html
│ │ ├── logo192.png
│ │ ├── logo512.png
│ │ ├── manifest.json
│ │ └── robots.txt
│ ├── README.md
│ └── src
│     ├── App.js
│     ├── components
│     │ ├── Auth
│     │ │ └── AuthForm.js
│     │ ├── Comments
│     │ │ ├── Comments.js
│     │ │ ├── PostForm.js
│     │ │ ├── Post.js
│     │ │ └── PostsList.js
│     │ ├── Errors
│     │ │ └── Errors.js
│     │ ├── Film
│     │ │ ├── Film.js
│     │ │ ├── FilmRow.js
│     │ │ ├── FilmsList.js
│     │ │ ├── PeopleList.js
│     │ │ └── PersonRow.js
│     │ ├── FilmSearch
│     │ │ ├── FilmCard.js
│     │ │ ├── FilmSearchContainer.js
│     │ │ └── FilmSearch.js
│     │ ├── Home
│     │ │ └── Home.js
│     │ ├── Layout
│     │ │ ├── Layout.css
│     │ │ ├── Layout.js
│     │ │ └── NavigationBar.js
│     │ ├── PeopleSearch
│     │ │ ├── PeopleSearchContainer.js
│     │ │ ├── PersonCard.js
│     │ │ └── PersonSearch.js
│     │ ├── Person
│     │ │ ├── Person.js
│     │ │ ├── RoleRow.js
│     │ │ └── RolesList.js
│     │ ├── Playlist
│     │ │ ├── Playlist.js
│     │ │ ├── PlaylistRow.js
│     │ │ └── PlaylistsList.js
│     │ ├── PlaylistCreation
│     │ │ └── PlaylistCreation.js
│     │ ├── Rate
│     │ │ └── Rate.js
│     │ └── User
│     │     └── User.js
│     ├── db
│     │ └── auth-context.js
│     ├── index.css
│     ├── index.js
│     ├── pages
│     │ ├── AuthPage.js
│     │ ├── FilmPage.js
│     │ ├── FilmSearchPage.js
│     │ ├── HomePage.js
│     │ ├── PersonPage.js
│     │ ├── PersonSearchPage.js
│     │ ├── PlaylistCreationPage.js
│     │ ├── PlaylistPage.js
│     │ └── UserPage.js
│     └── reportWebVitals.js
├── README.md
└── sql
    ├── db_creation.sql
    └── fill_db.sql
```