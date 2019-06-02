# ARTS-Week7

## Algorithm

##### Wildcard Matching

Given an input string (`s`) and a pattern (`p`), implement wildcard pattern matching with support for `'?'` and `'*'`.

'?' Matches any single character.
'\*' Matches any sequence of characters (including the empty sequence).

The matching should cover the **entire** input string (not partial).

```go
func isMatch(s, p string) bool {
    prev, now := make([]bool, len(p)+1), make([]bool, len(p)+1)
    for i := 0; i <= len(s); i++ {
        now, prev = prev, now
        now[0] = i == 0
        for j := 1; j <= len(p); j++ {
            if p[j-1] == '*' {
                now[j] = prev[j] || prev[j-1] || now[j-1]
            } else {
                now[j] = prev[j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
            }
        }
    }
    return now[len(p)]
}
```

## Review

[How To Get 10X The Value From Every Book You Read](https://medium.com/@Barry.Davret/how-to-get-10x-the-value-from-every-book-you-read-90166d672988)

This article is talking about the efficiency of reading. In summary,

- Read and Mark

- Reread it and get the points

- Edit your mark , make it thinner

- Review your mark two weeks a day.

## Tips

###### In MySQL, never use “utf8”. Use “utf8mb4”.

Beacause:

- MySQL’s “utf8mb4” means “UTF-8”.
- MySQL’s “utf8” means “a proprietary character encoding”. This encoding can’t encode many Unicode characters.

## Share

[数据结构概述与应用（一）](week-7-share.md)
