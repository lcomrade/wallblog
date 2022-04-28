[![Go report](https://goreportcard.com/badge/github.com/lcomrade/wallblog)](https://goreportcard.com/report/github.com/lcomrade/wallblog)
[![Docker Hub Pulls](https://img.shields.io/docker/pulls/lcomrade/wallblog)](https://hub.docker.com/r/lcomrade/wallblog)
[![Release](https://img.shields.io/github/v/release/lcomrade/wallblog)](https://github.com/lcomrade/wallblog/releases/latest)
[![License](https://img.shields.io/github/license/lcomrade/wallblog)](LICENSE)

**WallBlog** - lightweight blogging engine with markdown support.

Features:
- Page formatting in markdown and html
- Custom page header and footer (in markdown)
- Custom CSS style
- Automatic sitemap generation
- Custom error pages
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
make
```

You can find result of build in `./dist/` directory.



## ⚙️Configuration
### `/etc/wallblog/config.json`
To apply the changes in this file, you must restart the program.

Default:
```json
{
	"WebRoot": "/var/lib/wallblog",
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
	"Overwrite": {
		"Host": "",
		"Protocol": ""
	},
	"SiteMap": {
		"Enable": true,
		"URL": "/sitemap.xml",
		"SkipHidden": true
	},
	"Page": {
		"AutoTitle": {
			"Enable": true,
			"Prefix": "",
			"Sufix": ""
		},
		"AddToHead": [],
		"EnableBuiltInVars": true
	}
}
```

Read more about:
- [page head](docs/page_head.md)
- [site map](docs/sitemap.md)


### `/var/lib/wallblog/*`
`**.md` files - normal Markdown with support for HTML tags.

`**.htmlp` files - only HTML tags are supported.
Put there what is usually between `<body></body>`.

`**` files - serve according to their MIME type.

#### Pages and files
- `**/index.htmlp` or `**/index.md` - works like `index.html`

#### Custom page design
- `./article_start.htmlp` or `./article_start.md` - custom beginning of the article
- `./article_end.htmlp` or `./article_end.md` - custom end of article
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



## 📑Documentation
- [Article/Error pages](docs/page.md)
- [Config: Page Head](docs/page_head.md)
- [Config: Site Map](docs/sitemap.md)
- [Config: Reverse proxy](docs/reverse_proxy.md)
- [FAQ](docs/faq.md)
- [Changelog](CHANGELOG.md)
- [MD2HTML: Markdown Syntax Guide](https://github.com/lcomrade/md2html/blob/main/docs/syntax_guide.md)
