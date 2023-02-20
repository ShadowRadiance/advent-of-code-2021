#include <fstream>
#include <utility.h>

using std::vector;
using std::string;
using std::ifstream;

vector<string> load_data(const string& filename)
{
    vector<string> result;

    ifstream input{ filename };

    string line;
    if (input.is_open()) {
        while (getline(input, line)) {
            result.push_back(line);
        }
    }

    return result;
}

