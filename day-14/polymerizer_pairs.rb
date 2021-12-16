class PolymerizerPairs
  attr_reader :rules, :template, :pairs

  def initialize(rules, template)
    @rules = rules # { "string" => new-center-character }
    @template = template

    @elements = {}
    rules.keys.each { |k, v| @elements[k[0]] =  @elements[k[1]] = 0 }
    template.chars.each { |ch| @elements[ch] += 1 }

    @pairs = @rules.transform_values { 0 } # { "XY" => count }
    template.chars.each_cons(2) { |pair| @pairs[pair.join] += 1 }
  end

  def step
    return if @pairs.empty?

    # each pair is replaced with a specific duo of pairs so:
    # the count of the input pair goes down
    # the count of the output pairs go up
    # the elements count goes up for the rule_char

    changes = @rules.transform_values { 0 }
    @pairs.each do |pair, count|
      # remove count copies of this pair
      changes[pair] -= count

      rule_char = rules[pair]
      @elements[rule_char] += count

      new_pairs = [
        pair[0] + rule_char,
        rule_char + pair[1],
      ]

      # add count copies of the new pairs
      new_pairs.each { |np| changes[np] += count  }
    end

    changes.each { |pair, change| @pairs[pair] += change }
  end

  def steps(number)
    return template if number == 0

    number.times { step }
  end

  def most_common
    # [ab, bb, ba, aa] => "abbaa" => [a,3]
    @elements.max_by { |_, v| v }
  end

  def least_common
    @elements.select { |_, v| v.positive? }.min_by { |_, v| v }
  end

  def variance
    min, max = @elements.values.select(&:positive?).minmax
    max - min
  end
end