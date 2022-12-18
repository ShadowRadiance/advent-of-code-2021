#include <iostream>
#include <vector>
#include <format>

std::vector<std::string> load_data(const std::string& filename)
{
	std::vector<std::string> result;
	result.push_back(std::format("Loading {}!", filename));
	return result;
}
