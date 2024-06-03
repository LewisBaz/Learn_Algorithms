package floydstortoisehare

// https://leetcode.com/problems/find-the-duplicate-number/description/

func FindDuplicate(nums []int) int {
    tortoise, hare := 0, 0

    // Phase 1: Cycle Detection
    for {
        tortoise = nums[tortoise]
        hare = nums[nums[hare]]
        if tortoise == hare {
            break
        }
    }

    // Phase 2: Cycle Start Detection
    tortoise = 0
    for tortoise != hare {
        tortoise = nums[tortoise]
        hare = nums[hare]
    }

    return tortoise
}

// Функция работает корректно, но слишком долго
func FindDuplicate2(nums []int) int {
	hareStep := 2
	tortoise := 0
	hare := len(nums)/2
	var tortoise_value int = nums[tortoise]
	var hare_value int = nums[hare]

	met := false
    for tortoise_value != hare_value {
		tortoise += 1
		hare += hareStep
		if tortoise >= len(nums) {
			next := tortoise - len(nums) 
			tortoise = next
		}
		if hare >= len(nums) {
			next := hare - len(nums) + 1
			hare = next
		}
		if tortoise == hare {
			if met {
				return nums[tortoise]
			}
			met = true
			tortoise = 0
			hareStep = 1
		}
		tortoise_value = nums[tortoise]
		hare_value = nums[hare]
	}
	return tortoise_value
}