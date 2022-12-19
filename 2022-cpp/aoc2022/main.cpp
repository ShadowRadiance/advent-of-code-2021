#include <iostream>

#include <utility.h>
#include <day-01.h>
#include <day-02.h>
#include <day-03.h>
#include <day-04.h>
#include <day-05.h>
#include <day-06.h>
#include <day-07.h>
#include <day-08.h>
#include <day-09.h>
#include <day-10.h>
#include <day-11.h>
#include <day-12.h>
#include <day-13.h>
#include <day-14.h>
#include <day-15.h>
#include <day-16.h>
#include <day-17.h>
#include <day-18.h>
#include <day-19.h>
#include <day-20.h>
#include <day-21.h>
#include <day-22.h>
#include <day-23.h>
#include <day-24.h>
#include <day-25.h>

#include <filesystem>
namespace fs = std::filesystem;

int main(int argc, char** argv)
{
    std::cout << "Day 01 Answer A: " << day_01::answer_a(load_data("./data/day-01.txt")) << "\n";
    std::cout << "Day 01 Answer B: " << day_01::answer_b(load_data("./data/day-01.txt")) << "\n";
    std::cout << "Day 02 Answer A: " << day_02::answer_a(load_data("./data/day-02.txt")) << "\n";
    std::cout << "Day 02 Answer B: " << day_02::answer_b(load_data("./data/day-02.txt")) << "\n";
    std::cout << "Day 03 Answer A: " << day_03::answer_a(load_data("./data/day-03.txt")) << "\n";
    std::cout << "Day 03 Answer B: " << day_03::answer_b(load_data("./data/day-03.txt")) << "\n";
    std::cout << "Day 04 Answer A: " << day_04::answer_a(load_data("./data/day-04.txt")) << "\n";
    std::cout << "Day 04 Answer B: " << day_04::answer_b(load_data("./data/day-04.txt")) << "\n";
    std::cout << "Day 05 Answer A: " << day_05::answer_a(load_data("./data/day-05.txt")) << "\n";
    std::cout << "Day 05 Answer B: " << day_05::answer_b(load_data("./data/day-05.txt")) << "\n";
    std::cout << "Day 06 Answer A: " << day_06::answer_a(load_data("./data/day-06.txt")) << "\n";
    std::cout << "Day 06 Answer B: " << day_06::answer_b(load_data("./data/day-06.txt")) << "\n";
    std::cout << "Day 07 Answer A: " << day_07::answer_a(load_data("./data/day-07.txt")) << "\n";
    std::cout << "Day 07 Answer B: " << day_07::answer_b(load_data("./data/day-07.txt")) << "\n";
    std::cout << "Day 08 Answer A: " << day_08::answer_a(load_data("./data/day-08.txt")) << "\n";
    std::cout << "Day 08 Answer B: " << day_08::answer_b(load_data("./data/day-08.txt")) << "\n";
    std::cout << "Day 09 Answer A: " << day_09::answer_a(load_data("./data/day-09.txt")) << "\n";
    std::cout << "Day 09 Answer B: " << day_09::answer_b(load_data("./data/day-09.txt")) << "\n";
    std::cout << "Day 10 Answer A: " << day_10::answer_a(load_data("./data/day-10.txt")) << "\n";
    std::cout << "Day 10 Answer B: " << day_10::answer_b(load_data("./data/day-10.txt")) << "\n";
    std::cout << "Day 11 Answer A: " << day_11::answer_a(load_data("./data/day-11.txt")) << "\n";
    std::cout << "Day 11 Answer B: " << day_11::answer_b(load_data("./data/day-11.txt")) << "\n";
    std::cout << "Day 12 Answer A: " << day_12::answer_a(load_data("./data/day-12.txt")) << "\n";
    std::cout << "Day 12 Answer B: " << day_12::answer_b(load_data("./data/day-12.txt")) << "\n";
    std::cout << "Day 13 Answer A: " << day_13::answer_a(load_data("./data/day-13.txt")) << "\n";
    std::cout << "Day 13 Answer B: " << day_13::answer_b(load_data("./data/day-13.txt")) << "\n";
    std::cout << "Day 14 Answer A: " << day_14::answer_a(load_data("./data/day-14.txt")) << "\n";
    std::cout << "Day 14 Answer B: " << day_14::answer_b(load_data("./data/day-14.txt")) << "\n";
    std::cout << "Day 15 Answer A: " << day_15::answer_a(load_data("./data/day-15.txt")) << "\n";
    std::cout << "Day 15 Answer B: " << day_15::answer_b(load_data("./data/day-15.txt")) << "\n";
    std::cout << "Day 16 Answer A: " << day_16::answer_a(load_data("./data/day-16.txt")) << "\n";
    std::cout << "Day 16 Answer B: " << day_16::answer_b(load_data("./data/day-16.txt")) << "\n";
    std::cout << "Day 17 Answer A: " << day_17::answer_a(load_data("./data/day-17.txt")) << "\n";
    std::cout << "Day 17 Answer B: " << day_17::answer_b(load_data("./data/day-17.txt")) << "\n";
    std::cout << "Day 18 Answer A: " << day_18::answer_a(load_data("./data/day-18.txt")) << "\n";
    std::cout << "Day 18 Answer B: " << day_18::answer_b(load_data("./data/day-18.txt")) << "\n";
    std::cout << "Day 19 Answer A: " << day_19::answer_a(load_data("./data/day-19.txt")) << "\n";
    std::cout << "Day 19 Answer B: " << day_19::answer_b(load_data("./data/day-19.txt")) << "\n";
    std::cout << "Day 20 Answer A: " << day_20::answer_a(load_data("./data/day-20.txt")) << "\n";
    std::cout << "Day 20 Answer B: " << day_20::answer_b(load_data("./data/day-20.txt")) << "\n";
    std::cout << "Day 21 Answer A: " << day_21::answer_a(load_data("./data/day-21.txt")) << "\n";
    std::cout << "Day 21 Answer B: " << day_21::answer_b(load_data("./data/day-21.txt")) << "\n";
    std::cout << "Day 22 Answer A: " << day_22::answer_a(load_data("./data/day-22.txt")) << "\n";
    std::cout << "Day 22 Answer B: " << day_22::answer_b(load_data("./data/day-22.txt")) << "\n";
    std::cout << "Day 23 Answer A: " << day_23::answer_a(load_data("./data/day-23.txt")) << "\n";
    std::cout << "Day 23 Answer B: " << day_23::answer_b(load_data("./data/day-23.txt")) << "\n";
    std::cout << "Day 24 Answer A: " << day_24::answer_a(load_data("./data/day-24.txt")) << "\n";
    std::cout << "Day 24 Answer B: " << day_24::answer_b(load_data("./data/day-24.txt")) << "\n";
    std::cout << "Day 25 Answer A: " << day_25::answer_a(load_data("./data/day-25.txt")) << "\n";
    std::cout << "Day 25 Answer B: " << day_25::answer_b(load_data("./data/day-25.txt")) << "\n";
}
