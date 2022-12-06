require "./packet/operator"

class Product < Operator
  def value
    subpackets.map(&:value).reduce(&:*)
  end
end
