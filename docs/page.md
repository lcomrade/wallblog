# Page
## Page render template
Article page:
```html
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset='utf-8'>
		<link rel='stylesheet' type='text/css' href='/style.css'>
	</head>
	<body>
		<header> <!-- header.htmlp or header.md file --> </header>
		<article>
			<!-- article_start.htmlp or article_start.md file -->
			<!-- Requested file -->
			<!-- article_end.htmlp or article_end.md file -->
		</article>
		<footer> <!-- footer.htmlp or footer.md file --> </footer>
	</body>
</html>
```

Error page:
```html
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset='utf-8'>
		<link rel='stylesheet' type='text/css' href='/error.css'>
	</head>
	<body>
		<!-- .htmlp or .md file corresponding to the error code -->
	</body>
</html>
```


## Template mode
| Varible                              | Example                                                              |
| ------------------------------------ | -------------------------------------------------------------------- |
| `{{.Request.ClientIP}}`              | `172.17.0.1`                                                         |
| `{{.Request.Method}}`                | `GET`                                                                |
| `{{.Request.URL.Protocol}}`          | `http`                                                               |
| `{{.Request.URL.Path}}`              | `/kb/my_page.md`                                                     |
| `{{.Request.URL.RawQuery}}`          | `v=11&1=99`                                                          |
| `{{.Request.URL.Full}}`              | `https://example.org/kb/my_page.md`                                  |
| `{{.Request.Header.AcceptLanguage}}` | `en`                                                                 |
| `{{.Request.Header.Host}}`           | `example.org`                                                        |
| `{{.Request.Header.Referer}}`        | `https://yandex.ru`                                                  |
| `{{.Request.Header.UserAgent}}`      | `Mozilla/5.0 (Windows NT 10.0; rv:91.0) Gecko/20100101 Firefox/91.0` |

If you want to escape built in variable, you do this: `\{\{MY_VAR\}\}`.
