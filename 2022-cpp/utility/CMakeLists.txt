add_library(utility utility.cpp utility.h cpu_10.cpp cpu_10.h pathfinder.cpp pathfinder.h)
target_include_directories(utility PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

add_executable(utility-tests utility_test.cpp 
							 cpu_10_test.cpp cpu_10.cpp 
							 pathfinder_test.cpp pathfinder.cpp)
target_include_directories(utility-tests PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
target_link_libraries(utility-tests PUBLIC utility GTest::gtest_main)
include(GoogleTest)
gtest_discover_tests(utility-tests)
