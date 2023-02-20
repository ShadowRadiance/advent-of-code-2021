#include <gtest/gtest.h>

#include <days.h>

TEST(Day09, Part1Example)
{
    std::vector<std::string> data{
      "R 4",
      "U 4",
      "L 3",
      "D 1",
      "R 4",
      "D 1",
      "L 5",
      "R 2",
    };

    EXPECT_EQ("13", day_09::answer_a(data));
}

TEST(Day09, Part2Example)
{
    std::vector<std::string> data{
      "R 5",
      "U 8",
      "L 8",
      "D 3",
      "R 17",
      "D 10",
      "L 25",
      "U 20",
    };

    EXPECT_EQ("36", day_09::answer_b(data));
}