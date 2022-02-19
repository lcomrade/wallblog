[![Go report](https://goreportcard.com/badge/github.com/lcomrade/wallblog)](https://goreportcard.com/report/github.com/lcomrade/wallblog)
[![Docker Hub Pulls](https://img.shields.io/docker/pulls/lcomrade/wallblog)](https://hub.docker.com/r/lcomrade/wallblog)
[![Release](https://img.shields.io/github/v/release/lcomrade/wallblog)](https://github.com/lcomrade/wallblog/releases/latest)
[![License](https://img.shields.io/github/license/lcomrade/wallblog)](LICENSE)

**WallBlog** - lightweight blogging engine with markdown support.

Features:
- Page formatting in markdown and htmlp
- Custom page header and footer (in markdown)
- Custom error pages
- Custom CSS style
- HTTP and HTTPS


## 📦Install binary release
### UNIX-like systems
Example for Linux amd64.
For other OS/arch change the link end.
```
curl --output /usr/local/bin/wallblog -L https://github.com/lcomrade/wallblog/releases/latest/download/wallblog.linux.amd64
chmod 755 /usr/local/bin/wallblog
```

### Docker container
Read more on the [Docker Hub page](https://hub.docker.com/r/lcomrade/wallblog).



## 🔨Build from source
Build deps:
- Git
- GNU Make
- Golang <= 1.15

```bash
git clone https://github.com/lcomrade/wallblog.git
cd ./wallblog
go get github.com/lcomrade/md2html
make
```

You can find result of build in `./dist/` directory.



## ⚙️Configuration
### `/etc/wallblog/config.json`
To apply the changes in this file, you must restart the program.

Default:
```json
{
	"HTTP": {
		"Enable": true,
		"Port": ":80"
	},
	"HTTPS": {
		"Enable": false,
		"Port": ":443",
		"Cert": "",
		"Key": ""
	},
	"WebRoot": "/var/lib/wallblog"
}
```

### `/var/lib/wallblog/*`
#### Pages and files
- `**.md` - convert to HTML
- `**.htmlp` - like markdown, but only supports HTML tags
- `**/index.md` - works like `index.html`
- `**` - other files

#### Custom page design
- `./header.htmlp` or `./header.md` - page header
- `./footer.htmlp` or `./footer.md` - page footer
- `./style.css` - page CSS style
- `./favicon.ico` - site icon

#### Custom error pages
- `./error.css` - error page CSS style
- `./404.htmlp` or `./404.md`
- `./500_permission_denied.htmlp` or `./500_permission_denied.md`
- `./500_file_read_timeout.htmlp` or `./500_file_read_timeout.md`
- `./500_unknown.htmlp` or `./500_unknown.md`


### Markdown and HTMLP formats
`**.md` files - normal Markdown with support for HTML tags.

`**.htmlp` files - only HTML tags are supported.
Put there what is usually between `<body></body>`.


## How server renders pages?
This is the templates by which the server creates pages.

### Page
```html
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset='utf-8'>
		<link rel='stylesheet' type='text/css' href='/style.css'>
	</head>
	<body>
		<header> <!-- header.htmlp or header.md file --> </header>
		<article> <!-- Requested file --> </article>
		<footer> <!-- footer.htmlp or footer.md file --> </footer>
	</body>
</html>
```

### Error page
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



## 📑Documentation
- [Changelog](CHANGELOG.md)
- [MD2HTML: Markdown Syntax Guide](https://github.com/lcomrade/md2html/blob/main/docs/syntax_guide.md)
