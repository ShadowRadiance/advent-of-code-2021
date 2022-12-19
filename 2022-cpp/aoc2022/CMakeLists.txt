foreach(NUM RANGE 1 25)
	set(PADDED "${NUM}")
	if(NUM LESS 10)
		string(PREPEND NUM "0")
	endif()
	add_library(aoc2022-day-${NUM} src/day-${NUM}.cpp)
	target_include_directories(aoc2022-day-${NUM} PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include)
	target_link_libraries(aoc2022-day-${NUM} PUBLIC utility)

	set(DAY_LIBS ${DAY_LIBS} aoc2022-day-${NUM})
	file(COPY ${CMAKE_CURRENT_SOURCE_DIR}/data/day-${NUM}.txt DESTINATION ${CMAKE_CURRENT_BINARY_DIR}/data)
endforeach()

add_executable(aoc2022 main.cpp)
target_include_directories(aoc2022 PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/include)
target_link_libraries(aoc2022 PRIVATE ${DAY_LIBS} utility)



add_subdirectory(tests)

# target_include_directories(<target> PUBLIC|PRIVATE|INTERFACE <list>)
#	Source files use library headers but your headers do not use them: PRIVATE
#	Source files and your headers use library headers: PUBLIC
#	Source files don't use library headers but your headers do use them: INTERFACE


