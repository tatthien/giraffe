# Giraffe: An opinionated static site generator

I want to start 2022 by writing a static site generator that is used by my personal website (https://thien.dev)

About the name "Giraffe" it's because of my son. He really like giraffes.

## How to write posts?

All content should be localed in `/content/{post-type}` folder, in markdown format. Each posts should have frontmatter and content:

```
---
title: <string>
date: YYYY-MM-DD
tags: <string>, <string>
description: <string>
---

<your-content>
```

Each folder inside `content` is corresponding to a post type. For example, `/content/posts/` is a folder that contains all `posts`.