#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
 public:
  vector<int> twoSum(vector<int>& nums, int target) {
    unordered_map<int, int> map;  // value index
    for (int i = 0; i < nums.size(); i++) {
      if (!map.count(nums[i])) {
        map.insert({target - nums[i], i});
      } else {
        return {i, map[nums[i]]};
      }
    }
    return {};
  }
};