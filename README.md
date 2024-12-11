名前候補探索器
==============

条件に当てはまる名前の候補を列挙します。五格分類法に対応しています。


使い方
-----

<details>
<summary>各コマンドのヘルプ</summary>

```console
$ name -h
Usage: name [subcommand] [options]

SUBCOMMANDS
  filter    name filter related commands
  search    search for given names
  info    show information about a given name

$ name search -h
Usage: name [options] <familyName>

OPTIONS
  -dir-dict string
        Directory of MeCab dictionary (full space only) (default "/opt/homebrew/opt/mecab-ipadic/lib/mecab/dic/ipadic")
  -max-length int
        Maximum length of a given name (default 3)
  -min-length int
        Minimum length of a given name (default 1)
  -space string
        Search spaces (available: full, common) (default "common")
  -yomi-count int
        Number of Yomi-Gana candidates (default 5)

STDIN
        See $ name filter try -h

EXAMPLES
        $ name search 山田 < ./filter.example.json
        評点    画数    名前    読み    天格    地格    人格    外格    総格
        15      13      一喜    イッキ  吉      大吉    大吉    大大吉  大吉
        15      13      一喜    イッキ  吉      大吉    大吉    大大吉  大吉
        ...

$ name info -h
Usage: name info [options] <familyName> <givenName> <yomi>

EXAMPLES
        $ name info 山田 太郎 タロウ
        評点    画数    名前    読み    天格    地格    人格    外格    総格
        8       13      太郎    タロウ  吉      大吉    大凶    大凶    大吉

$ name filter validate -h
Usage: name filter validate

EXAMPLES
        $ name filter validate < valid-filter.json
        $ echo $?
        0

        $ name filter validate < invalid-filter.json
        $ echo $?
        1

$ name filter try -h
Usage: name filter try <familyName> <givenName> <yomi>

STDIN
        JSON filter:

                filter: true or false or and or or or not or minRank or minTotalRank or mora or strokes or yomiCount or length
                true: {"true":{}}
                false: {"false":{}}
                and: {"and":[filter...]}
                or: {"or":[filter...]}
                not: {"not":filter}
                minRank: {"minRank":rank}
                rank: 0-4 (4=大大吉, 3=大吉, 2=吉, 1=凶, 0=大凶)
                minTotalRank: {"minTotalRank":byte}
                mora: {"maxMora":count}
                strokes: {"strokes":count}
                yomiCount: {"yomiCount":{"rune":string,"count":count}}
                length: {"length":count}
                count: {"equal":byte} or {"greaterThan":byte} or {"lessThan":byte}

EXAMPLES
        $ name filter try 田中 太郎 たなかたろう < filter.json
        $ echo $?
        0

        $ name filter try 田中 太郎 たなかたろう < filter.json
        $ echo $?
        1
```
</details>

まずフィルタを用意します。フィルタは名前の候補について真であれば候補を残し、偽であれば候補を除去します。

| 説明         | 構文                                                                   | 例                                                                                                                     |
|:-----------|:---------------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------|
| 真          | `{"true": {}}`                                                       | `{"true": {}}`                                                                                                        |
| 偽          | `{"false": {}}`                                                      | `{"false": {}}`                                                                                                       |
| 論理積        | `{"and": [filter...]}`                                               | `{"and": [{"yomiCount": {"rune": "ア", "count": {"equal": 1}}}, {"yomiCount": {"rune": "イ", "count": {"equal": 1}}}]}` |
| 論理和        | `{"or": [filter...]}`                                                | `{"or": [{"yomiCount": {"rune": "ア", "count": {"equal": 1}}}, {"yomiCount": {"rune": "イ", "count": {"equal": 1}}}]}`  |
| 否定論理       | `{"not": filter}`                                                    | `{"not": {"yomiCount": {"rune": "ア", "count": {"equal": 1}}}}`                                                        |
| 五格それぞれの最小値 | `{"minRank": count}`（`4`=大大吉, `3`=大吉, `2`=吉, `1`=凶, `0`=大凶）          | `{"minRank": 3}`                                                                                                      |
| 五格の合計値の最小値 | `{"minTotalRank": byte}`                                             | `{"minTotalRank": 11}`                                                                                                |
| 読み仮名の文字数   | `{"mora": count}`                                                    | `{"mora": {"equal": 3}}`                                                                                              |
| 画数         | `{"strokes": count}`                                                 | `{"strokes": {"lessThan": 25}}`                                                                                       |
| 指定した読み仮名の数 | `{"yomiCount": {"rune": "ア", "count": count}}`                       | `{"yomiCount": {"rune": "ア", "count": {"equal": 1}}}`                                                                 |
| よくある読み仮名   | `{"commonYomi": {}}`                                                 | `{"commonYomi": {}}`                                                                                                  |
| `count`    | `{"equal": byte}` or `{"lessThan": byte}` or `{"greaterThan": byte}` | `{"equal": 1}`                                                                                                        |

