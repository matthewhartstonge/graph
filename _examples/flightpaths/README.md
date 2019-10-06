## Flightpaths

flightpaths implements a trivial graph solving problem to showcase the 
differences in the order solution outcomes are retrieved due to the search 
strategy used.

### Problem
Given a flight schedule, we want to be able to generate a flight path that we 
can take that departs from:

* Christchurch, New Zealand

And gets us to:

* The Gold Coast, Australia


### Outcomes

```
Depth-First Solution:
Solution 1:
Total Cost: 3020.00
- start
- (Christchurch) -(305.0)-> (Wellington)
- (Wellington) -(491.0)-> (Auckland)
- (Auckland) -(2224.0)-> (Gold Coast)

Solution 2:
Total Cost: 2742.00
- start
- (Christchurch) -(305.0)-> (Wellington)
- (Wellington) -(2437.0)-> (Gold Coast)

Solution 3:
Total Cost: 2987.00
- start
- (Christchurch) -(763.0)-> (Auckland)
- (Auckland) -(2224.0)-> (Gold Coast)

Solution 4:
Total Cost: 2434.00
- start
- (Christchurch) -(2434.0)-> (Gold Coast)

No more solutions found!

Breadth-First Solution:
Solution 1:
Total Cost: 2434.00
- start
- (Christchurch) -(2434.0)-> (Gold Coast)

Solution 2:
Total Cost: 2987.00
- start
- (Christchurch) -(763.0)-> (Auckland)
- (Auckland) -(2224.0)-> (Gold Coast)

Solution 3:
Total Cost: 2742.00
- start
- (Christchurch) -(305.0)-> (Wellington)
- (Wellington) -(2437.0)-> (Gold Coast)

Solution 4:
Total Cost: 3020.00
- start
- (Christchurch) -(305.0)-> (Wellington)
- (Wellington) -(491.0)-> (Auckland)
- (Auckland) -(2224.0)-> (Gold Coast)

No more solutions found!

Lowest-Cost First Solution:
Solution 1:
Total Cost: 2434.00
- start
- (Christchurch) -(2434.0)-> (Gold Coast)

Solution 2:
Total Cost: 2742.00
- start
- (Christchurch) -(305.0)-> (Wellington)
- (Wellington) -(2437.0)-> (Gold Coast)

Solution 3:
Total Cost: 2987.00
- start
- (Christchurch) -(763.0)-> (Auckland)
- (Auckland) -(2224.0)-> (Gold Coast)

Solution 4:
Total Cost: 3020.00
- start
- (Christchurch) -(305.0)-> (Wellington)
- (Wellington) -(491.0)-> (Auckland)
- (Auckland) -(2224.0)-> (Gold Coast)

No more solutions found!
```
