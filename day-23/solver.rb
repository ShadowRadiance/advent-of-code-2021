# frozen_string_literal: true

require "./pods/map"
require "./heap"

class Solver
  Map = Pods::Map
  Node = Struct.new(:state, :cost_to_get_here, :parent, :moves_to_get_here, keyword_init: true) do
    include Comparable
    def ==(other)
      state == other.state
    end
    def <=>(other)
      cost_to_get_here <=> other.cost_to_get_here
    end
  end

  def initialize(initial, solved=".|.|AA|.|BB|.|CC|.|DD|.|.")
    @initial_state = Map.new(initial).state
    @solved_state  = Map.new(solved).state
  end

  def solve
    winning_node = do_a_dijkstra
    
    puts "STATE:\n  #{winning_node.state}"
    puts "COST: \n  #{winning_node.cost_to_get_here}"
    puts "MOVES:\n  #{winning_node.moves_to_get_here.join("\n") }"

    winning_node.cost_to_get_here
  end

  private

  def do_a_dijkstra
    # 1. Set all points distance to infinity except for the starting point set distance to 0.
    # 2. Set all points, including starting point as a non-visited node.
    # 3. Set the non-visited node with the smallest current distance as the current node "C."
    # 4. For each neighbor "N" of your current node: add the current distance of "C" with the weight of
    # the edge connecting "C"->"N." If it's smaller than the current distance of "N", set it as the new
    # current distance of "N".
    # 5. Mark the current node, "C", as visited.
    # 6. Repeat the step above from step 3 until the destination point is visited.

    # 1. Set all points distance to infinity except for the starting point set distance to 0.
      # no other nodes yet, we need to ask the map of root's state for the next moves
    root_node = Node.new(state: @initial_state, cost_to_get_here: 0, parent: nil, moves_to_get_here: [])
    # 2. Set all points, including starting point as a non-visited node.
      # no other nodes yet, we need to ask the map of root's state for the next moves
    unvisited_nodes = Heap.new
    unvisited_nodes << root_node
    visited_nodes = []

    until unvisited_nodes.empty? do
      # 3. Set the non-visited node with the smallest current distance as the current node "C."
      # unvisited_nodes = unvisited_nodes.sort_by(&:cost_to_get_here).reverse
      current_node = unvisited_nodes.extract
      current_map = Map.new(current_node.state)
      
      # 4. For each neighbor "N" of your current node: 
      #    add the current distance of "C" with the weight of the edge connecting "C"->"N". 
      #    If it's smaller than the current distance of "N", set it as the new current distance of "N".
      moves = current_map.all_possible_moves.map(&:to_s) # to_s to detach from current_map
      moves.each { |move| 
        new_map = Map.new(current_node.state)
        cost = new_map.apply_move_str(move)

        create_or_update_node(
          unvisited_nodes, 
          Node.new(
            state: new_map.state,
            cost_to_get_here: current_node.cost_to_get_here + cost,
            parent: current_node,
            moves_to_get_here: current_node.moves_to_get_here + [move],
          )
        )
      }
      
      # 5. Mark the current node, "C", as visited.
      visited_nodes.push(current_node)
      
      puts "#{visited_nodes.size} / #{current_node.cost_to_get_here} / #{unvisited_nodes.size}"

      # 6. Repeat the step above from step 3 until the destination point is visited.
      return current_node if current_node.state == @solved_state
    end
    raise "WTF"
  end

  def create_or_update_node(node_list, potential)
    existing_node = node_list.to_a.find { |node| node.state == potential.state }

    if existing_node
      return unless potential.cost_to_get_here < existing_node.cost_to_get_here
      existing_node.cost_to_get_here = potential.cost_to_get_here
      existing_node.parent = potential.parent
      existing_node.moves_to_get_here = potential.moves_to_get_here
      node_list.rerank(existing_node)
    else
      node_list << potential
    end
  end
end