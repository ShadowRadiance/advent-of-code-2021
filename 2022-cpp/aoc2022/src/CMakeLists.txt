add_library(aoc2022-day-01 day-01.cpp day-01.h)
target_include_directories(aoc2022-day-01 INTERFACE ${CMAKE_CURRENT_SOURCE_DIR})

target_link_libraries(aoc2022-day-01
	PUBLIC utility)

#add_executable(day-01_test
#	test.cpp
#)
#target_link_libraries(day-01_test
#  GTest::gtest_main
#)
#include(GoogleTest)
#gtest_discover_tests(day-01_test)
