add_subdirectory(src)

add_executable(aoc2022 main.cpp)
target_link_libraries(aoc2022
	PUBLIC aoc2022-day-01)

