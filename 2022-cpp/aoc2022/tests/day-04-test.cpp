#include <gtest/gtest.h>

#include <days.h>

TEST(Day04, Part1Example)
{
    std::vector<std::string> data{
      "2 - 4,6 - 8",
      "2 - 3,4 - 5",
      "5 - 7,7 - 9",
      "2 - 8,3 - 7",
      "6 - 6,4 - 6",
      "2 - 6,4 - 8",
    };

    EXPECT_EQ("2", day_04::answer_a(data));
}

TEST(Day04, Part2Example)
{
    std::vector<std::string> data{
      "2 - 4,6 - 8",
      "2 - 3,4 - 5",
      "5 - 7,7 - 9",
      "2 - 8,3 - 7",
      "6 - 6,4 - 6",
      "2 - 6,4 - 8",
    };

    EXPECT_EQ("4", day_04::answer_b(data));
}