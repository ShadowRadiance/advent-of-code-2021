require "./packet/operator"

class LessThan < Operator
  def value
    subpackets.first.value < subpackets.last.value ? 1 : 0
  end
end
