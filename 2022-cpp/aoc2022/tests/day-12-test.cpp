#include <gtest/gtest.h>

#include <days.h>

const std::vector<std::string> data{
    "Sabqponm",
    "abcryxxl",
    "accszExk",
    "acctuvwj",
    "abdefghi",
};

TEST(Day12, Part1Example)
{
  EXPECT_EQ("31", day_12::answer_a(data));
}

TEST(Day12, Part2Example)
{
  EXPECT_EQ("29", day_12::answer_b(data));
}