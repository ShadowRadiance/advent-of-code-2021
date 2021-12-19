# frozen_string_literal: true

class SnailfishNumber

  def initialize(other)
    @top_level_node = case other
                      when SnailfishNumber then init_from_array(other.to_a)
                      when Array then init_from_array(other)
                      when String then init_from_string(other)
                      else raise TypeError
                      end
    normalize
  end

  def left
    @top_level_node.left
  end

  def right
    @top_level_node.right
  end

  def magnitude
    @top_level_node.magnitude
  end

  def +(other)
    raise TypeError unless SnailfishNumber === other

    SnailfishNumber.new([self, other])
  end

  def to_a
    @top_level_node.to_a
  end

  def to_s
    @top_level_node.to_s
  end

  def normalize
    # puts to_s
    loop do
      break unless explode || split
    end
  end

  def find_node(under, level, &block)
    return under if block.call(under, level)
    return nil unless Pair===under

    find_node(under.left, level+1, &block) || find_node(under.right, level+1, &block)
  end

  def explode
    explodanode = find_node(@top_level_node, 0) do |node, level|
      level >= 4 && Pair===node && Regular===node.left && Regular===node.right
    end
    return false if explodanode.nil?

    explodanode.explode!
    # puts to_s

    true
  end

  def split
    splittanode = find_node(@top_level_node, 0) do |node, _level|
      Regular===node && node.value >= 10
    end
    return false if splittanode.nil?

    splittanode.split!
    # puts to_s
    true
  end

  private

  def init_from_array(arr)
    Node.from(arr, nil)
  end

  def init_from_string(_str)
    raise "Ugh... still have two write an array string parser"
  end

  class Node
    attr_reader :parent

    def initialize(parent = nil, *_args, **_kwargs)
      @parent = parent
    end

    class << self
      def from(other, parent=nil)
        case other
        when Regular then Regular.new(parent, other.value)
        when Pair, SnailfishNumber then Pair.new(parent, other.left, other.right)
        when Integer then Regular.new(parent, other)
        when Array then Pair.new(parent, other.first, other.last)
        else raise TypeError
        end
      end

      protected :new
    end

  end

  class Regular < Node
    attr_reader :value

    def initialize(parent, value)
      super
      @value = value
    end

    def absorb(other)
      @value += other
    end

    def split!
      # To split a regular number, replace it with a pair; 
      # the left element of the pair should be the regular number divided by two and rounded down, 
      # while the right element of the pair should be the regular number divided by two and rounded up. 
      # For example, 10 becomes [5,5], 11 becomes [5,6], 12 becomes [6,6], and so on.
      parent&.replace(self, [value/2, (value+1)/2])
    end

    def magnitude
      value
    end

    def to_a
      value
    end

    def to_s
      value.to_s
    end
  end

  class Pair < Node
    def initialize(parent, left, right)
      super
      @left = Node.from(left, self)
      @right = Node.from(right, self)
    end

    def left(except: nil)
      return nil if @left.equal?(except)

      @left
    end

    def right(except: nil)
      return nil if @right.equal?(except)

      @right
    end

    def magnitude
      3 * @left.magnitude + 2 * @right.magnitude
    end

    def explode!
      l_val = @left.value
      r_val = @right.value

      closest_regular_node_left&.absorb(l_val)
      closest_regular_node_right&.absorb(r_val)

      parent&.replace(self, 0)
    end

    def replace(node_to_replace, value)
      new_node = Node.from(value, self)
      if left.equal?(node_to_replace)
        @left = new_node
      elsif right.equal?(node_to_replace)
        @right = new_node
      else
        raise ArgumentError, "cannot replace a node I don't own"
      end
    end

    def to_a
      [@left, @right].map(&:to_a)
    end

    def to_s
      "[#{[@left, @right].map(&:to_s).join(",")}]"
    end

    def right_most_regular
      node = right
      node = node.right until Regular===node
      node
    end

    def left_most_regular
      node = left
      until Regular===node
        node = node.left
      end
      node
    end

    def closest_regular_node_left
      node = parent&.left(except: self)
      return node if Regular===node

      node&.right_most_regular || parent&.closest_regular_node_left
    end

    def closest_regular_node_right
      node = parent&.right(except: self)
      return node if Regular===node

      node&.left_most_regular || parent&.closest_regular_node_right
    end
  end
end
