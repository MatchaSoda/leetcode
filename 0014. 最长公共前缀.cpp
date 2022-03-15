#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  string longestCommonPrefix(vector<string>& strs) {
    string result = "";
    for (int i = 0; i < strs[0].size(); i++) {
      char first_line_char = strs[0][i];
      for (int j = 1; j < strs.size(); j++) {
        if (strs[j][i] != first_line_char) {
          return result;
        }
      }
      result += first_line_char;
    }
    return result;
  }
};
