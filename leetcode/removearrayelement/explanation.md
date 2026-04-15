### Explanation
Given int array `nums` and int `val`, remove all occurrences of `val` **in-place** (do not create a new array, modify `nums` itself), and return the amount of elements `k` thatare not equal to `val`. The remaining elements of `nums` are not important and *do not have to be deleted*.

This problem uses a *two pointer approach*. Essentially, have two variables for the current element of the array `k` and for the end of the array `n`.

`k := 0; n := len(nums)`

This gives us the first element in `nums` and the last element in `nums`. Next, we use a while loop, `for k < n`, and check if the current element of `nums` is equal to `val`. If it is, set the current element to the last element, and decrement `n`. This is because the last element is now equal to `val`, so we don't want to set the next instance of `val` to `val` again. If it it not equal to `val`, we increment `k` instead to go to the next element in the array.

```go
for k < n {
    // Check if current nums element = val
    if nums[k] == val {
        // Set current element to the last element
        // Remember n = length of the array, so we subtract 1
        // to get the actual last element.
        nums[k] = nums[n-1]
        
        // Decrement n to get the new last element that isn't val
        n--
        
        // This works because k is not incremented until it is not val,
        // so even if n == val, it will check the current element again
    } else {
        // Current element != val, increment k to go to the next element
        k++
    }

    return k
}
```
