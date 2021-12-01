class DepthScanner
  attr_reader :data
  def initialize(data = [])
    @data = data
  end

  def increases
    data.each_cons(2).filter { |a,b| b > a }.size
  end
end
