#include <unordered_map>
#include <vector>

class Solution {
 public:
  std::vector<int> twoSum(std::vector<int>& nums, int target) {
    std::unordered_map<int, int> existedNumbers;  // value index
    for (int i = 0; i < nums.size(); i++) {
      if (existedNumbers.count(target - nums[i]) == 0) {
        existedNumbers.insert({nums[i], i});
      } else {
        return {i, existedNumbers[target - nums[i]]};
      }
    }
    return {};
  }
};
