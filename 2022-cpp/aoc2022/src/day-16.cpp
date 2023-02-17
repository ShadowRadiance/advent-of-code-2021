#include <days.h>

#include <cassert>
#include <functional>
#include <iterator>
#include <memory>
#include <optional>
#include <sstream>
#include <unordered_map>

#include <iostream>
#include <algorithm>

namespace day_16
{
    using std::string;
    using std::vector;
    using maybe_int = std::optional<size_t>;

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

    auto floyd_warshall(const vector<Valve>& valves)
    {
        auto size = valves.size();
        std::vector<std::vector<size_t>> distances(size, std::vector<size_t>(size, INT_MAX));
        std::vector<std::vector<maybe_int>> nexts(size, std::vector<maybe_int>(size));

        // record the vertices
        for (int i = 0; i < size; ++i) {
            distances[i][i] = 0;
            nexts[i][i] = i;
        }

        // record the edges
        for (int i = 0; i < size; ++i) {
            auto& valveV = valves[i];
            for (auto& name : valveV.connectedValveNames) {
                auto it = std::find_if(
                    valves.begin(),
                    valves.end(),
                    [&](auto& valveU) { return valveU.name == name; });
                if (it != valves.end()) {
                    auto j = std::distance(valves.begin(), it);
                    auto& valveU = *it;
                    distances[i][j] = 1; // edge weight
                    distances[j][i] = 1; // reverse link
                    nexts[i][j] = j;
                }
            }
        }

        // determine shortest paths to all other nodes
        for (int k = 0; k < size; ++k) {
            for (int i = 0; i < size; ++i) {
                for (int j = 0; j < size; ++j) {
                    size_t newDistance = distances[i][k] + distances[k][j];
                    if (distances[i][j] > newDistance) {
                      distances[i][j] = newDistance;
                      nexts[i][j] = nexts[i][k];
                    }
                }
            }
        }

        return std::make_tuple(distances, nexts);
    }

    vector<int> findPath(
        size_t fromIndex, 
        size_t toIndex,
        std::vector<std::vector<maybe_int>> &f_w_next_matrix)
    {
        vector<int> path;
        if (!f_w_next_matrix[fromIndex][toIndex].has_value())
            return path;
        
        path.push_back(fromIndex);
        while (fromIndex != toIndex) {
            fromIndex = f_w_next_matrix[fromIndex][toIndex].value();
            path.push_back(fromIndex);
        }
        return path;
    }

    void printPaths(auto nexts, const vector<Valve> &valves) {
        auto size = valves.size();
        for (size_t u = 0; u < size; u++) {
            std::cout << "Paths from " << valves[u].name << "\n";
            for (size_t v = 0; v < size; v++) {
                if (u == v)
                    continue;
                std::cout << "\t to " << valves[v].name << ": ";
                auto indices = findPath(u, v, nexts);
                for (int idx : indices) {
                    std::cout << valves[idx].name << " ";
                }
                std::cout << "(" << indices.size() << ")\n";
            }
        }
    }

    int solve(const vector<Valve>& valves)
    {
        auto [distances, nexts] = floyd_warshall(valves);
        
        printPaths(nexts, valves);

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
