enable_testing()

add_executable(
  day-01_test
  test.cpp
)
target_link_libraries(
  day-01_test
  GTest::gtest_main
)

include(GoogleTest)
gtest_discover_tests(day-01_test)
