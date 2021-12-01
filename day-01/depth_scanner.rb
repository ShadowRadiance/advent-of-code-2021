class DepthScanner
  attr_reader :data, :strategy

  def initialize(data = [], strategy: SimpleDepthScanner.new )
    @data = data
    @strategy = strategy
  end

  def increases
    strategy.execute(data)
  end
end

class SimpleDepthScanner
  def execute(data)
    data.each_cons(2).filter { |a,b| b > a }.size
  end
end

class SlidingWindowScanner
  attr_reader :window_size

  def initialize(window_size)
    @window_size = window_size
  end

  def execute(data)
    data.each_cons(window_size).map { |array| array.sum }.each_cons(2).filter { |a,b| b>a }.size
  end
end
