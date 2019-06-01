# ARTS-Week7

## Algorithm

##### Wildcard Matching

Given an input string (`s`) and a pattern (`p`), implement wildcard pattern matching with support for  `'?'`  and  `'*'`.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).

The matching should cover the  **entire**  input string (not partial).

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

## Tips

## Share
