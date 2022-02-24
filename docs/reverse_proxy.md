# Reverse proxy
WallBlog supports working behind a reverse proxy.

```json
{
	"Overwrite": {
		"Host": "",
		"Protocol": ""
	}
}
```

- `Host` - domain name of the host (e.g. `example.org`, `192.168.1.15`)
- `Protocol` - specifies the preferred protocol (`http`/`https`)


## Example
This example shows how to set up Nginx as a reverse proxy.
WallBlog is *http* and only listens to *localhost*,
Nginx is *https* and listens to requests from *external IPs*.


`/etc/nginx/nginx.conf`:
```nginx
events {
	worker_connections 1024;
}

http {
	include /etc/nginx/mime.types;
	include /etc/nginx/conf.d/*.conf;

	error_log /var/log/nginx/error.log warn;
	server_tokens off;

	ssl_protocols TLSv1.2 TLSv1.3;

	proxy_http_version 1.1;

	server {
		server_name example.org;
		listen 443 ssl http2;
		listen [::]:443 ssl http2;
		ssl_certificate /etc/letsencrypt/live/example.org/fullchain.pem;
		ssl_certificate_key /etc/letsencrypt/live/example.org/privkey.pem;
	
		access_log /var/log/nginx/example.org.access.log combined;
		
		location / {
			proxy_pass http://localhost:8000/;
		}
	}
}
```


`/etc/wallblog/config.json`:
```json
{
	"WebRoot": "/var/lib/wallblog",
	"HTTP": {
		"Enable": true,
		"Port": "localhost:8000"
	},
	"HTTPS": {
		"Enable": false
	},
	"Overwrite": {
		"Host": "example.org",
		"Protocol": "https"
	}
}
```
