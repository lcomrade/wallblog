# Changelog
Semantic versioning is used (https://semver.org/).

## v1.3.2
- Added `article_start` and `article_end`
- Added build in variables (`{{URL.Full}}` and `{{URL.Path}}`)
- Disallow access to config files. Now return 403 HTTP code.
- Fix: site URL detection
- Update Golang 1.15.15 => 1.18.1

## v1.3.1
- Fix: HTML tags in autotitle

## v1.3.0
- Added automatic page title generation
- Allowed to add additional tags to `<head>...</head>`

## v1.2.0
- Added automatic sitemap generation
- Added reverse proxy support
- Allow use `index.htmlp` as index page
- Documented sitemap, reverse proxy and FAQ

## v1.1.0
- Added custom error pages
- Added support for HTMLP format
- Page rendering is documented

## v1.0.0
This is the first release with only basic functionality implemented here:
- Page formatting in Markdown
- Custom page header and footer (in Markdown)
- Custom CSS style
- HTTP and HTTPS
