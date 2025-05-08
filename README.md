### QUIZ

Ovo je jednostavna aplikacija kviza napisana u [Go](https://go.dev/) programskom jeziku. Aplikacija učitava pitanja i odgovore iz plain text fajla i pokreće kviz u komandnoj liniji.

Pitanja i odgovori će biti automatski izmešana (shuffle) prilikom svakog pokretanja kviza, tako da redosled neće uvek biti isti.

1. Instalacija
```sh
git clone https://github.com/mstanimirovic/quiz.git
```
ili skinite .zip file.

2. Pokretanje aplikacije
```sh
go run main.go [ime fajla]
```
