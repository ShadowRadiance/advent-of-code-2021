#include <gtest/gtest.h>

#include <days.h>

const std::vector<std::string> data{
    "        ...#",
    "        .#..",
    "        #...",
    "        ....",
    "...#.......#",
    "........#...",
    "..#....#....",
    "..........#.",
    "        ...#....",
    "        .....#..",
    "        .#......",
    "        ......#.",
    "",
    "10R5L5R10L4R5L5",
};

TEST(Day22, Part1Example)
{
    EXPECT_EQ("6032", day_22::answer_a(data));
}

TEST(Day22, Part2Example)
{
    day_22::set_test_mode(true);
    EXPECT_EQ("5031", day_22::answer_b(data));
}