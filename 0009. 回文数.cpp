class Solution {
 public:
  bool isPalindrome(int input) {
    if (input < 0) {
      return false;
    }
    long long original_input = input;
    long long reversed_input = 0;
    while (input) {
      reversed_input = reversed_input * 10 + input % 10;
      input /= 10;
    }
    return original_input == reversed_input;
  }
};
