cmake_minimum_required(VERSION 3.23)
cmake_policy(SET CMP0135 NEW)

project(AdventOfCode2022)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

add_subdirectory(utility)
add_subdirectory(aoc2022)
