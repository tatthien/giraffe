# Giraffe: An opinionated static site generator

![build badge](https://img.shields.io/github/workflow/status/tatthien/giraffe/Test)

I want to start 2022 by writing a static site generator that is used by my personal website (https://thien.dev)

About the name "Giraffe" it's because of my son. He really like giraffes.

## Install

Via `go install` with go 1.16 or higher:

```
go install github.com/tatthien/giraffe@latest
```

Or download binary files from [Releases](https://github.com/tatthien/giraffe/releases) page  

## Commands

**`giraffe help`**: Print the usage guide.

```bash
$ giraffe help

Giraffe: An opinionated static site generator

Usage: giraffe [command] [arguments]

Available commands:
        serve           start a server, watch files changed and rebuild
        build           generate static files
        new [path]      create a new markdown file
        version         print the cli version
```

**`giraffe`**: Build everything into `dist` folder.

```bash
$ giraffe

CONTENT         TOTAL
Tags            15
Post types      3
Pages           18
Build time      105.455465ms
```

**`giraffe serve`**: Serving the site at `localhost:3333` for preview (:3333 is the default port, you can change it in `config.yaml`). It also rebuilds the site when a file in `contents` and `theme` changes.

**`giraffe new [path]`**: Quickly create a new markdown file insert the `content` directory.

```bash
# This command will create a new file inside /content/posts and automatically set the date.
$ giraffe new posts/welcome-to-girrafe.md
```

**`giraffe version`**: Show the latest version.

Sample output:

```bash
$ giraffe version

giraffe version: v0.9.0
```

## Configuration

There are some options that you can change in `config.yaml`.

```yaml
baseURL: <site base url>
title: <site title>
description: <site description>
port: <server port>
```

## How to write posts?

All content should be localed in `/content/>post-type>` folder, in markdown format. Each post should have frontmatter and content:

```
---
title: <string>
date: YYYY-MM-DD
draft: <boolean> # This is the item status. Set ` if you don't want to publish the item.
tags: <string>, <string>
---

<your-content>
```

Each folder inside `content` is corresponding to a post type. For example:

- The folder `/content/posts/` contains all items which it post type is `posts`
- The folder `/content/pages` contains all items which it post type is `pages`
## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2022-present, Thien Nguyen
