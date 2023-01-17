#include <days.h>

#include <cassert>
#include <functional>
#include <iterator>
#include <memory>
#include <optional>
#include <sstream>
#include <unordered_map>

namespace day_16
{
  using std::string;
  using std::vector;

    struct Valve
    {
        string name{};
        int flow{0};
        vector<string> connectedValveNames{};
    };

    int parseInt(auto& it, auto end)
    {
        size_t size{0};
        int parsed = std::stoi(string{it, end}, &size);
        it += size;
        return parsed;
    }

    Valve parseValve(const string& str)
    {
        auto it = str.begin();
        auto end = str.end();

        assert(*it == 'V');
        it += 6;                 // Valve_
        string name{it, it + 2}; // AA
        it += 2;
        assert(*it == ' ');
        it += 15; // _has_flow_rate=
        int flowRate = parseInt(it, end);
        assert(*it == ';');
        it += 23;             // ;_tunnels_lead_to_valve
        if (*it == 's') ++it; // s
        vector<string> connectedValveNames;
        while (it != end) { // DD, II, BB
            if (*it >= 'A' && *it <= 'Z') {
                connectedValveNames.push_back({it, it + 2});
                it += 2;
            } else {
                ++it;
            }
        }

        return Valve{name, flowRate, connectedValveNames};
    }

    vector<Valve> parseValves(const vector<string>& data)
    {
        vector<Valve> valves;
        std::transform(data.begin(),
                       data.end(),
                       back_inserter(valves),
                       parseValve);
        return valves;
    }

    vector<vector<int>> buildAdjancyMatrix(const vector<Valve>& valves) {
        int size = valves.size();
        vector<vector<int>> adjancyMatrix(size, vector<int>(size));
        for (size_t i = 0; i < size; i++) {
            Valve valve = valves[i];
            for (string name : valve.connectedValveNames) {
                auto it = std::find_if(
                    valves.begin(),
                    valves.end(),
                    [&name](auto& _) { return _.name == name; });
                if (it == valves.end())
                    throw std::runtime_error("expected to find " + name +
                                             " in valves!");
                int indexOfConnectedValve = std::distance(valves.begin(), it);
                adjancyMatrix[i][indexOfConnectedValve] = 1; // link forward
                adjancyMatrix[indexOfConnectedValve][i] = 1; // link reverse
            }
        }
        return adjancyMatrix;
    }

    int solve(const vector<Valve>& valves)
    {
        vector<vector<int>> adjancyMatrix = buildAdjancyMatrix(valves);

        return 0;
    }

    string answer_a(const vector<string>& input_data)
    {
        return std::to_string(solve(parseValves(input_data)));
    }

  string answer_b(const vector<string>& input_data)
  {
    return "PENDING";
  }
}
