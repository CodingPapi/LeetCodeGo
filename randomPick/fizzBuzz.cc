#include <iostream>                // std::cout
#include <thread>                // std::thread
#include <mutex>                // std::mutex, std::unique_lock
#include <condition_variable>    // std::condition_variable
using namespace std;

class FizzBuz {
private:
  int n;
public:
  FizzBuz(int n) {
      this->n = n;
  }

  void fizz(function<void()> printFizz) {

  }
  void buzz(function<void()> printBuzz) {

  }
  void fizzBuzz(function<void()> printFizzBuzz) {

  }
  void number(function<void(int)> printNumber) {

  }

}
