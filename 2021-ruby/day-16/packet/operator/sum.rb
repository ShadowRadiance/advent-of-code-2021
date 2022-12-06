require "./packet/operator"

class Sum < Operator
  def value
    subpackets.sum(&:value)
  end
end
