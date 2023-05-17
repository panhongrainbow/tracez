# package compress

> 1. This package mainly compresses `Json data`.
>    It makes it convenient for the `tracez client` to process the data later.
> 2. The whole program is divided into `two parts`.
>    One part `uses unsafe pointer`, the other does `not use unsafe pointer`.

## Introduction

When I used unsafe pointers to rewrite the string processing program, this action is called `a hack`.

At this time, it will `bypass the compiler's checks and directly operate the memory` to obtain the highest performance. It is an act that many people will do.

I found that the performance increased a lot, `at least 3 to 10 times faster`.
(直接绕过编译器检查机制去操作内存，很快，但会怕 race，直接处理不用复制，当然快，转成 string ，复制来复制去，要死啦，但是这样做，就算上锁，也是很快)

There are two main reasons:

1. When reading the file, the whole data is a byte slice, which can be processed directly `without converting it into a string`.
2. `Avoid data copying behavior` as much as possible. Reducing one data copy can maximize performance.

However, I felt it was unfair later.

It is because `slices and strings` are `not` as afraid of `data race` as `unsafe pointers`.

To compare performance, the programs written with `unsafe pointers must all be locked` before they can be compared.

Later I found that even if the programs written with `unsafe pointers were locked`, they were still `much faster than string functions`.

However, although the programs written with unsafe pointers are very fast, they have `not` gone through `the compiler's checking mechanism`. So, start using them from `unit testing`.

## Comparison

make a comparison through benchmark testing

| function | Created by unsafe pointers | Operated by string |
| -------- | -------------------------- | ------------------ |
| OneLine  | 33.68 ns/op ~ 36.91 ns/op  | 1029 ns/op         |
| Separate | 221.4 ns/op ~ 239.0 ns/op  | 671.8 ns/op        |