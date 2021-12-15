class Polymerizer
  attr_reader :rules
  def initialize(rules)
    @rules = rules
  end

  def step(template)
    return "" if template.empty?

    result = []
    template.chars.each_cons(2) { |first, second|
      result << apply_rule(first, second)
    }
    result.flatten.join + template[-1]
  end

  def steps(number, template)
    return template if number == 0

    step( steps(number-1, template) )
  end

  private

  def apply_rule(first, second)
    # [first, rules[[first,second].join]]

    key = [first,second].join
    injection = rules[key]
    [first, injection]
  end
end