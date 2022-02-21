# Joplin Multi-Note Viewer

Use [Joplin's Data API](https://joplinapp.org/api/references/rest_api/) to
view and browse multiple notes at once. Supports title and body search and internal
links, as well as dragging and dropping notes to rearrange them in your view.

Visible note IDs are put into the URL as you go, so if you want to return to a
particular note set, bookmark the generated URL.

## Installing

* Enable the Web Clipper service in Joplin (Tools > Options > Web Clipper > Enable the clipper service).
* Download the latest release binary for your platform.
* Download the sample `config.yml` file from Releases. Then edit it and
  fill in the following information:

```yaml
joplin:
  api_key: <the authorization token from Tools > Options > Web Clipper > Advanced Options>
  url: http://localhost:41184
```

## Running

* Start Joplin.
* Start the viewer from the command line.
* Visit [http://localhost:8181/](http://localhost:8181).

## Using

* The search box can be refocused and brought into view by hitting `s`.
* You can use the up and down arrow keys to navigate autocomplete results,
  and `[Enter]` to open a result.
* Open notes can be rearranged by dragging and dropping.
* Close a note with the Close button.
* Notes will reload every few seconds so if you make changes in Joplin,
  they'll be reflected in the viewer.

## Contributing

This is a hobby project for me to learn Vue 3 and Go and to be able to use my
interlinked notes better, so if I fix a bug, it'll either be because it's
directly affecting me, it's easy to fix, or it's interesting. Pull
requests are preferred over issues if you have the skill, but even if you
don't, someone else might!

## Developing

I built this using Go 1.18 and Node 16.

* Run `dev-setup.sh`. This will make a build of the Vue app and put its
  `dist` folder into `server`. This is so `go:embed` works.
* Start the server with `go run .`.
* Start the client with `npm run dev`. It will proxy API requests to a locally
  running server so the embedded client files will be overridden by the files
  Vite produces.

Add tests if you want. I may get around to it eventually.

## Building & Releasing

* Bump the version. This project uses semantic versioning as best as possible.
* Update the changelog with details on the changes. It uses the
  [Keep a Changelog format](https://raw.githubusercontent.com/olivierlacan/keep-a-changelog/main/CHANGELOG.md).
* Run `build.sh`. It will create the different platform binaries in `build` as
  well as a `config.yml` sample.

## License

MIT License
