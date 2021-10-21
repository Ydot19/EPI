# Asymptotic Bounding

The graph of relative time complexities:

![time complexity](./assets/time-complexity.png)
(credits: https://adrianmejia.com)

Terminology:
- Big-O notation: 
    - represents the upper bound of the running time of an algorithm
    - provides the worst case scenario
- Omega Notation:
    - represents the lower bound of the running time of an algorithm
- Theta Notatoin:
    - represents the fitness function that encloses the lower and upper bound

## Complexities

### Constant Time Complexity

**`O(1)`**

```python
# assume lst is a list
lst.pop()  # remove the element
lst.append(5)  # add an element, in the rare case of array resizing could be linear time
lst[0]  # retrieving an element at a position

# assume dct is a dictionary/hashmap

dct["key"] = 1  # retrieving or assigning a key-value

# assum st is a set

print(5 in st)  # hash sets are constant time
```

### Logarithmic Time
**`O(log(n))`**
- Popularly associated with divide and conquer problems
- Binary Search
- Traversing a Heap

### Linear Time
**`O(n)`**
- Associated popularly with Linked Lists and Arrays

- O(n)
- O(max(m,n))
- O(min(m,n))
- O(m + n)