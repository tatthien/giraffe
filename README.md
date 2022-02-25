# Giraffe: An opinionated static site generator

I want to start 2022 by writing a static site generator that is used by my personal website (https://thien-space.netlify.app)

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

There are only 2 commands:

- `giraffe`: Build everything into `dist` folder.
- `giraffe serve`: Serving the site at `localhost:3333` for preview. It also rebuilds the site when a file in `contents` and `theme` changes.

## How to write posts?

All content should be localed in `/content/{post-type}` folder, in markdown format. Each posts should have frontmatter and content:

```
---
title: <string>
date: YYYY-MM-DD
tags: <string>, <string>
description: <string>
draft: false <boolean> # This is the item status. Set true if you don't want to publish the item.
---

<your-content>
```

Each folder inside `content` is corresponding to a post type. For example, the folder `/content/posts/` contains all `posts`.

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2022-present, Thien Nguyen
