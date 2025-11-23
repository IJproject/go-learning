# 簡易在庫管理

## やること

商品在庫を管理するプログラムを書け。

1. Item 構造体
   1. ID string
   2. Name string
2. StockItem 構造体
   1. embedded: Item
   2. Quantity intID で
3. map[string]\*StockItem を作り、在庫を引けるようにする。
4. 次の操作を順に実行して、都度在庫を更新・表示せよ。
   1. ADD id qty：在庫を加算
   2. SELL id qty：在庫を減算（0 未満になるなら減算しないで警告）

```go:初期在庫
initial := []StockItem{
  {Item{"p1", "Pen"}, 10},
  {Item{"n1", "Notebook"}, 5},
  {Item{"e1", "Eraser"}, 20},
}
```

```go:操作例
ops := []string{
  "SELL p1 3",
  "ADD n1 10",
  "SELL e1 25",
  "SELL n1 4",
  "ADD p1 1",
}
```

## 出力要件

- 各操作後の在庫状態（全商品）
- 不正な SELL（在庫不足）のときは警告を出す

## 目的

- embedded struct の扱い
- map とポインタでのデータ更新
- 参照を介した値の変更の感覚を掴む
- 入力コマンドのパースと分岐
