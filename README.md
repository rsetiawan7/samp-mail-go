# samp-mail-go

[![sampctl](https://img.shields.io/badge/sampctl-samp--mail--go-2f2f2f.svg?style=for-the-badge)](https://github.com/rsetiawan7/samp-mail-go)

This library was heavily inspired by [SAMPMailJS](https://github.com/bruxo00/SAMPMailJS) and initiated because I'm learning Go, make this library run as simple as configure `.env` file, and run it!

Any feedbacks and pull requests are welcome! I really appreciate it.

<!--
Short description of your library, why it's useful, some examples, pictures or
videos. Link to your forum release thread too.

Remember: You can use "forumfmt" to convert this readme to forum BBCode!

What the sections below should be used for:

`## Installation`: Leave this section un-edited unless you have some specific
additional installation procedure.

`## Testing`: Whether your library is tested with a simple `main()` and `print`,
unit-tested, or demonstrated via prompting the player to connect, you should
include some basic information for users to try out your code in some way.

And finally, maintaining your version number`:

* Follow [Semantic Versioning](https://semver.org/)
* When you release a new version, update `VERSION` and `git tag` it
* Versioning is important for sampctl to use the version control features

Happy Pawning!
-->

## Installation

Simply install to your project:

```bash
sampctl package install rsetiawan7/samp-mail-go
```

Include in your code and begin using the library:

```pawn
#include <samp-mail-go>
```

## Usage

<!--
Write your code documentation or examples here. If your library is documented in
the source code, direct users there. If not, list your API and describe it well
in this section. If your library is passive and has no API, simply omit this
section.
-->
#### --- Make sure you're using Go 1.19

Install dependencies
```
go mod tidy
```

Build your mail server

```
go build
```

Run it

```
./samp-mail-go
```

In your SA-MP code, use it

```pawn
#include <samp-mail-go>

// Send e-mail
SendEmail("Your Server", "your-email@example.com", "Hello", "You should receive this e-mail")
```

## To-do

- [ ] CI / CD for Releases build (Windows & Linux)
- [ ] Unit test
