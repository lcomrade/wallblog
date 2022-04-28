# Frequently Asked Questions
## Why is there no support for Windows and OSX?
There are several reasons for this:
1. I do not support proprietary software. And I'm not going to help users who use it.
2. Windows and OSX are not server OSs.
3. Windows is not POSIX compliant.

You can try to build a program for these operating systems.
But it may not work at all or it may not work as intended.


## How server renders pages?
This is the templates by which the server creates pages.

Page:
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
