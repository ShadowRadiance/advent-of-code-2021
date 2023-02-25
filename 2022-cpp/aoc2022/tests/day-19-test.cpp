#include <gtest/gtest.h>

#include <days.h>

std::vector<std::string> const data{
    "Blueprint 1: "
        "Each ore robot costs 4 ore. "
        "Each clay robot costs 2 ore. "
        "Each obsidian robot costs 3 ore and 14 clay. "
        "Each geode robot costs 2 ore and 7 obsidian.",
    "Blueprint 2: "
        "Each ore robot costs 2 ore. "
        "Each clay robot costs 3 ore. "
        "Each obsidian robot costs 3 ore and 8 clay. "
        "Each geode robot costs 3 ore and 12 obsidian.",
};

TEST(Day19, Part1Example)
{
    EXPECT_EQ("33", day_19::answer_a(data));
}

TEST(Day19, Part2Example)
{
    EXPECT_EQ("PENDING", day_19::answer_b(data));
}
