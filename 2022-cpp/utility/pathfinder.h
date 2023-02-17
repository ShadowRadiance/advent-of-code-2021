#pragma once

#include <vector>
#include <unordered_map>
#include <functional>

namespace advent_of_code
{
	namespace Pathfinder
	{
		class IGraph
		{
		public:
			typedef void* NodeHandle;
			virtual std::vector<NodeHandle> neighbours(NodeHandle handle) const = 0;
		};

		class IWeightedGraph : public IGraph
		{
		public:
			virtual double cost(NodeHandle from, NodeHandle to) const = 0;
		};

		class SimpleGraph : public IWeightedGraph
		{
		public:
			struct Node
			{
				char id;
				std::vector<char> edges;
				bool hasEdge(char id) const;
			};
			double cost(NodeHandle from, NodeHandle to) const override;
			std::vector<NodeHandle> neighbours(NodeHandle handle) const override;
			NodeHandle findNode(char id) const;
			NodeHandle addNode(char id, std::vector<char> edges);
			NodeHandle addNode(const Node& node);
			char id(NodeHandle handle) const;
		private:
			std::vector<Node> nodes;

			NodeHandle handle(ptrdiff_t index) const;
			ptrdiff_t index(NodeHandle handle) const;
			ptrdiff_t index(char id) const;
		};

		using CameFromMap = std::unordered_map<IGraph::NodeHandle, IGraph::NodeHandle>;
		using CostMap = std::unordered_map<IGraph::NodeHandle, double>;
		using Path = std::vector<IGraph::NodeHandle>;

		CameFromMap breadth_first_search(const IGraph& graph,
										 IGraph::NodeHandle start,
										 IGraph::NodeHandle goal);

		void dijkstra_search(const IWeightedGraph& graph,
							 IGraph::NodeHandle start,
							 IGraph::NodeHandle goal,
							 CameFromMap& came_from,
							 CostMap& cost_so_far);

		Path reconstruct_path(IGraph::NodeHandle start,
							  IGraph::NodeHandle goal,
							  CameFromMap came_from);

		using HeuristicFunction = std::function<double(IGraph::NodeHandle, IGraph::NodeHandle)>;

		void a_star_search(const IWeightedGraph& graph,
						   IGraph::NodeHandle start,
						   IGraph::NodeHandle goal,
						   HeuristicFunction heuristic_fn,
						   CameFromMap& came_from,
						   CostMap& cost_so_far);
	} // namespace Pathfinder
} // namespace advent_of_code
