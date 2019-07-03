# Why this way
The site provides content as HTML, rather than JSON (which is much more parse-able).
Hence first comes transforming byte into HTML into String.
Then slicing, iterating, matching, printing.

# How To:

```
go get ./...
go build -o tester main.go
./tester http://hiring.test.everon.io/
```


# Potential improvements:

1. Using proper CLI package like spf13/cobra
2. Better HTML-to-StringSlice conversion (if exists)
3. More helper methods (to avoid some code repetition)