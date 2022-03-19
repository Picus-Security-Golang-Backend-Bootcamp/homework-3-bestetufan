# homework-3-bestetufan
homework-3-bestetufan created by GitHub Classroom

Beste Tufan

1-) docker-compose up -d
2-) docker exec -it -u postgres db psql
3-) CREATE DATABASE bookstore;
5-) run main.go list
    5.a-) [InsertSampleData] -> optional
    5.b-) [AutoMigration] -> (books, author)
6-) to test from db:
    6.a-) connect db -> \c bookstore
    6.b-) SELECT * FROM book;
7-) teardown -> docker-compose down