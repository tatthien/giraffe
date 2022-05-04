# Giraffe: An opinionated static site generator

I want to start 2022 by writing a static site generator that is used by my personal website (https://thien.dev)

About the name "Giraffe" it's because of my son. He really like giraffes.

## Install

```
go get github.com/tatthien/giraffe
```

For non-Go users.

```
curl -sf https://gobinaries.com/tatthien/giraffe | sh
```

## Usage

- `giraffe`: Build everything into `dist` folder.
- `giraffe serve`: Serving the site at `localhost:3333` for preview. It also rebuilds the site when a file in `contents` and `theme` changes.
- `giraffe version`: Show the latest version.

## How to write posts?

All content should be localed in `/content/{post-type}` folder, in markdown format. Each posts should have frontmatter and content:

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
