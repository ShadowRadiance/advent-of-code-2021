#include <gtest/gtest.h>

#include <days.h>

TEST(Day05, Part1Example)
{
  std::vector<std::string> data{
    "    [D]    ",
    "[N] [C]    ",
    "[Z] [M] [P]",
    " 1   2   3 ",
    " ",
    "move 1 from 2 to 1",
    "move 3 from 1 to 3",
    "move 2 from 2 to 1",
    "move 1 from 1 to 2",
  };

  EXPECT_EQ("CMZ", day_05::answer_a(data));
}

TEST(Day05, Part2Example)
{
  std::vector<std::string> data{
    "    [D]    ",
    "[N] [C]    ",
    "[Z] [M] [P]",
    " 1   2   3 ",
    " ",
    "move 1 from 2 to 1",
    "move 3 from 1 to 3",
    "move 2 from 2 to 1",
    "move 1 from 1 to 2",
  };

  EXPECT_EQ("MCD", day_05::answer_b(data));
}