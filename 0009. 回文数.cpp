class Solution {
 public:
  bool isPalindrome(int input) {
    if (input < 0) {
      return false;
    }
    long long originalInput = input;
    long long reversedInput = 0;
    while (input != 0) {
      reversedInput = reversedInput * 10 + input % 10;
      input /= 10;
    }
    return originalInput == reversedInput;
  }
};
