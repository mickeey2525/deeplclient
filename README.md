# DeepL Cli Client
This is Unofficial Cli Client for the API of [deepl.com](https://www.deepl.com/)

## Usage

```
deepl -text これはテストです -srclang JA -targetlang EN
```

The response would be like below
```
{"translations":[{"detected_source_language":"JA","text":"This is a test."}]}%
```

## Installation

```
go get -u github.com/mickeey2525/deeplclient
```