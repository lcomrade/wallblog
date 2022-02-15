[![Go report](https://goreportcard.com/badge/github.com/lcomrade/wallblog)](https://goreportcard.com/report/github.com/lcomrade/wallblog)
[![Docker Hub Pulls](https://img.shields.io/docker/pulls/lcomrade/wallblog)](https://hub.docker.com/r/lcomrade/wallblog)
[![Release](https://img.shields.io/github/v/release/lcomrade/wallblog)](https://github.com/lcomrade/wallblog/releases/latest)
[![License](https://img.shields.io/github/license/lcomrade/wallblog)](LICENSE)

**wallblog** - lightweight blogging engine with markdown support.


## Install binary release
### UNIX-like systems
Example for Linux amd64.
For other OS/arch change the link end.
```
curl --output /usr/local/bin/wallblog -L https://github.com/lcomrade/wallblog/releases/latest/download/wallblog.linux.amd64
chmod 755 /usr/local/bin/wallblog
```

### Docker container
Read more on the [Docker Hub page](https://hub.docker.com/r/lcomrade/wallblog).



## Build from source
Build deps:
- GNU Make
- Golang <= 1.15

```
git clone https://github.com/lcomrade/wallblog.git
cd ./wallblog
make
```

You can find result of build in `./dist/` directory.



## Documentation
- [Configuration Guid](docs/config.md)
- [Markdown Syntax Guide](https://github.com/lcomrade/md2html/blob/main/docs/syntax_guide.md)
- [Changelog](CHANGELOG.md)
