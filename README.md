名前候補探索器
==============

名前の候補を列挙します。五格分類法

使い方
-----

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
    {
      "mora": {
        "equal": 3
      }
    },
    {
      "minRank": 2
    },
    {
      "minTotalRank": 11
    },
    {
      "commonYomi": {}
    },
    {
      "or": [
        {
          "and": [
            {
              "yomiCount": {
                "rune": "ユ",
                "count": {
                  "equal": 1
                }
              }
            },
            {
              "yomiCount": {
                "rune": "ウ",
                "count": {
                  "equal": 0
                }
              }
            },
            {
              "yomiCount": {
                "rune": "サ",
                "count": {
                  "lessThan": 2
                }
              }
            },
            {
              "yomiCount": {
                "rune": "キ",
                "count": {
                  "equal": 0
                }
              }
            }
          ]
        },
        {
          "and": [
            {
              "yomiCount": {
                "rune": "ユ",
                "count": {
                  "equal": 0
                }
              }
            },
            {
              "yomiCount": {
                "rune": "ウ",
                "count": {
                  "equal": 1
                }
              }
            },
            {
              "yomiCount": {
                "rune": "サ",
                "count": {
                  "lessThan": 2
                }
              }
            },
            {
              "yomiCount": {
                "rune": "キ",
                "count": {
                  "equal": 0
                }
              }
            }
          ]
        },
        {
          "and": [
            {
              "yomiCount": {
                "rune": "ユ",
                "count": {
                  "equal": 0
                }
              }
            },
            {
              "yomiCount": {
                "rune": "ウ",
                "count": {
                  "equal": 0
                }
              }
            },
            {
              "yomiCount": {
                "rune": "サ",
                "count": {
                  "equal": 0
                }
              }
            },
            {
              "yomiCount": {
                "rune": "キ",
                "count": {
                  "equal": 1
                }
              }
            }
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
12	25	票羽河	ヒョウ	大凶	吉	大大吉	大吉	大吉
13	25	票応応	ヒョウ	大凶	吉	大大吉	大大吉	大吉
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

ライセンス
---------
MIT License
