# Site Map
Site Map is used by search engines (like Google).
This is a simple XML file whose specification is described at [sitemaps.org](https://www.sitemaps.org/protocol.html).

WallBlog generates a minimal version. Example:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<url>
		<loc>https://example.org/about.md</loc>
	</url>
	<url>
		<loc>https://example.org/download/v1.0.0.md</loc>
	</url>
</urlset>
```

Only files with extensions are indexed:
- `.md`
- `.htmlp`
- `.html`
- `.txt`

All configuration files (e.g. `header.md`, `footer.md`) are excluded from the sitemap.


## Configuration
All settings are in the configuration file.

Site map generation settings:
```json
{
	"SiteMap": {
		"Enable": true,
		"URL": "/sitemap.xml",
		"SkipHidden": true
	}
}
```

- `Enable` - disable automatic sitemap generation
- `URL` - address where the map will be located
- `SkipHidden` - skip files and dirs that begin with `.`


### Reverse proxy
**If you are using a reverse proxy** be sure to configure these two settings.
Without these, the file will be generated with an error.

```json
{
	"Overwrite": {
		"Host": "",
		"Protocol": ""
	}
}
```
