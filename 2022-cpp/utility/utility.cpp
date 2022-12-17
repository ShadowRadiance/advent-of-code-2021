#include <iostream>
#include <vector>

void test_setup()
{
  for (auto v = std::vector{1, 2, 3}; auto &e : v)
  {
    std::cout << e;
  }

  std::cout << std::endl
            << "Hello, library!" << std::endl;
}