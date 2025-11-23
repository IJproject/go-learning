# 図形の面積計算

## やること

図形ごとの面積を計算して一覧表示するプログラムを書け。

1. Shape インターフェースを定義
   1. Area() float64
   2. Name() string
2. 次の構造体を作り Shape を満たすようにメソッド実装
   1. Rectangle{Width, Height float64}
   2. Circle{Radius float64}
   3. Triangle{Base, Height float64}
3. 図形のスライスを作って全部の面積を出し、合計面積と最大面積の図形名も出力せよ。

```go:サンプルデータ
shapes := []Shape{
  Rectangle{3, 4},
  Circle{2},
  Triangle{5, 2},
  Rectangle{10, 2},
  Circle{1.5},
}
```

## 出力要件

- 各図形の名前と面積
- 合計面積
- 最大面積の図形名と面積

※ 円周率は math.Pi を使用せよ

## 目的

- interface の設計と実装
- 多態を使った共通処理
- []Shape の走査と集計
- 標準ライブラリ利用の最低限の慣れ
