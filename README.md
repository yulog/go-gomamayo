# go-gomamayo

[![Go Reference](https://pkg.go.dev/badge/github.com/yulog/go-gomamayo.svg)](https://pkg.go.dev/github.com/yulog/go-gomamayo)

[ゴママヨ](https://thinaticsystem.com/glossary/gomamayo)検出器のGo言語版です。[na2na-p/gomamayo-deno](https://github.com/na2na-p/gomamayo-deno)を参考にした再実装です。  
形態素解析には[ikawaha/kagome](https://github.com/ikawaha/kagome)を使っています。  
辞書にはIPADIC、NEologd、unidic、unidic3を使えます。

> [!WARNING]
> 辞書を含むためファイルサイズが大きいです。(700MB+)

## Library

```
go get -u github.com/yulog/go-gomamayo@latest
```

## CLI
```
go install github.com/yulog/go-gomamayo/cmd/gomamayo@latest
```

### Example

```
gomamayo analyze ごまマヨネーズ
```

```json
{
  "isGomamayo": true,
  "combo": 1,
  "detail": [
    {
      "surface": "ごま|マヨネーズ",
      "dimension": 1,
      "rawResult1": {
        "id": 28368,
        "start": 0,
        "end": 2,
        "surface": "ごま",
        "class": "KNOWN",
        "pos": [
          "名詞",
          "一般",
          "*",
          "*"
        ],
        "base_form": "ごま",
        "reading": "ゴマ",
        "pronunciation": "ゴマ",
        "features": [
          "名詞",
          "一般",
          "*",
          "*",
          "*",
          "*",
          "ごま",
          "ゴマ",
          "ゴマ"
        ]
      },
      "rawResult2": {
        "id": 99158,
        "start": 2,
        "end": 7,
        "surface": "マヨネーズ",
        "class": "KNOWN",
        "pos": [
          "名詞",
          "一般",
          "*",
          "*"
        ],
        "base_form": "マヨネーズ",
        "reading": "マヨネーズ",
        "pronunciation": "マヨネーズ",
        "features": [
          "名詞",
          "一般",
          "*",
          "*",
          "*",
          "*",
          "マヨネーズ",
          "マ ヨネーズ",
          "マヨネーズ"
        ]
      }
    }
  ]
}
```

### Usage

```
gomamayo analyze --disable-ignore ごまマヨネーズ
```

```
gomamayo ignore add サラダ
```

```
gomamayo ignore remove サラダ
```

## License

MIT

## Author

yulog