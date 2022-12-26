#include <gtest/gtest.h>

#include <cpu_10.h>

TEST(CPU_10, SignalStrength)
{
    std::vector<std::string> instructions{
        "noop",
        "addx 3",
        "addx -5"
    };

    CPU_10 cpu{ instructions };
    EXPECT_EQ(1, cpu.x());
    EXPECT_EQ(0, cpu.current_cycle());
    EXPECT_EQ(0, cpu.signal_strength());
    EXPECT_EQ(5, cpu.cycles_required());

    cpu.begin_cycle();              // noop begins
    EXPECT_EQ(1, cpu.x());
    EXPECT_EQ(1, cpu.current_cycle());
    EXPECT_EQ(1, cpu.signal_strength());
    cpu.end_cycle();                // noop completes
    EXPECT_EQ(1, cpu.x());
    EXPECT_EQ(1, cpu.current_cycle());
    EXPECT_EQ(1, cpu.signal_strength());

    cpu.begin_cycle();              // addx 3 begins
    EXPECT_EQ(1, cpu.x());
    EXPECT_EQ(2, cpu.current_cycle());
    EXPECT_EQ(2, cpu.signal_strength());
    cpu.end_cycle();                // addx 3 executing
    EXPECT_EQ(1, cpu.x());
    EXPECT_EQ(2, cpu.current_cycle());
    EXPECT_EQ(2, cpu.signal_strength());

    cpu.begin_cycle();              // addx 3 executing
    EXPECT_EQ(1, cpu.x());
    EXPECT_EQ(3, cpu.current_cycle());
    EXPECT_EQ(3, cpu.signal_strength());
    cpu.end_cycle();                // addx 3 completes
    EXPECT_EQ(4, cpu.x());
    EXPECT_EQ(3, cpu.current_cycle());
    EXPECT_EQ(12, cpu.signal_strength());

    cpu.begin_cycle();              // addx -5 begins
    EXPECT_EQ(4, cpu.x());
    EXPECT_EQ(4, cpu.current_cycle());
    EXPECT_EQ(16, cpu.signal_strength());
    cpu.end_cycle();                // addx -5 executing
    EXPECT_EQ(4, cpu.x());
    EXPECT_EQ(4, cpu.current_cycle());
    EXPECT_EQ(16, cpu.signal_strength());

    cpu.begin_cycle();              // addx -5 executing
    EXPECT_EQ(4, cpu.x());
    EXPECT_EQ(5, cpu.current_cycle());
    EXPECT_EQ(20, cpu.signal_strength());
    cpu.end_cycle();                // addx -5 completes
    EXPECT_EQ(-1, cpu.x());
    EXPECT_EQ(5, cpu.current_cycle());
    EXPECT_EQ(-5, cpu.signal_strength());
    EXPECT_EQ(-5, cpu.signal_strength());
}