<details>
<summary>フィルタの例</summary>

```json
{
  "and": [
    {"mora": {"equal": 3}},
    {"minRank": 2},
    {"minTotalRank": 11},
    {"commonYomi": {}},
    {
      "or": [
        {
          "and": [
            {"yomiCount": {"rune": "ユ", "count": {"equal": 1}}},
            {"yomiCount": {"rune": "ウ", "count": {"equal": 0}}},
            {"yomiCount": {"rune": "サ", "count": {"lessThan": 2}}},
            {"yomiCount": {"rune": "キ", "count": {"equal": 0}}}
          ]
        },
        {
          "and": [
            {"yomiCount": {"rune": "ユ", "count": {"equal": 0}}},
            {"yomiCount": {"rune": "ウ", "count": {"equal": 1}}},
            {"yomiCount": {"rune": "サ", "count": {"lessThan": 2}}},
            {"yomiCount": {"rune": "キ", "count": {"equal": 0}}}
          ]
        },
        {
          "and": [
            {"yomiCount": {"rune": "ユ", "count": {"equal": 0}}},
            {"yomiCount": {"rune": "ウ", "count": {"equal": 0}}},
            {"yomiCount": {"rune": "サ", "count": {"equal": 0}}},
            {"yomiCount": {"rune": "キ", "count": {"equal": 1}}}
          ]
        }
      ]
    }
  ]
}
```
</details>

### 頻出空間探索

よくある人名の空間から名前候補を探索します。時間はほとんどかかりません。

```console
$ name search --space full 山田 < ./filter.json | tee result.tsv
評点    画数    名前    読み    天格    地格    人格    外格    総格
11      23      亜希保  アキホ  大凶    大吉    大吉    吉      大吉
11      26      啓穂    アキホ  大凶    吉      大大吉  大吉    吉
...
```

### 全空間探索

常用漢字+人名用漢字の空間から名前候補を探索します。かなり時間がかかります。現実的な時間で探索を終えるために `--max-length` を指定するなら `3` 以下を推奨します。

```console
$ name search --space common 山田 < ./filter.json | tee result.tsv
評点	画数	名前	読み	天格	地格	人格	外格	総格
14      15      隅己    スミキ  吉      大大吉  吉      大吉    大吉
15      24      隅期    スミキ  吉      大大吉  吉      大大吉  大吉
...
```

### 名前判定

```console
$ name info 山田 太郎 タロウ
評点    画数    名前    読み    天格    地格    人格    外格    総格
8       13      太郎    タロウ  吉      大吉    大凶    大凶    大吉
```

### フィルタ検査

```console
$ name filter validate < filter.json
$ echo $?
0
```

### フィルタ試験

```console
$ name filter try 山田 太郎 タロウ < filter.json
$ echo $?
1
```

インストール方法
----------------
### macOS

1. `brew install mecab mecab-ipadic` (+ 必要なら mecab-ipadic-neologd の辞書を用意)
2. 以下を実行：

    ```console
    $ export CGO_LDFLAGS="`mecab-config --libs`"
    $ export CGO_CFLAGS="-I`mecab-config --inc-dir`"
    $ go get github.com/Kuniwak/name
    ```


### Debian / Ubuntu

1. `sudo apt install mecab libmecab-dev mecab-ipadic-utf8` (+ 必要なら mecab-ipadic-neologd の辞書を用意)
2. 以下を実行：

    ```console
    $ export CGO_LDFLAGS="`mecab-config --libs`"
    $ export CGO_CFLAGS="-I`mecab-config --inc-dir`"
    $ go get github.com/Kuniwak/name
    ```

ライセンス
---------
MIT License
