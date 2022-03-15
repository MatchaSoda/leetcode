#include <stack>
#include <string>

using namespace std;

class Solution {
 public:
  bool isValid(string s) {
    stack<char> right_char_stack;
    for (char c : s) {
      if (c == '(') {
        right_char_stack.push(')');
      } else if (c == '{') {
        right_char_stack.push('}');
      } else if (c == '[') {
        right_char_stack.push(']');
      } else {
        if (right_char_stack.empty()) {
          return false;
        }
        if (c == right_char_stack.top()) {
          right_char_stack.pop();
        } else {
          return false;
        }
      }
    }
    if (!right_char_stack.empty()) {
      return false;
    }
    return true;
  }
};
