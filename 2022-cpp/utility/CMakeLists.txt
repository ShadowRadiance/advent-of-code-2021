add_library(utility utility.cpp utility.h)
target_include_directories(utility PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

add_executable(utility-tests utility_test.cpp)
target_include_directories(utility-tests PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
target_link_libraries(utility-tests PUBLIC utility GTest::gtest_main)
include(GoogleTest)
gtest_discover_tests(utility-tests)
