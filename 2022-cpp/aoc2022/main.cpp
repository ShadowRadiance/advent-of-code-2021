#include <iostream>
#include <vector>

#include "utility.h"

int main(int argc, char** argv)
{
    for (auto v = std::vector{ 1, 2, 3 }; auto & e : v)
    {
        std::cout << e;
    }

    std::cout << std::endl
        << "Hello, world!" << std::endl;

    test_setup();
}
