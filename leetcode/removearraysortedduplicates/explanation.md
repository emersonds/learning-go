### Remove Duplicates from Sorted Array
*Time complexity: O(n)*
*Space complexity: O(1)*

```
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
```

This problem uses a two pointer approach. In this case, `k` is the number of unique elements, which also keeps track of our current unique element in the `nums` array. `i` is declared as it usually is, an iterator through the `for` loop. But, `i` is also our second pointer in this case, where each element is checked against the element at `nums[k]`.

The way this two pointer problem works is simple. We use `k` to keep track of the current unique element in the array, then check each other value with `i`, setting `nums[k] = nums[i]` if `nums[i] != nums[k-1]` and incrementing `k`. We check `nums[k-1]` so `i` is always one element *ahead* of `k`. Indexes 0 and 1 will always be unique to each other, so it makes sense to check `i` against `k-1`.

If `nums[i] == nums[k]`, this if-block will be ignored. So, `i` will keep incrementing, while `k` stays the same until a unique element is found.

For example:
```
nums = [0,1,1,1,1,2,2,3,3,4]
            k   i 

nums[i] = 1
nums[k - 1] = 1

Since `nums[i] == nums[k - 1]`, we do nothing
and move `i` to the next.

`k` remains 2.

nums = [0,1,1,1,1,2,2,3,3,4]
            k     i
```
