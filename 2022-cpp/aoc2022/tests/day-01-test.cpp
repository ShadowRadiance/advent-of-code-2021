#include <gtest/gtest.h>

#include <days.h>

TEST(Day01, SingleElf)
{
  EXPECT_EQ("1000", day_01::answer_a({ "1000" }));
  EXPECT_EQ("3000", day_01::answer_a({ "1000", "2000" }));
  EXPECT_EQ("6000", day_01::answer_a({ "1000", "2000", "3000" }));
}

TEST(Day01, MultipleElves)
{
  EXPECT_EQ("6000", day_01::answer_a({ "1000", "2000", "3000", "", "5000" }));
  EXPECT_EQ("9000", day_01::answer_a({ "1000", "2000", "3000", "", "4000", "5000" }));
}


TEST(Day01, FullExample)
{
  std::vector<std::string> data{
      "1000",
      "2000",
      "3000",
      "",
      "4000",
      "",
      "5000",
      "6000",
      "",
      "7000",
      "8000",
      "9000",
      "",
      "10000",
  };


  // The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
  // The second Elf is carrying one food item with 4000 Calories.
  // The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
  // The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
  // The fifth Elf is carrying one food item with 10000 Calories.

  EXPECT_EQ("24000", day_01::answer_a(data));
}

TEST(Day01, Part2Example)
{
  std::vector<std::string> data{
      "1000",
      "2000",
      "3000",
      "",
      "4000",
      "",
      "5000",
      "6000",
      "",
      "7000",
      "8000",
      "9000",
      "",
      "10000",
  };

  EXPECT_EQ("45000", day_01::answer_b(data));
}