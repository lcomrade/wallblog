# Page Head

```json
{
	"Page": {
		"AutoTitle": {
			"Enable": true,
			"Prefix": "",
			"Sufix": ""
		},
		"AddToHead": []
	}
}
```


## `AutoTitle`
Sets the content of `<head><title>....</title></head>`.

If this option is enabled, the header will be created using following pattern:
`Prefix + AutoTitle + Sufix`.
`AutoTitle` is the first level header in Markdown or the first tag `<h1>....</h1>` in HTMLP.


## `AddToHead`
Adds extra lines to `<head>....</head>`.

With this directive you can add custom JavaScript or other metadata.
Example:
```json
{
	"Page": {
		"AddToHead": [
			"<script src='https://example.org/extra.js'></script>",
			"<meta name='robots' content='noindex, nofollow'>"
		]
	}
}
```
