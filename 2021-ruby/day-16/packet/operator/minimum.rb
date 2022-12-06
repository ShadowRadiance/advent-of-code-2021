require "./packet/operator"

class Minimum < Operator
  def value
    subpackets.map(&:value).min
  end
end
