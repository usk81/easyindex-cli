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

Mac/Linux:

```sh
wget -O- https://github.com/usk81/easyindex-cli/releases/download/{version}/easyindex-cli_{version}_{os}.tar.gz | tar xz
# e.g.
# mac M1 / version v0.0.1
# wget -O- https://github.com/usk81/easyindex-cli/releases/download/v0.0.1/easyindex-cli_v0.0.1_darwin_ard64.tar.gz | tar xz
```

Windows (PowerShell):

```sh
iwr -outf easyindex-cli.tar.gz https://github.com/usk81/easyindex-cli/releases/download/{version}/easyindex-cli_{version}_{os}.tar.gz
# e.g.
# 64bit / version v0.0.1
# iwr -outf easyindex-cli.tar.gz https://github.com/usk81/easyindex-cli/releases/download/{version}/easyindex-cli_v0.0.1_windows_amd64.tar.gz
```

### Usage

### quickstart

Mac / Linux:

```sh
## updated
easyindex publish updated -c (your credential json file path) https://example.com/foobar https://example.com/fizzbizz
 
## deleted
easyindex.exe publish deleted -c (your credential json file path) https://example.com/foobar https://example.com/fizzbizz
```

Windows:

```sh
## updated
easyindex publish updated -c (your credential json file path) https://example.com/foobar https://example.com/fizzbizz
 
## deleted
easyindex.exe publish deleted -c (your credential json file path) https://example.com/foobar https://example.com/fizzbizz
```

### basic usage

> :smile:NOTE:  
>   You can omit specifying the credential file path by setting it in an environment variable (`EASYINDEX_CREDENTIAL_PATH`).

```sh
## Mac /Linux:
easyindex publish updated https://example.com/foobar https://example.com/fizzbizz

## Windows:
easyindex.exe publish updated https://example.com/foobar https://example.com/fizzbizz
```

### skip error pages

> By default, if a problem is found on specified pages before calling the indexing API, it will exit with an error.
>
> By enabling the skip flag, you can exclude problematic pages and send requests to the API.

```sh
## Mac /Linux:
easyindex publish updated --skip https://example.com/foobar https://example.com/fizzbizz

## Windows:
easyindex.exe publish updated --skip https://example.com/foobar https://example.com/fizzbizz
```

### set quota

> You can limit the number of requests.
>
> Please set based on the quota of Google Indexing API.

```sh
## Mac /Linux:
easyindex publish updated --limit 200 https://example.com/foobar https://example.com/fizzbizz

## Windows:
easyindex.exe publish updated --limit 200 https://example.com/foobar https://example.com/fizzbizz
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
