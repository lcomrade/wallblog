# Page Head

```json
{
	"Page": {
		"AutoTitle": {
			"Enable": true,
			"Prefix": "",
			"Sufix": ""
		}
	}
}
```

## AutoTitle
Sets the content of `<head><title>....</title></head>`.

If this option is enabled, the header will be created using following pattern:
`Prefix + AutoTitle + Sufix`.
`AutoTitle` is the first level header in Markdown or the first tag `<h1>....</h1>` in HTMLP.
