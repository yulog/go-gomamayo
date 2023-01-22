# go-gomamayo

[ゴママヨ](https://thinaticsystem.com/glossary/gomamayo)検出器のGo言語版です。[na2na-p/gomamayo-deno](https://github.com/na2na-p/gomamayo-deno)を参考にした部分的な再実装です。
形態素解析の辞書にはIPADICを使っています。

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


## License

MIT

## Author

yulog