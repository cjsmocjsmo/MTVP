# MTVP

MTVP is a Go-based media catalog and web UI. It scans your configured movie, TV, and video directories, stores metadata in SQLite, and serves the site plus static assets over HTTP.

## Requirements

- Go 1.19 or newer
- A writable SQLite database path
- Media directories for the content you want to import
- Thumbnail directories if you want image URLs to resolve locally

## Configuration

The application loads environment variables from `env/.env` automatically when you start it from the `go/` directory.

At minimum, set:

- `MTVGO_DB_PATH` - path to the SQLite database file
- `MTVGO_RAW_ADDR` - host address for the HTTP server, such as `127.0.0.1`
- `MTVGO_SERVER_PORT` - port for the HTTP server, such as `8090`

Common optional variables:

- `MTVGO_MOVIES_PATH` - directory of movie files
- `MTVGO_TV_PATH` - directory of TV show files
- `MTVGO_VIDEOS_PATH` - directory of general video files
- `MTVGO_POSTER_PATH` - directory containing movie poster images
- `MTVGO_TV_POSTER_PATH` - directory containing TV poster images
- `MTVGO_THUMBNAIL_PATH` - directory served at `/thumbnails/`
- `MTVGO_TV_THUMBNAIL_PATH` - directory served at `/tvthumbnails/`
- `MTVGO_UPDATE_TOKEN` - token used by update handlers, if you use them

Example `env/.env`:

```env
MTVGO_DB_PATH=/absolute/path/to/mtvp.db
MTVGO_RAW_ADDR=127.0.0.1
MTVGO_SERVER_PORT=8090
MTVGO_MOVIES_PATH=/media/movies
MTVGO_TV_PATH=/media/tv
MTVGO_VIDEOS_PATH=/media/videos
MTVGO_POSTER_PATH=/media/posters
MTVGO_TV_POSTER_PATH=/media/tv-posters
MTVGO_THUMBNAIL_PATH=/media/thumbs
MTVGO_TV_THUMBNAIL_PATH=/media/tv-thumbs
```

## Build

From the `go/` directory:

```bash
go build
```

This produces a local `mtvp` binary in the same directory.

## Run

From the `go/` directory:

```bash
./mtvp
```

You can also run it without building first:

```bash
go run .
```

The program will:

1. Load `../env/.env`
2. Open the SQLite database from `MTVGO_DB_PATH`
3. Create and populate tables
4. Start the HTTP server on `MTVGO_RAW_ADDR:MTVGO_SERVER_PORT`

## Notes

- Run the application from `go/` so the relative `.env` path resolves correctly.
- If media paths are not set, the setup step skips those imports instead of failing.