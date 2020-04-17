# example-dependency-injection

Cara mudah melakukan unit test pada golang. Mocking, dependency injection in Go

Dalam melakukan unit test terdapat salah satu teknik Dependency Injection (DI) simple namun sangat powerful dalam menulis unit test pada golang. Salah satu sarat baik dalam menuliskan unit test yaitu adalah fungsi yang dibuat haruslah independent dan testable.

Goals:
- Diharapkan dapat menulis unit test yang mampu berdiri sendiri.

Kasus Studi:

Pase 1 membuat fungsi sebenarnya

Pase 1.2 Output
nakama@roronoa:~/go/src/github.com/firdasafridi/example-dependency-injection/cmd$ go run main.go
2020/04/17 14:42:13 Afghanistan
2020/04/17 14:42:13 Åland Islands
2020/04/17 14:42:13 Albania
2020/04/17 14:42:13 Algeria
2020/04/17 14:42:13 American Samoa
2020/04/17 14:42:13 Andorra
2020/04/17 14:42:13 Angola
2020/04/17 14:42:13 Anguilla
2020/04/17 14:42:13 Antarctica
2020/04/17 14:42:13 Antigua and Barbuda
...


Pase 2 membuat interface yang menyerupai fungsi sebenarnya
- menulis interface
- menjalankan mock generator

mockgen -destination=./_mocks/mock_repository.go -package=_mocks github.com/firdasafridi/example-dependency-injection/country Repository

Pase 3 Menulis unit test menggunakan gomock (dependency injection)

Conclusion
