#include <days.h>

#include <cassert>
#include <functional>
#include <iterator>
#include <memory>
#include <optional>
#include <sstream>
#include <unordered_map>
#include <stdexcept>
#include <unordered_set>
#include <ranges>

#include <iostream>
#include <algorithm>
#include <pathfinder.h>
#include <deque>
#include <numeric>

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
	using NEIGHBOURS = std::unordered_set<string>;

	namespace vw = std::views;

	struct Valve
	{
		string name{};
		int flow{ 0 };
		NEIGHBOURS tunnels{};

		bool hasEdge(string name) const
		{
			return tunnels.contains(name);
		}
	};
	using VALVES = std::unordered_map<string, Valve>;

	class ValveGraph : public IWeightedGraph
	{
	public:
		const Valve& valve(NodeHandle handle) const
		{
			return nodes.at(name(handle));
		}
		double cost(NodeHandle from, NodeHandle to) const override
		{
			const Valve& fromNode = nodes.at(name(from));
			const Valve& toNode = nodes.at(name(to));

			if (!fromNode.hasEdge(toNode.name)) throw invalid_argument("from and to must be adjacent");

			return 1;
		}
		vector<NodeHandle> allNodes() const override
		{
			vector<NodeHandle> result(nodes.size());
			transform(nodes.begin(), nodes.end(), result.begin(),
					  [&](auto& pair) { return handle(pair.second.name); });
			return result;
		}
		size_t numberOfNodes() const override
		{
			return nodes.size();
		}
		vector<NodeHandle> neighbours(NodeHandle handle) const override
		{
			const Valve& node = nodes.at(name(handle));

			vector<NodeHandle> result(node.tunnels.size());
			transform(node.tunnels.begin(),
					  node.tunnels.end(),
					  result.begin(),
					  [&](string name) { return findNode(name); });
			return result;
		}

		NodeHandle findNode(const string& name) const
		{
			if (!nodes.contains(name)) return nullptr;

			return handle(nodes.at(name).name);
		}

		NodeHandle addNode(const Valve& node)
		{
			if (nodes.contains(node.name)) throw invalid_argument("id already exists");

			nodes[node.name] = node;
			return handle(nodes.at(node.name).name);
		}

	private:
		VALVES nodes;

		NodeHandle handle(const string& name) const
		{
			return reinterpret_cast<NodeHandle>(&name);
		}

		const string& name(NodeHandle handle) const
		{
			return *reinterpret_cast<const string*>(handle);
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
		vector<string> tunnels;
		while (it != end) { // DD, II, BB
			if (*it >= 'A' && *it <= 'Z') {
				tunnels.push_back({ it, it + 2 });
				it += 2;
			} else {
				++it;
			}
		}

		return Valve{ name, flowRate, NEIGHBOURS(tunnels.begin(), tunnels.end()) };
	}

	VALVES parseValves(const vector<string>& data)
	{
		VALVES valves;
		for (auto& s : data) {
			Valve v = parseValve(s);
			valves[v.name] = v;
		}
		return valves;
	}

	void printDistances(const ValveGraph& graph, const FloydWarshallDistances& distances)
	{
		for (auto& handleV : graph.allNodes()) {
			std::cout << "Distances from " << graph.valve(handleV).name << "\n";
			for (auto& handleU : graph.allNodes()) {
				if (handleV == handleU) continue;
				std::cout
					<< "\t to "
					<< graph.valve(handleU).name << ": "
					<< distances.at(handleV).at(handleU) << "\n";
			}
		}
	}

	void printPaths(const ValveGraph& graph, const FloydWarshallNexts& nexts)
	{
		for (auto& handleV : graph.allNodes()) {
			std::cout << "Paths from " << graph.valve(handleV).name << "\n";
			for (auto& handleU : graph.allNodes()) {
				if (handleV == handleU) continue;
				std::cout << "\t to " << graph.valve(handleU).name << ": ";
				auto path = reconstruct_path(handleV, handleU, nexts);
				for (auto& handle : path) {
					std::cout << graph.valve(handle).name << " ";
				}
				std::cout << "(" << path.size() << ")\n";
			}
		}
	}


	struct State
	{
		NodeHandle current;
		std::unordered_set<NodeHandle> opened;
		int elapsed_seconds;
		int relieved_pressure;
	};

	struct SeenState
	{
		std::unordered_set<NodeHandle> opened;
		int elapsed_seconds;
		int relieved_pressure;

	};
	bool operator==(const SeenState& lhs, const SeenState& rhs)
	{
		return lhs.elapsed_seconds == rhs.elapsed_seconds
			&& lhs.relieved_pressure == rhs.relieved_pressure
			&& lhs.opened == rhs.opened;
	}

	auto relieved_per_min(const std::unordered_set<NodeHandle>& opened, const ValveGraph& graph)
	{
		auto relieved_per_min_by_valve = opened
			| vw::transform([&graph](auto handle) { return graph.valve(handle).flow; })
			| vw::common;
		return std::accumulate(
			relieved_per_min_by_valve.begin(),
			relieved_per_min_by_valve.end(),
			0);
	}

	int wait_until_end(int end_time, int elapsed_seconds, int relieved_pressure, const std::unordered_set<NodeHandle>& opened, const ValveGraph& graph)
	{
		auto time_left = end_time - elapsed_seconds;
		return relieved_pressure + time_left * relieved_per_min(opened, graph);
	}

	// If I am at the valve at [position], I've opened a set of valves [opened], and I have [remaining]
	// minutes remaining,
	//		(PART TWO) and there are [helpers] acting after me, 
	// how many points can I score from this position?
	int solve(const VALVES& valves, int end_time = 30/*, helpers = 0*/)
	{
		ValveGraph graph; for (auto& valve : valves) { graph.addNode(valve.second); }
		FloydWarshallDistances distances;	// distances[u][v] = (double) distance from u to v
		FloydWarshallNexts nexts;			// nexts[u][target] = (handle) next step from u to get to target
		floyd_warshall(graph, distances, nexts);
		printDistances(graph, distances);	// checks out
		printPaths(graph, nexts);			// checks out

		auto hasFlow = [&graph](auto& handle) { return graph.valve(handle).flow > 0; };
		auto flowing_view = graph.allNodes() | vw::filter(hasFlow);
		auto flowing = std::vector<NodeHandle>(flowing_view.begin(), flowing_view.end());

		auto maxRelieved = 0;
		auto queue = std::deque<State>();
		auto seen = std::vector<SeenState>();

		queue.push_back(State{ graph.findNode("AA"), {}, 0, 0 });
		seen.push_back(SeenState{ {}, 0, 0 });

		while (!queue.empty()) {
			State state = queue.front(); queue.pop_front();
			// if all flowing valves are opened, wait until the end
			if (state.opened.size() == flowing.size() || state.elapsed_seconds >= end_time) {
				auto relieved_at_end = wait_until_end(end_time, state.elapsed_seconds, state.relieved_pressure, state.opened, graph);
				if (relieved_at_end > maxRelieved) maxRelieved = relieved_at_end;
				continue;
			}

			// for every unopened valve, run simulation
			auto handleNotInOpened = [&state](auto& handle) {
				return std::none_of(
					state.opened.begin(),
					state.opened.end(),
					[&handle](auto& openHandle) {
						return openHandle == handle;
					}
				);
			};
			auto unopened = flowing | vw::filter(handleNotInOpened);

			for (auto& destination : unopened) {
				auto cost = int(distances[state.current][destination] + 1);
				auto newElapsed = state.elapsed_seconds + cost;
				// if opening the dest valve would exceed the time limit, wait until the end
				if (newElapsed >= end_time) {
					auto relieved_at_end = wait_until_end(end_time, state.elapsed_seconds, state.relieved_pressure, state.opened, graph);
					if (relieved_at_end > maxRelieved) maxRelieved = relieved_at_end;
					continue;
				}
				// relieve pressure of existing opened valves while we move to dest and open it
				auto newRelieved = state.relieved_pressure + relieved_per_min(state.opened, graph) * cost;
				std::unordered_set<NodeHandle> newOpened(state.opened);
				newOpened.insert(destination);

				SeenState newSeenState{ newOpened, newElapsed, newRelieved };
				if (std::none_of(seen.begin(), seen.end(), [&newSeenState](auto& seenState) { return seenState == newSeenState; })) {
					queue.push_back(State{ destination, newOpened, newElapsed, newRelieved });
				}
			}

		}

		return maxRelieved;
	}

	string answer_a(const vector<string>& input_data)
	{
		return std::to_string(solve(parseValves(input_data), 30));
	}

	string answer_b(const vector<string>& input_data)
	{
		// return std::to_string(solve(parseValves(input_data), 26, 1));
		return "PENDING";
	}
}
