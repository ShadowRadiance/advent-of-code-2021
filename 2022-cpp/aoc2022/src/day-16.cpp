#include <days.h>

#include <cassert>
#include <functional>
#include <iterator>
#include <memory>
#include <optional>
#include <sstream>
#include <unordered_map>
#include <stdexcept>

#include <iostream>
#include <algorithm>
#include <pathfinder.h>

namespace day_16
{
	using std::string;
	using std::vector;
	using std::invalid_argument;
	using advent_of_code::Pathfinder::floyd_warshall;
	using advent_of_code::Pathfinder::reconstruct_path;
	using advent_of_code::Pathfinder::FloydWarshallDistances;
	using advent_of_code::Pathfinder::FloydWarshallNexts;
	using advent_of_code::Pathfinder::IWeightedGraph;
	using NodeHandle = advent_of_code::Pathfinder::IGraph::NodeHandle;

	struct Valve
	{
		string name{};
		int flow{ 0 };
		vector<string> connectedValveNames{};

		bool hasEdge(string name) const
		{
			return std::find(connectedValveNames.begin(), connectedValveNames.end(), name)
				!= connectedValveNames.end();
		}
	};

	class ValveGraph : public IWeightedGraph
	{
	public:
		Valve valve(NodeHandle handle) const
		{
			return nodes[index(handle)];
		}
		double cost(NodeHandle from, NodeHandle to) const override
		{
			const Valve& fromNode = nodes[index(from)];
			const Valve& toNode = nodes[index(to)];

			if (!fromNode.hasEdge(toNode.name)) throw invalid_argument("from and to must be adjacent");

			return 1;
		}
		vector<NodeHandle> allNodes() const override
		{
			vector<NodeHandle> result(nodes.size());
			for (ptrdiff_t i = 0; i < nodes.size(); i++) {
				result[i] = handle(i);
			}
			return result;
		}
		size_t numberOfNodes() const override
		{
			return nodes.size();
		}
		vector<NodeHandle> neighbours(NodeHandle handle) const override
		{
			const Valve& node = nodes[index(handle)];

			vector<NodeHandle> result(node.connectedValveNames.size());
			transform(node.connectedValveNames.begin(),
					  node.connectedValveNames.end(),
					  result.begin(),
					  [&](string name) { return findNode(name); });
			return result;
		}

		NodeHandle findNode(string name) const
		{
			return handle(index(name));
		}

		NodeHandle addNode(string name, int flow, vector<string> edges)
		{
			return addNode(Valve{ name, flow, edges });
		}

		NodeHandle addNode(const Valve& node)
		{
			if (index(node.name) != nodes.size()) throw invalid_argument("id already exists");

			nodes.push_back(node);
			return handle(index(node.name));
		}

	private:
		vector<Valve> nodes;

		ptrdiff_t index(string name) const
		{
			auto it = find_if(nodes.begin(), nodes.end(), [name](auto& node) { return node.name == name; });
			return it - nodes.begin();
		}

		NodeHandle handle(ptrdiff_t index) const
		{
			// the 0 handle, nullptr, would equate to index==-1
			// valid handles are > 0

			return reinterpret_cast<NodeHandle>(index + 1);
		}

		ptrdiff_t index(NodeHandle handle) const
		{
			// the 0 handle, nullptr, would equate to index==-1
			// valid handles are > 0

			return reinterpret_cast<ptrdiff_t>(handle) - 1;
		}

	};


	int parseInt(auto& it, auto end)
	{
		size_t size{ 0 };
		int parsed = std::stoi(string{ it, end }, &size);
		it += size;
		return parsed;
	}

	Valve parseValve(const string& str)
	{
		auto it = str.begin();
		auto end = str.end();

		assert(*it == 'V');
		it += 6;                 // Valve_
		string name{ it, it + 2 }; // AA
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
				connectedValveNames.push_back({ it, it + 2 });
				it += 2;
			} else {
				++it;
			}
		}

		return Valve{ name, flowRate, connectedValveNames };
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


	void printPaths(const ValveGraph& graph, const FloydWarshallNexts& nexts, const vector<Valve>& valves)
	{
		for (auto& valve : valves) {
			std::cout << "Paths from " << valve.name << "\n";
			NodeHandle vHandle = graph.findNode(valve.name);
			for (auto& other : valves) {
				if (&valve == &other) continue;
				std::cout << "\t to " << other.name << ": ";
				NodeHandle uHandle = graph.findNode(other.name);
				auto path = reconstruct_path(vHandle, uHandle, nexts);
				for (auto& handle : path) {
					std::cout << graph.valve(handle).name << " ";
				}
				std::cout << "(" << path.size() << ")\n";
			}
		}
	}

	int solve(const vector<Valve>& valves)
	{
		ValveGraph graph;
		for (auto& valve : valves) { graph.addNode(valve); }
		FloydWarshallDistances distances;
		FloydWarshallNexts nexts;
		floyd_warshall(graph, distances, nexts);

		printPaths(graph, nexts, valves);

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
