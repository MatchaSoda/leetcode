#include <stack>
#include <string>

class Solution {
 public:
  bool isValid(std::string s) {
    std::stack<char> leftParenthesis;
    for (char currentChar : s) {
      if (currentChar == '(' || currentChar == '[' || currentChar == '{') {
        leftParenthesis.push(currentChar);
      } else {
        if (leftParenthesis.empty()) {
          return false;
        }
        if (currentChar == ')' && leftParenthesis.top() == '(' ||
            currentChar == ']' && leftParenthesis.top() == '[' ||
            currentChar == '}' && leftParenthesis.top() == '{') {
          leftParenthesis.pop();
        } else {
          return false;
        }
      }
    }
    return leftParenthesis.empty();
  }
};
