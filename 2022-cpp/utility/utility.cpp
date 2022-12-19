#include <fstream>
#include <vector>
#include <format>

#include <utility.h>

std::vector<std::string> load_data(const std::string& filename)
{
	std::vector<std::string> result;

	std::ifstream input{ filename };

	std::string line;
	if (input.is_open()) {
		while (std::getline(input, line)) {
			result.push_back(line);
		}
	}

	return result;
}

