require "./packet/operator"

class Maximum < Operator
  def value
    subpackets.map(&:value).max
  end
end
