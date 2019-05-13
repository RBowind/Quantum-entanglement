# ARTS-Week5

## Algorithm
### problem:
**Generate Parentheses**

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
For example, given n = 3, a solution set is:
```
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```
To be honest, I really had no idea at first.

[solution](index.go)

## Review
[10 mistakes you should avoid](https://medium.com/@ajay.dutta94/10-mistakes-you-should-avoid-as-a-web-developer-c4ad4d2570f6)

## Tips
nohub 以及 &
- nohub

    不挂断地运行命令
- &
  
    在后台运行

一般两个一起使用: `nohup command &` 这样就能使命令永久在后台运行。但如果出现进程bug导致的崩溃，比如golang里面的panic，又该如何？见下便知。
## Share
[Supervisor 帮你解决进程崩溃烦恼](week-5-share.md)