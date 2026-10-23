# go-statistics.github.io

The go-statistics landing page (Hugo), served at <https://go-statistics.github.io/>.

The layout is the go-* family template (go-fileshare.github.io's, through
`logos/landing-kit`, as go-numerics and go-dataframe use it):
`layouts/index.html`, the shared stylesheet `layouts/partials/styles.html`,
and `[params.brand]` in `hugo.toml`, read from the logo's gradient.

The figure is real: `figure/main.go` draws 10 000 gamma variates with
go-statistics/stats, fits them and tests the fit, and draws all of it with
go-plot into `static/img/figure.svg`; the page's terminal shows that file,
read at build time. The `ci` workflow draws the figure again and fails if a
byte differs. (The figure is the same drawn on amd64 and on arm64.)

```sh
cd figure && go run . > ../static/img/figure.svg
hugo server
```

| Workflow | What |
| --- | --- |
| `hugo.yml` | builds on pull requests; builds and deploys to Pages on `main` |
| `ci.yml` | the site builds and names every repository; the figure is what its program draws |
| `links.yml` | every link and image of the built site resolves |

BSD-3-Clause.

## The design is not here

It comes from the shared Hugo module
[`go-fleettools/landingkit`](https://github.com/go-fleettools/landingkit),
pinned in `go.mod` (Renovate bumps it); `kit.go` keeps `go mod tidy` from
deleting that requirement. `[params.brand]` in `hugo.toml` is read from the logo by
`landingkit brand`. A file under `layouts/partials/` with the name of one of the
kit's would silently override it.
