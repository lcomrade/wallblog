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


## Built-in variables
| Varible        | Description                              | Example                             |
| -------------- | -----------------------------------------| ----------------------------------- |
| `{{URL.Path}}` | Path without a domain name and protocol. | `/kb/my_page.md`                    |
| `{{URL.Full}}` | Full path with domain name and protocol. | `https://example.org/kb/my_page.md` |

If you want to escape built in variable, you do this: `\{{\{MY_VAR\}\}`.
