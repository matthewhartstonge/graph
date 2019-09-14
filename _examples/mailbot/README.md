### Mailbot

This solving example is derived from Chapter 3.3.1, Example 3.5 in "[Artificial 
Intelligence: Foundations of Computational Agents, Poole & Mackworth.](https://artint.info/2e/html/ArtInt2e.Ch3.S3.SS1.html)"

```
N = {mail, ts, o103, ... }
A = {⟨ts,mail⟩, ⟨o103,ts⟩, ⟨o103,b3⟩, ⟨o103,o109⟩, ...}
```

Consider the problem of the delivery robot finding a path from location `o103`
to location `r123`. 
 
There are three paths from `o103` to `r123`:

``` 
 ⟨o103,o109,o119,o123,r123⟩
 ⟨o103,b3,b4,o109,o119,o123,r123⟩
 ⟨o103,b3,b1,b2,b4,o109,o119,o123,r123⟩
```
 
If `o103` were the start node and `r123` were the unique goal node, each of 
these three paths would be a solution to the graph-searching problem.
The first of these is an optimal solution, with a solution cost of
`12 + 16 + 9 + 4 = 41`.