require "./packet/operator"

class EqualTo < Operator
  def value
    subpackets.first.value == subpackets.last.value ? 1 : 0
  end
end
