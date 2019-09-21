## Mailbot

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

### Output

Using a depth-first search strategy:

```
Solution 1:
- start
- (o103) -(12.0)-> (o109)
- (o109) -(16.0)-> (o119)
- (o119) -(9.0)-> (o123)
- (o123) -(4.0)-> (r123)

Solution 2:
- start
- (o103) -(4.0)-> (b3)
- (b3) -(7.0)-> (b4)
- (b4) -(7.0)-> (o109)
- (o109) -(16.0)-> (o119)
- (o119) -(9.0)-> (o123)
- (o123) -(4.0)-> (r123)

Solution 3:
- start
- (o103) -(4.0)-> (b3)
- (b3) -(4.0)-> (b1)
- (b1) -(6.0)-> (b2)
- (b2) -(3.0)-> (b4)
- (b4) -(7.0)-> (o109)
- (o109) -(16.0)-> (o119)
- (o119) -(9.0)-> (o123)
- (o123) -(4.0)-> (r123)

No more solutions found!

time="2019-09-22T02:16:54+12:00" level=info took="849.117µs"
```