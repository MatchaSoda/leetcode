#include <unordered_map>

class Solution {
 public:
  std::unordered_map<char, int> romanCharValues = {
      // roman value
      {'I', 1},   {'V', 5},   {'X', 10},   {'L', 50},
      {'C', 100}, {'D', 500}, {'M', 1000},
  };
  int romanToInt(std::string s) {
    int result = 0;
    for (int charIndex = 0; charIndex < s.size(); charIndex++) {
      int currentValue = romanCharValues[s[charIndex]];
      if (charIndex + 1 < s.size() &&
          romanCharValues[s[charIndex + 1]] > currentValue) {
        result -= currentValue;
      } else {
        result += currentValue;
      }
    }
    return result;
  }
};
