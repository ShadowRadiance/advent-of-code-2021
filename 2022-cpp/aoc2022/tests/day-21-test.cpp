#include <gtest/gtest.h>

#include <days.h>

std::vector<std::string> const data{
    "root: pppw + sjmn",
    "dbpl: 5",
    "cczh: sllz + lgvd",
    "zczc: 2",
    "ptdq: humn - dvpt",
    "dvpt: 3",
    "lfqf: 4",
    "humn: 5",
    "ljgn: 2",
    "sjmn: drzm * dbpl",
    "sllz: 4",
    "pppw: cczh / lfqf",
    "lgvd: ljgn * ptdq",
    "drzm: hmdt - zczc",
    "hmdt: 32",
};

TEST(Day21, Part1Example)
{
    EXPECT_EQ("152", day_21::answer_a(data));
}

TEST(Day21, Part2Example)
{
    EXPECT_EQ("301", day_21::answer_b(data));
}