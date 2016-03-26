# Feed

An RSS parser written in Go.

## Why?

There's already a few RSS packages/projects written in Go:

* [jteeuwen/go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss/)
* [gorilla/feeds](https://github.com/gorilla/feeds)
* [SlyMarbo/rss](https://github.com/SlyMarbo/rss)

Why did I feel the need to write another?

Well each of these seemed to either be for reading/parsing or generating rss feeds. Essentially, they seemed like they were designed for either consumption or production of RSS feeds but weren't interested in supporting both of these feature sets.

This made me sad.

Why shouldn't we have a single package that could be used for both use cases and cut down on the duplicated code in between?

That still doesn't explain why I didn't try to contribute these features to one of the above projects... Well, to be honest, I just wanted an excuse to play around more with the Go language and to learn more about the RSS/Atom standards!
