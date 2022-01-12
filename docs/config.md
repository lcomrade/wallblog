# Configuration
## /etc/wallblog/config.json
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

## /var/lib/wallblog/*
- `./header.md` - Page header
- `./footer.md` - Page footer
- `./style.css` - Page CSS style
- `./favicon.ico` - site icon
- `**/index.md` - works like `index.html`
- `**.md` - convert to HTML
