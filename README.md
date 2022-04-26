# easyindex-cli

## preinstall

- create Google service account
- create credential json file for Google Indexing API
- add your service account as a site owner on search console

ref. [Google Search Central](https://developers.google.com/search/apis/indexing-api/v3/prereqs)

## required

- Google service account
- credential json file for Google Indexing API

### install

Mac:

```sh
# example

```

Linux:

```sh

```

Windows:

```sh

```

### Usage

### quickstart

Mac / Linux:

```sh
## updated
indexing publish updated -c (your credential json file path) https://example.com/foobar https://example.com/fizzbizz
 
## deleted
indexing.exe publish deleted -c (your credential json file path) https://example.com/foobar https://example.com/fizzbizz
```

Windows:

```sh
## updated
indexing publish updated -c (your credential json file path) https://example.com/foobar https://example.com/fizzbizz
 
## deleted
indexing.exe publish deleted -c (your credential json file path) https://example.com/foobar https://example.com/fizzbizz
```

### basic usage

> :smile:NOTE:  
>   You can omit specifying the credential file path by setting it in an environment variable (`EASYINDEX_CREDENTIAL_PATH`).

```sh
## Mac /Linux:
indexing publish updated https://example.com/foobar https://example.com/fizzbizz

## Windows:
indexing.exe publish updated https://example.com/foobar https://example.com/fizzbizz
```

### skip error pages

> By default, if a problem is found on specified pages before calling the indexing API, it will exit with an error.
>
> By enabling the skip flag, you can exclude problematic pages and send requests to the API.

```sh
## Mac /Linux:
indexing publish updated --skip https://example.com/foobar https://example.com/fizzbizz

## Windows:
indexing.exe publish updated --skip https://example.com/foobar https://example.com/fizzbizz
```

### set quota

> You can limit the number of requests.
>
> Please set based on the quota of Google Indexing API.

```sh
## Mac /Linux:
indexing publish update --limit 200 https://example.com/foobar https://example.com/fizzbizz

## Windows:
indexing.exe publish update --limit 200 https://example.com/foobar https://example.com/fizzbizz
```

### flags

| flag | shorthand | description |
|---|---|---|
| `credentials` | c | credential file path |
| `ignore` | i | ignore pre-check |
| `skip` | s | skip error pages |
| `limit` | l | limit the number of API request |

### environment variables

| key | description |
|---|---|
| `EASYINDEX_CREDENTIALS_PATH` | credential file path |
| `EASYINDEX_IGNORE_PRECHECK` | ignore pre-check |
| `EASYINDEX_SKIP` | skip error pages |
| `EASYINDEX_REQUEST_LIMIT` | limit the number of API request |
