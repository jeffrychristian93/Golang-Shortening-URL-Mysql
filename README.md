# Golang-Shortening-URL-Mysql

Shortening url example using go language.

Installation
1. Download main.go
2. Edit your database connection on this code:
    - sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db-redirects")
    - sql.Open("mysql", "username:password@tcp(yourIpAddress:port)/yourDataabse")
3. Run using this command -> go run main.go

If you got some error, please download the libraries first.
1. go get github.com/gin-gonic/gin" -> Gin is a web framework written in Go (Golang)
2. go get github.com/go-sql-driver/mysql" -> Mysql driver for Go

______________________________________________________________________________


You can create manual table with this or auto create when program first running.

CREATE TABLE `redirect` (
	`id` int NOT NULL AUTO_INCREMENT,
	`slug` varchar(6) collate utf8mb4_unicode_ci NOT NULL,
	`url` varchar(620) collate utf8mb4_unicode_ci NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='URL shortener Table';

______________________________________________________________________________

To create short url you can try with postman:
- Method : POST
- Request URL : http://localhost:8000/create
- Param : 
  - key : url
  - value : http://example.com/this-is-a-very-long-url-bla-bla-bla

result:
{
  "message": "201 Created",
  "url": "Location: http://domain.com/5fbbd6"
}

______________________________________________________________________________

Then you can try GET request:
- Method : GET
- Request URL : http://localhost:8000/5fbbd6

result:
{
  "message": "301 Found",
  "url": "Location: http://example.com/this-is-a-very-long-url-bla-bla-bla"
}

Thank you.
