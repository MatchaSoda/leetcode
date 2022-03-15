#include <unordered_map>

using namespace std;

class Solution {
 public:
  unordered_map<char, int> map = {
      // roman value
      {'I', 1},   {'V', 5},   {'X', 10},   {'L', 50},
      {'C', 100}, {'D', 500}, {'M', 1000},
  };
  int romanToInt(string s) {
    int result = 0;
    for (int i = 0; i < s.size(); i++) {
      int current_value = map[s[i]];
      if (map[s[i + 1]] > current_value) {
        current_value *= -1;
      }
      result += current_value;
    }
    return result;
  }
};
