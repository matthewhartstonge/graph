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
- start
- (Christchurch) --> (Wellington)
- (Wellington) --> (Auckland)
- (Auckland) --> (Gold Coast)

Solution 2:
- start
- (Christchurch) --> (Wellington)
- (Wellington) --> (Gold Coast)

Solution 3:
- start
- (Christchurch) --> (Auckland)
- (Auckland) --> (Gold Coast)

Solution 4:
- start
- (Christchurch) --> (Gold Coast)

No more solutions found!
```

```
Breadth-First Solution:
Solution 1:
- start
- (Christchurch) --> (Gold Coast)

Solution 2:
- start
- (Christchurch) --> (Auckland)
- (Auckland) --> (Gold Coast)

Solution 3:
- start
- (Christchurch) --> (Wellington)
- (Wellington) --> (Gold Coast)

Solution 4:
- start
- (Christchurch) --> (Wellington)
- (Wellington) --> (Auckland)
- (Auckland) --> (Gold Coast)

No more solutions found!
```
