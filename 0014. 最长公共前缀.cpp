#include <string>
#include <vector>

class Solution {
 public:
  std::string longestCommonPrefix(std::vector<std::string>& strs) {
    std::string result = "";
    for (int charIndex = 0; charIndex < strs[0].size(); charIndex++) {
      char firstLineChar = strs[0][charIndex];
      for (int stringIndex = 1; stringIndex < strs.size(); stringIndex++) {
        if (strs[stringIndex][charIndex] != firstLineChar) {
          return result;
        }
      }
      result += firstLineChar;
    }
    return result;
  }
};
