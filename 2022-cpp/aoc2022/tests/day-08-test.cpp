#include <gtest/gtest.h>

#include <days.h>

TEST(Day08, Part1Example)
{
    std::vector<std::string> data{
        "30373",
        "25512",
        "65332",
        "33549",
        "35390",
    };

    EXPECT_EQ("21", day_08::answer_a(data));
}

TEST(Day08, Part2Example)
{
    std::vector<std::string> data{};

    EXPECT_EQ("PENDING", day_08::answer_b(data));
}
